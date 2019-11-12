package rpc

import (
	"encoding/binary"
	"github.com/google/uuid"
)

//type UUID [16]byte
type UUID struct {
	C1 uint32 `struc:"uint32,little"`
	C2 uint16 `struc:"uint16,little"`
	C3 uint16 `struc:"uint16,little"`
	C4 uint16 //Last two components are encoded as bigendian.
	C5 [6]byte
}

func (id *UUID) FromBytes(bytes []byte) {
	if len(bytes) != 16 {
		panic("UUID->FromBytes expected 16bytes")
	}
	id.C1 = binary.BigEndian.Uint32(bytes[0:4])
	id.C2 = binary.BigEndian.Uint16(bytes[4:6])
	id.C3 = binary.BigEndian.Uint16(bytes[6:8])
	id.C4 = binary.BigEndian.Uint16(bytes[8:10])
	copy(id.C5[:], bytes[10:16])
}

func (id *UUID) SetRandom() {
	rand := uuid.New()
	id.FromBytes(rand[:])
}

const (
	ProtoVerMajor       = 5
	ProtoVerMinor       = 1 //Fragmentation is supported for Auth
	ProtoVerMinorNoFrag = 0
)

const RPCConnProtoID = 0x0B //Protocol ID for connection oriendted DCERPC.

// RPC Packet types.
const (
	PktRequest     = 0
	PktPing        = 1
	PktResponse    = 2
	PktFault       = 3
	PktWorking     = 4
	PktNoCall      = 5
	PktReject      = 6
	PktAck         = 7
	PktQuit        = 8
	PktFack        = 9
	PktQuack       = 10
	PktBind        = 11
	PktBindAck     = 12
	PktBindNack    = 13
	PktAlterCxt    = 14
	PktAlterCxtR   = 15
	PktAuth3       = 16
	PktShutdown    = 17
	PktRemoteAlert = 18
	PktOrphaned    = 19
	PktMaxType     = 20
	PktInvalid     = 0xff
)

//RPC Packet flags.
const (
	FlagFirstFrag         = 0x01
	FlagLastFrag          = 0x02
	FlagAlertPending      = 0x03
	FlagSuppHeaderSigning = 0x04
	FlagConcurrMultiplex  = 0x10
	FlagDidNotExecute     = 0x20
	FlagMaybe             = 0x40
	FlagObjectUUID        = 0x80
)

type RPCOptData struct {
	RPCVer       uint8
	RPCVerMinor  uint8
	Padding      [2]uint8
	DataRepr     [4]uint8
	RejectStatus uint32
	Padding2     [2]uint8
}

type RPCRejectOptData struct {
	Reason uint16
	Info   RPCOptData
}

// Presentation negotiation structures, these are part of PktBind & PktAck

// This is the presentation context identifier.
type RpcPresCtxId uint16

type UUIDInfo struct {
	Id      UUID
	Version uint32 `struc:"uint32,little"`
}

type RPCSyntaxNegotiation struct {
	ContexId            uint16 `struc:"uint16,little"`
	TransferSyntaxCount uint8  `struc:"uint8,sizeof=TransferSyntax"`
	Padding             uint8  `struc:"uint8"`
	AbstractSyntax      UUIDInfo
	TransferSyntax      []UUIDInfo
}

type RPCSyntaxNegotiationList struct {
	Count      uint8    `struc:"uint8,sizeof=SyntaxList"`
	Padding    [3]uint8 `struc:"[3]uint8"`
	SyntaxList []RPCSyntaxNegotiation
}

//type RPCSyntaxNegtiationResult uint16
// RPCSyntaxNegtiationResult values
const (
	RPCSynNegAccepted     = 0
	RPCSynNegUserRejected = 1
	RPCSynNegProvRejected = 2
)

// FIXME: Refer to the specs to find the difference below two error codes.
//Reason codes for RPCSynNegProvRejected - uint16
const (
	PSynNegRejNotSpec                    = 0
	PSynNegRejAbstractSyntaxNotSupported = 1
	PSynNegRejTransferSyntaxNotSupported = 2
	PSynNegRejLocalLimitExceeded         = 3
)

//uint16
const (
	SynNegRejNotSpec                     = 0
	SynNegRejTempCongestion              = 1
	SynNegRejLocalLimitExceeded          = 2
	SynNegRejProtocalVersionNotSupported = 4
	SynNegRejAuthTypeUnknown             = 8
	SynNegRejInvalidChecksum             = 9
)

