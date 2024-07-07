package libGatlingWSServer

import (
	"net/http"

	"github.com/gatlinglab/libGatlingWSServer/internal/modDataPackage"
)

type IWJWSServer interface {
	Initialize(port int) error
	HttpHandleFunc(pattern string, fn http.HandlerFunc)
	WSHandleConnected(fn modDataPackage.CBWJConnectedHandler)
	WSHandleClosed(fn modDataPackage.CBWJClosedHandler)
	WSHandleMessage(fn modDataPackage.CBWJMessageHandler)
	Start() error
}

func WWS_NewServer(port int) IWJWSServer {
	return modDataPackage.NewWSServer()
}
