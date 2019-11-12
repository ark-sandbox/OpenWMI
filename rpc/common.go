package rpc

import (
	"net"
)

type RpcInterface interface {
	GetUuid() UUID
	InvokeOp(opnum uint16, ctx *RpcContext) error
}

type RpcContext struct {
	cmnHdr            RpcCommonHdr
	currentBinding    RpcInterface
	conn              net.Conn
	supportedBindings []RpcInterface
}
