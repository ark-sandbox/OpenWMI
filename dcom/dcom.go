/* DCOM - Distributed Component Object Model */
package dcom

import (
	"fmt"
	"io"
	"local/msrpc/rpc"
	"math/rand"
	"net"
)

type RpcContext rpc.RpcContext
type RpcInterface rpc.RpcInterface

/* DCOM definations */
type OID uint64   //Object identifier.
type SETID uint64 //Ping Set Identifier.
type HRESULT uint32
type DWORD uint32
type error_status uint32 //FIXME
type GUID = rpc.UUID
type CID = GUID
type CLSID = GUID //Class Identifier.
type IID = GUID   //interface Idenfifier.
type IPID = GUID  //Interafce Pointer Identifier.
type OXID uint64  //Object Exporter Identifier.

type ComVersion struct {
	MajVersion uint16 //5
	MinVersion uint16 //1-7 depending on capability.
}

//Out of band data, that is not part of method signature.
type ORpcExtent struct {
	ReprId GUID   //Representation format of this Extent.
	Size   uint32 //Size should be multiple of 8.
	Data   []byte
}

type ORpcExtentArray struct {
	Size        uint32 //Size should be multiple of 2.
	Reserved    uint32
	ExtentArray []ORpcExtent
}

//Sent as part of ORPC requests(implicit 1st arg) & Activation request.
type ORpcThisTyp struct {
	Version    ComVersion
	Flags      uint32
	Reserverd  uint32 //Set to zero.
	Cid        CID
	Extensions ORpcExtentArray
}

//Sent as part of ORCP response(implicit 1st arg)
type ORpcThatTyp struct {
	Flags      uint32
	Extensions ORpcExtentArray
}

//iid_is is an IDL attribute, refers to Interface as parameter.
//Contains hand-marshalled OBJREF.
type MInterfacePointer struct {
	Size uint32
	Data []byte
}

//Pointer to MInterfacePointer.
type PtrTo_MInterfacePointer struct {
	Pointer    uint32
	MInterfPtr MInterfacePointer
}

type PMInterfacePointer PtrTo_MInterfacePointer

//OBJREF is the marshalld format for DCOM object reference.
const (
	ObjRefStandardTyp = 0x01
	ObjRefHandlerTyp  = 0x02
	ObjRefCustomTyp   = 0x3
	ObjRefExtendedTyp = 0x4
)

type ObjRef struct {
	Signature uint32 //0x5747454d
	flags     uint32 //0x01 = ObjRefStd, 0x02 = ObjRefHandler, 0x3 = ObjRefCustom, 0x4 = ObjRefExtended
	Iid       IID    //Interface Identifier - iid_is IDL attraibute.
	//Followed by either ObjRefStd, ObjRefHandler, ObjRefCustom, ObjRefExtended.
}

const (
	SorfPing   = 0x00000000 //Client should perform GC pinging.
	SorfNoPing = 0x00000000 //Client shouldn't perform garbage collection pinging.
)

//Common structure for ObjRef formats.
type StdObjRef struct {
	Flags      uint32 //SORF_PING or SORF_NOPING
	NoRefCount uint32
	Oxid       OXID
	Oid        OID
	Ipid       IPID
}

//Following structs define four different Object Reference formats.
//OBJREF_STANDARD
type ObjRefStd struct {
	Common             StdObjRef
	ObjResolverAddress DualStringArray
}

//OBJREF_HANDLER
type ObjRefHandler struct {
	Common             StdObjRef
	Clsid              CLSID
	ObjResolverAddress DualStringArray
}

//OBJREF_CUSTOM
type ObjRefCustom struct {
	Clsid      CLSID
	Reservered uint64
	ObjectData []byte
}

//OBJREF_EXTENDED
type ObjRefExtended struct {
	Common             StdObjRef
	Sign1              uint32 //0x4E535956
	ObjResolverAddress DualStringArray
	NoElems            uint32 // Set to 0x01
	Sign2              uint32 //0x4E535956
	ElemArray          ObjRefDataElemet
}

type ObjRefDataElemet struct {
	DataID      GUID //Context identifier.
	SizeInBytes uint32
	SizeInWords uint32
	Data        []byte
}

