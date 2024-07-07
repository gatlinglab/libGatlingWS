package libGatlingWSServer

import (
	"net/http"

	iWSServer "github.com/gatlinglab/libGatlingWSServer/internal"
)

type IWJWSSocket interface {
	Write(msg []byte) error
	WriteBinary(msg []byte) error
	Close() error
	IsClosed() bool
	LocalAddr()
	RemoteAddr()
}

type CBWJConnectedHandler func(IWJWSSocket)
type CBWJClosedHandler func(IWJWSSocket)
type CBWJMessageHandler func(IWJWSSocket, int, []byte) // len, data;

type IWJWSServer interface {
	Initialize(port int) error
	HttpHandleFunc(pattern string, fn http.HandlerFunc)
	Start() error
	WSHandleConnected(fn CBWJConnectedHandler)
	WSHandleClosed(fn CBWJConnectedHandler)
	WSHandleMessage(fn CBWJMessageHandler)
}

//var g_singleWSServer IWJWSServer = nil

func WWS_NewServer(port int) IWJWSServer {
	return iWSServer.IWS_NewServer(port)
}

// func WWS_NewDefaultServer(port int) IWJWSServer {
// 	inst := iWSServer.IWS_NewServer(port)
// 	g_singleWSServer = inst
// 	return inst
// }
// func WWS_SetDefaultServer(inst IWJWSServer) {
// 	g_singleWSServer = inst
// }

// func WWS_GetDefaultServer() IWJWSServer {
// 	return g_singleWSServer
// }
