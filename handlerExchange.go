package libGatlingWSServer

import (
	iWSServer "github.com/gatlinglab/libGatlingWSServer/internal"
)

type cWSExchangeHandler struct {
	handlerConnected CBWJConnectedHandler
}

func newWSExchangeHandler() *cWSExchangeHandler {
	return &cWSExchangeHandler{handlerConnected: emptyOnConnect}
}

func (pInst *cWSExchangeHandler) WSHandlerConnected(socket iWSServer.CWSSocket) {
	pInst.handlerConnected(socket)
}

// //////////////// empty function for default call
func emptyOnConnect(IWJWSSocket) {

}