type NDRString16 []uint16
type DualStringArray struct {
	NumEntires     uint16
	SecurityOffset uint16
	StringArray    []NDRString16 //NumEntries of string.
}
type StringBinding struct {
	WTowerId       uint16   //Protocol Sequence Identifier.
	NetworkAddress []uint16 //host or host[endpoint]
}
type SecurityBinding struct {
	AuthnService uint16 //If RPC_C_AUTHN_DEFAULT, next two fields are not present.
	Reservered   uint16 //Set to be 0xffff
	PrincpleName uint16 //Service Principal Name.

}

var OxidResolverID OXID
var ObjectStore map[CLSID]([]DComClass)
var IIDInstanceStore map[IID]([]IPID)
var IPIDStore map[IPID]DComClass
var RunningORpcListeners []ORpcListener

type DComClass interface {
	CreateInstance() interface{}
	ActivateInterface(IID, uint32) IPID
	DispatchOp(iid IID, opnum uint16)
}

type DComObject struct {
	Clsid        CLSID
	ImplInterf   []IID
	ObjId        OID
	ActiveInterf map[IID]IPID
	RefCnt       uint32
	Bindings     DualStringArray
}

type ORpcListener struct { // Is this ObjectExporter?
	ServerAddress string
	PortNumber    string
	BoundObjects  []*DComObject
}

func (li *ORpcListener) Start() {
	addr := fmt.Sprintf("%s:%s\n", li.ServerAddress, li.PortNumber)
	//Start Listening on a TCP port.
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error on binding ", err)
		return
	}
	fmt.Println("ORpcListener bound on", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Some error occurred in connection accept")
		}
		go li.HandleConnection(conn)
	}
}

func (li *ORpcListener) AddObject(obj *DComObject) {
	//Check If the obj bindings matches ORpcListerner binding.
	li.BoundObjects = append(li.BoundObjects, obj)
}

func (li *ORpcListener) HandleConnection(conn io.Reader) {
}

//ORPC call - read orpc request headers, then look for the ObjectUUID
//provided in the call in IPIDStore & do the call dispatch.
func (li *ORpcListener) HandleORpcRequest(ctx *RpcContext) {
}
func (li *ORpcListener) HandleRpcBindRequest(ctx *RpcContext) {
}
func (li *ORpcListener) HandleRpcAlterRequest(ctx *RpcContext) {
}

var SimpleDComObjectClsId CLSID
var SimpleDComInterfId IID

type SimpleDComObject struct {
	DComObject
	Count int16
}

func (obj *SimpleDComObject) CreateInstance() interface{} {
	obj.Clsid = SimpleDComObjectClsId
	//Adding all the inteface implemented by this DComObject.
	var interfList []IID
	interfList = append(interfList, SimpleDComInterfId)
	obj.ImplInterf = interfList

	//Generating Random ObjectId.
	obj.ObjId = OID(rand.Uint64())
	ObjectStore[obj.Clsid] = append(ObjectStore[obj.Clsid], obj)
	obj.RefCnt = 1

	/* Update the bindings.
	obj.Bindings = {} //Populate bindings here.
	srvAddr =
	portNo =
	for _, orpcListener := RunningORpcListeners {
		if orpcListener.ServerAddress = srvAddr && orpcListener.PortNumber == portNo { //Fill up.
			orpcListener.AddObject(obj)
			return
		}
	}
	//Creating a new ORpcListener since none of them running matches binding.
	orpcListener := orpcListener{srvAddr, portNo, nil}
	orpcListener.Start()
	orpcListener.AddObject(obj)
	*/
	return obj
}

func (obj *SimpleDComObject) ActivateInterface(iid IID, refCnt uint32) IPID {
	var piid IPID
	(&piid).SetRandom()
	IIDInstanceStore[iid] = append(IIDInstanceStore[iid], piid)
	//Check if for current object alreary this interface is activated if so return it.
	//Otherwise add it.
	IPIDStore[piid] = obj
	obj.RefCnt += refCnt
	obj.ActiveInterf[iid] = piid
	return piid
}

func (obj *SimpleDComObject) DispatchOp(iid IID, opnum uint16) {
	switch iid {
	case SimpleDComInterfId:
		switch opnum {
		case 0: //Add to counter.
			fmt.Println("Adding to the counter.")
			obj.Count += 1
		case 1: //Decrement counter.
			fmt.Println("Decrementing from the counter")
			obj.Count -= 1
		}
	default:
		fmt.Printf("for interface %v is yet to be implemented\n")
	}
}

func DispatchORpcCall() {
}
