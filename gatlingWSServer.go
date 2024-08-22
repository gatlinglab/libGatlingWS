package libGatlingWS

import (
	"net/http"

	"github.com/gatlinglab/libGatlingWS/modDataPackage"
	"github.com/gatlinglab/libGatlingWS/modProtocol"
)

type IWJWSServer interface {
	Initialize(port int) error
	HttpHandleFunc(pattern string, fn http.HandlerFunc)
	WSHandleConnected(fn modProtocol.CBWJConnectedHandler)
	WSHandleClosed(fn modProtocol.CBWJClosedHandler)
	WSHandleMessage(fn modProtocol.CBWJMessageHandler)
	WSHandleMessageBinary(fn modProtocol.CBWJMessageBinaryHandler)
	Start() error
	Stop()
}

func WWS_NewServer() IWJWSServer {
	return modDataPackage.NewWSServer()
}
