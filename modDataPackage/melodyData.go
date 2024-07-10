package modDataPackage

import (
	"fmt"

	"github.com/gatlinglab/libGatlingWSServer/internal/honorMelody"
	"github.com/gatlinglab/libGatlingWSServer/modProtocol"
)

type CWJMelodyDataHandler struct {
	handlerClose   modProtocol.CBWJClosedHandler
	handlerConnect modProtocol.CBWJConnectedHandler
	adapter        *modProtocol.ProtocolAdapter
}

func newMelodyDataHandler() *CWJMelodyDataHandler {
	return &CWJMelodyDataHandler{adapter: modProtocol.NewProtocolAdapter()}
}

func (pInst *CWJMelodyDataHandler) OnClose(session *honorMelody.Session, code int, reason string) error {
	wjSession := session.DataAdapter.(modProtocol.IWJSocket)
	pInst.handlerClose(wjSession)
	return nil
}

func (pInst *CWJMelodyDataHandler) OnConnect(session *honorMelody.Session) {
	fmt.Println("connect: ", session)
	wjSession := NewCWJSessionServer(session)
	session.DataAdapter = wjSession
	pInst.handlerConnect(wjSession)
}

func (pInst *CWJMelodyDataHandler) OnMessage(session *honorMelody.Session, msg []byte) {
	wjSession := session.DataAdapter.(modProtocol.IWJSocket)
	pInst.adapter.OnMessage(wjSession, msg)
}

func (pInst *CWJMelodyDataHandler) OnMessageBinary(session *honorMelody.Session, msg []byte) {
	wjSession := session.DataAdapter.(modProtocol.IWJSocket)
	pInst.adapter.OnMessageBinary(wjSession, msg)
}

func (pInst *CWJMelodyDataHandler) WsHandlerClose(fn modProtocol.CBWJClosedHandler) {
	pInst.handlerClose = fn
}

func (pInst *CWJMelodyDataHandler) WsHandlerConnect(fn modProtocol.CBWJConnectedHandler) {
	pInst.handlerConnect = fn
}

func (pInst *CWJMelodyDataHandler) WsHandlerMessage(fn modProtocol.CBWJMessageHandler) {
	pInst.adapter.WsHandlerMessage(fn)
}

func (pInst *CWJMelodyDataHandler) WsHandlerMessageBinary(fn modProtocol.CBWJMessageBinaryHandler) {
	pInst.adapter.WsHandlerMessageBinary(fn)
}
