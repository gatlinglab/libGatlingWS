package modDataPackage

import (
	"github.com/gatlinglab/libGatlingWSServer/internal/honorMelody"
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

type CWJMelodyDataHandler struct {
	handlerClose   CBWJClosedHandler
	handlerConnect CBWJConnectedHandler
	handlerMessage CBWJMessageHandler
}

func newMelodyDataHandler() *CWJMelodyDataHandler {
	return &CWJMelodyDataHandler{}
}

func (pInst *CWJMelodyDataHandler) OnClose(session *honorMelody.Session, code int, reason string) error {
	//pInst.handlerClose(session)
	return nil
}

func (pInst *CWJMelodyDataHandler) OnConnect(session *honorMelody.Session) {
	//.Write(data)
}

func (pInst *CWJMelodyDataHandler) OnMessage(session *honorMelody.Session, msg []byte) {

}

func (pInst *CWJMelodyDataHandler) WsHandlerClose(fn CBWJClosedHandler) {
	pInst.handlerClose = fn
}

func (pInst *CWJMelodyDataHandler) WsHandlerConnect(fn CBWJConnectedHandler) {
	pInst.handlerConnect = fn
}

func (pInst *CWJMelodyDataHandler) WsHandlerMessage(fn CBWJMessageHandler) {
	pInst.handlerMessage = fn
}
