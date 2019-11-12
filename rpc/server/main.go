package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"local/msrpc/rpc"
	"net"

	"github.com/ThomsonReutersEikon/go-ntlm/ntlm"
	"github.com/davecgh/go-spew/spew"
	"github.com/lunixbochs/struc"
)

type RpcContext rpc.RpcContext
type RpcInterface rpc.RpcInterface

type RemoteActivator struct {
	uuid rpc.UUID
}

func (o *RemoteActivator) Init() {
	o.uuid.FromBytes([]byte("\x4d\x9f\x4a\xb8\x7d\x1c\x11\xcf\x86\x1e\x00\x20\xaf\x6e\x7c\x57"))
}

func (o *RemoteActivator) GetUuid() rpc.UUID {
	return o.uuid
}

func (o *RemoteActivator) InvokeOp(opnum uint16, ctx *RpcContext) error {
	switch opnum {
	case 0:
		break
	default:
		return errors.New("Unsupported RemoteActivator method.")
	}
	return nil
}

type RemoteActivateReq struct {
	ORpcThis             ORpcThisTyp
	ClassID              rpc.UUID
	PwszObjectName       LpwStr
	PObjectStorage       PMInterfacePointer
	ClientImpLevel       DWORD
	Mode                 DWORD
	Interfaces           DWORD
	PIIDs                PIID_ARRAY
	CRequestedProtseqs   uint16
	AReqeuestedProtoseqs UShortArray
}

type RemoteAcitvateResp struct {
	ORpcThat          ORpcThatTyp
	POxid             OXID
	PpdsaOxidBindings PDualStringArray
	PipidRemUnknown   IPID
	PAuthHint         DWord
	PServerVersion    ComVersion
	PHresult          HResult
	PPInterfaceData   PMInterfacePointer_ARRAY
	PResults          HResult_Array
	ErrorCode         ErrorStatusTyp
}

func (o *RemoteActivator) RemoteActivate(ctx *conn) {

}

type OxidResolver struct {
	uuid rpc.UUID
}

func (o *OxidResolver) Init() {
	o.uuid.FromBytes([]byte("\x99\xfc\xfe\xc4\x52\x60\x10\x1b\xbb\xcb\x00\xaa\x00\x21\x34\x7a"))
}

func (o *OxidResolver) GetUuid() rpc.UUID {
	return o.uuid
}

func (o *OxidResolver) InvokeOp(opnum uint16, ctx *RpcContext) error {
	switch opnum {
	case 3:
		o.ServerAlive(ctx)
	default:
		return errors.New("Unsupported OxidResolver method.")
	}
	return nil
}

