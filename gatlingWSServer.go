package libGatlingWSServer

import (
	"net/http"

	"github.com/gatlinglab/libGatlingWSServer/modDataPackage"
	"github.com/gatlinglab/libGatlingWSServer/modProtocol"
)

type IWJWSServer interface {
	Initialize(port int) error
	HttpHandleFunc(pattern string, fn http.HandlerFunc)
	WSHandleConnected(fn modProtocol.CBWJConnectedHandler)
	WSHandleClosed(fn modProtocol.CBWJClosedHandler)
	WSHandleMessage(fn modProtocol.CBWJMessageHandler)
	Start() error
}

func WWS_NewServer() IWJWSServer {
	return modDataPackage.NewWSServer()
}