type RPCSyntaxNegotiationResult struct {
	Result       uint16 `struc:"uint16,little"`
	Reason       uint16 `struc:"uint16,little"`
	AcceptedUuid UUIDInfo
}

type RPCSyntaxNegotiationResultList struct {
	NResults uint8  `struc:"uint8,little,sizeof=Results"`
	Padding  uint8  `struc:"uint8,little"`
	Padding2 uint16 `struc:"uint16,little"`
	Results  []RPCSyntaxNegotiationResult
}

//Version datatypes - used in bind_nack packet.
type RpcVersion struct {
	Major uint8
	Minor uint8
}

type RpcVersionSupported struct {
	NVersions uint8
	Versions  []RpcVersion
}

//Common header of all RPC packets. - 16 bytes in size.
type RpcCommonHdr struct {
	MajVer  uint8    `struc:"uint8,little"`
	MinVer  uint8    `struc:"uint8,little"`
	PType   uint8    `struc:"uint8,little"`
	Flags   uint8    `struc:"uint8,little"`
	Repr    [4]uint8 `struc:"[4]uint8,little"`
	FragLen uint16   `struc:"uint16,little"`
	AuthLen uint16   `struc:"uint16,little"`
	CallId  uint32   `struc:"uint32,little"`
}

// Common Authentication trailer - 8 bytes in length.
// The authentication trailer is only present if auth_length
// in common part of header is non-zero.
// The authentication trailer is variable size.
// The authentication trailer is aligned on a 4-byte boundary.
// The encodings of the auth_value field is authentication service specific.
type RpcComAuthTrailer struct {
	AuthType   uint8
	AuthLevel  uint8
	PaddingLen uint8
	Padding    uint8
	KeyID      uint32 `struc:"uint32,little"`
	//Followed by bytes of size auth_length
}

const RpcComAuthTrailerLen = 8

//TODO: Authentication headers to be added.
// ...

// RpcBindHeader template.
type RpcBindHdr struct {
	MaxXmitFrag   uint16 `struc:"uint16,little"` /* Max transmit fragment size in kb */
	MaxRecvFrag   uint16 `struc:"uint16,little"` /* Max Receive fragement size in kb */
	AssocGrpId    uint32 `struc:"uint32,little"` /* client-server group */
	SyntaxNegList RPCSyntaxNegotiationList
	//RpcComAuthTrailer follows
}

type RpcPortAddress struct {
	Length  uint16 `struc:"uint16,little,sizeof=Address"`
	Address []byte `struc:"[]uint8"`
}

// RpcBindAckHeader template.
type RpcBindAckHdr struct {
	MaxXmitFrag uint16 `struc:"uint16,little"` /* Max transmit fragment size in kb */
	MaxRecvFrag uint16 `struc:"uint16,little"` /* Max Receive fragement size in kb */
	AssocGrpId  uint32 `struc:"uint32,little"` /* client-server group */
	PortAddress RpcPortAddress
	Padding     [3]uint8
	//Fourbyte alignment to be restored.
	SyntaxNegResultList RPCSyntaxNegotiationResultList
	//RPCSyntaxNegotiationList follows
	//RpcComAuthTrailer follows
}

// RpcBindNackHeader template.
type RpcBindNackHdr struct {
	CommonHdr     RpcCommonHdr
	ProvRejReason uint16
	SuppVersions  RpcVersionSupported
}

//RpcRequestHeader template.
type RpcRequestHeader struct {
	AllocHint uint32       `struc:"uint32,little"`
	PresCtxId RpcPresCtxId `struc:"uint16,little"`
	OpNum     uint16       `struc:"uint16,little"` //OpNum will be followed by UUID if FlagObjectUUID is present rpcflags
	//StubData  []byte
	//TODO: will be follwed by optionals RpcAuthHeaders
}

//RpcResponseHeader template.
type RpcResponseHeader struct {
	AllocHint  uint32       `struc:"uint32,little,sizeof=StubData"`
	PresCtxId  RpcPresCtxId `struc:"uint16,little"`
	AlertCount uint8        `struc:"uint8,little"`
	Padding    uint8        `struc:"uint8,little"`
	StubData   []byte
	//Followed by stub data and Authentication headers.
}

//RpcFaultPktHeader template
type RpcFaultPktHeader struct {
	CommonHdr  RpcCommonHdr
	AllocHint  uint32
	PresCtxId  RpcPresCtxId
	AlertCount uint8
	Padding    uint8
	Status     uint32
	Reserved   uint32
	StubData   []byte
	//TODO: will be follwed by optionals RpcAuthHeaders
}