func (o *OxidResolver) ServerAlive(ctx *RpcContext) {
	fmt.Println("OxidResolver->ServerAlive() called")
	cmnHdr := ctx.cmnHdr
	cmnHdr.PType = rpc.PktResponse
	respHdr := rpc.RpcResponseHeader{}
	respHdr.AllocHint = 4
	respHdr.PresCtxId = 0
	respHdr.AlertCount = 0
	respHdr.Padding = 0
	respHdr.StubData = make([]byte, 4, 4)
	//Calculate Frame length
	cmnHdrLen, _ := struc.Sizeof(&cmnHdr)
	respHdrLen, err := struc.Sizeof(&respHdr)
	cmnHdr.FragLen = uint16(cmnHdrLen + respHdrLen)
	fmt.Printf("Setting fragment length to %v", cmnHdr.FragLen)
	//spew.Dump(cmnHdr)
	buff := new(bytes.Buffer)
	struc.Pack(buff, &cmnHdr)
	fmt.Printf("Packed commonHeader\n")
	//spew.Dump(respHdr)
	err = struc.Pack(buff, &respHdr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Packed responseHdr\n")
	fmt.Println("Debug: bindAckHdr size:", len(buff.Bytes()))
	ctx.conn.Write(buff.Bytes())
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	ctx := new(RpcContext)
	ctx.currentBinding = nil
	ctx.conn = conn
	oxidResolver := OxidResolver{}
	oxidResolver.Init()
	remoteActivator := RemoteActivator{}
	remoteActivator.Init()
	ctx.supportedBindings = append(ctx.supportedBindings, &oxidResolver)
	ctx.supportedBindings = append(ctx.supportedBindings, &remoteActivator)
	for {
		cmnHdr := rpc.RpcCommonHdr{}
		struc.Unpack(conn, &cmnHdr)
		ctx.cmnHdr = cmnHdr
		if !(cmnHdr.MajVer == 5 && cmnHdr.MinVer == 0) {
			fmt.Printf("RPC version mismatched %v, %v", cmnHdr.MajVer, cmnHdr.MinVer)
			conn.Close()
			break
		}

		fmt.Println("*_******************* NEW PDU ********************_*")
		spew.Dump(cmnHdr)
		switch cmnHdr.PType {
		case rpc.PktRequest:
			handleRpcRequest(ctx)
		case rpc.PktBind:
			handleRpcBindRequest(ctx)
		case rpc.PktAlterCxt:
			handleRpcAlterContextRequest(ctx)
		case rpc.PktAuth3:
			handleRpcPktAuth(ctx)
		default:
			fmt.Println("Unsupported Packet Type\n")
			conn.Close()
		}
	}
}

func handleRpcPktAuth(ctx *RpcContext) {
	fmt.Println("AUTH3 packet received\n")
	hdr := make([]byte, 4)
	io.ReadFull(ctx.conn, hdr)
	authHdr := rpc.RpcComAuthTrailer{}
	struc.Unpack(ctx.conn, &authHdr)
	spew.Dump(authHdr)
	buff := make([]byte, ctx.cmnHdr.AuthLen+uint16(authHdr.PaddingLen))
	fmt.Printf("Read %v bytes ", len(buff))
	io.ReadFull(ctx.conn, buff)
}
func handleAuthHeader(ctx *RpcContext) (rpc.RpcComAuthTrailer, []byte) {
	if ctx.cmnHdr.AuthLen != 0 {
		authHdr := rpc.RpcComAuthTrailer{}
		buff := make([]byte, ctx.cmnHdr.AuthLen)
		struc.Unpack(ctx.conn, &authHdr)
		if authHdr.AuthType != 10 {
			fmt.Println("Only NTLMSSP is supported")
		}
		spew.Dump(authHdr)
		io.ReadFull(ctx.conn, buff)
		spew.Dump(buff)
		session, err := ntlm.CreateServerSession(ntlm.Version2, ntlm.ConnectionlessMode)
		if err != nil {
			fmt.Println("ntlm CreateServerSession is returned error\n")
		}
		session.SetUserInfo("ark", "pass@123", "WORKGROUP")
		challenge, err := session.GenerateChallengeMessage()
		spew.Dump(challenge.Bytes())
		return authHdr, challenge.Bytes()
	}
	return rpc.RpcComAuthTrailer{}, nil
}

func handleRpcRequest(ctx *RpcContext) {
	fmt.Println("handleRpcRequest\n")
	reqHdr := rpc.RpcRequestHeader{}
	struc.Unpack(ctx.conn, &reqHdr)
	spew.Dump(reqHdr)
	fmt.Printf("Opnum Received%v\n", reqHdr.OpNum)
	ctx.currentBinding.InvokeOp(reqHdr.OpNum, ctx)
}

func handleRpcBindRequest(ctx *RpcContext) {
	fmt.Println("handleRpcBindRequest\n")
	bindHdr := rpc.RpcBindHdr{}
	struc.Unpack(ctx.conn, &bindHdr)
	var reqUuid rpc.UUID
	//spew.Dump(bindHdr)
	//We have received dump request.
	for _, uuid := range bindHdr.SyntaxNegList.SyntaxList {
		reqUuid = uuid.AbstractSyntax.Id
		break
	}
	fmt.Println("reqUUID", reqUuid)
	var bindingSucceess bool = false
	//Bind if the UUID in the request matches one of the supportedBindings
	for _, interf := range ctx.supportedBindings {
		fmt.Println("availableUUID", interf.GetUuid())
		if interf.GetUuid() == reqUuid {
			ctx.currentBinding = interf
			bindingSucceess = true
		}
	}
	if bindingSucceess == false {
		sendBindNack(ctx, rpc.PSynNegRejAbstractSyntaxNotSupported)
		return
	}
	var authHdr rpc.RpcComAuthTrailer
	var authResponse []byte
	//Successfully bound, but for check for auth headers.
	if ctx.cmnHdr.AuthLen > 0 {
		authHdr, authResponse = handleAuthHeader(ctx)
	}
	if bindingSucceess == true {
		sendBindAck(ctx, &bindHdr, authHdr, authResponse)
		return
	}
}

func sendBindAck(ctx *RpcContext, bindHdr *rpc.RpcBindHdr, authHdr rpc.RpcComAuthTrailer,
	authResponse []byte) {
	fmt.Println("sendBindAck is called")
	cmnHdr := ctx.cmnHdr
	cmnHdr.PType = rpc.PktBindAck
	cmnHdr.AuthLen = 1
	cmnHdr.FragLen = 0
	//Construct BindAckHdr
	bindAckHdr := rpc.RpcBindAckHdr{}
	bindAckHdr.MaxXmitFrag = bindHdr.MaxXmitFrag
	bindAckHdr.MaxRecvFrag = bindHdr.MaxRecvFrag
	bindAckHdr.PortAddress.Address = []byte("135")
	bindAckHdr.PortAddress.Length = uint16(len(bindAckHdr.PortAddress.Address))
	bindAckHdr.SyntaxNegResultList = rpc.RPCSyntaxNegotiationResultList{
		1, 0, 0, []rpc.RPCSyntaxNegotiationResult{{0, 0, bindHdr.SyntaxNegList.SyntaxList[0].AbstractSyntax}}}

	cmnHdrLen, _ := struc.Sizeof(cmnHdr)
	bindAckHdrLen, _ := struc.Sizeof(&bindAckHdr)
	//spew.Dump(cmnHdr)
	buff := new(bytes.Buffer)
	cmnHdr.FragLen = uint16(cmnHdrLen + bindAckHdrLen)
	if authResponse != nil {
		cmnHdr.AuthLen = uint16(len(authResponse))
		cmnHdr.FragLen = uint16(cmnHdrLen+bindAckHdrLen) + cmnHdr.AuthLen + rpc.RpcComAuthTrailerLen
	}
	struc.Pack(buff, &cmnHdr)
	struc.Pack(buff, &bindAckHdr)
	if authResponse != nil {
		struc.Pack(buff, &authHdr)
		buff.Write(authResponse)
	}
	fmt.Println("Debug: bindAckHdr size:", len(buff.Bytes()))
	ctx.conn.Write(buff.Bytes())
}

func sendBindNack(ctx *RpcContext, reason uint16) {
	fmt.Println("sendBindNack is called")
}

func handleRpcAlterContextRequest(ctx *RpcContext) {
	fmt.Println("handleRpcAlterContextRequest\n")
}

func main() {
	fmt.Println("Hello")
	listener, err := net.Listen("tcp", "localhost:135")
	if err != nil {
		fmt.Println("Error on binding ", err)
		return
	}
	fmt.Println("Listening on 10135")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Some error occurred in connection accept")
		}
		go handleConnection(conn)
	}
	wait := make(chan int)
	<-wait
}
