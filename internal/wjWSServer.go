package iWSServer

import (
	"errors"
	"net/http"

	//gatlingWS "github.com/gatlinglab/libGatlingWSServer"
	"github.com/gatlinglab/libGatlingWSServer/internel/honorMelody"
)

type CWJWSServer struct {
	serverPort       int
	wsServer         *honorMelody.Melody
	router           *cWJWSRouter
	handlerOnConnect func(CWSSocket)
	handlerOnClosed  func(CWSSocket)
	handlerOnMessage func(CWSSocket, []byte)
}

func newWSServer() *CWJWSServer {
	melody := honorMelody.New()
	router := newWJWSRouter()
	return &CWJWSServer{wsServer: melody, router: router}
}

func (pInst *CWJWSServer) Initialize(port int) error {
	if port < 10 || port > 65535 {
		return errors.New("port error")
	}
	m := honorMelody.New()
	pInst.wsServer = m

	return nil
}

func (pInst *CWJWSServer) Start() error {

	return nil
}

func (pInst *CWJWSServer) HttpHandleFunc(pattern string, fn http.HandlerFunc) {
	pInst.router.HandlerFunc(pattern, fn)
}
func (pInst *CWJWSServer) WSHandleConnected(fn func(CWSSocket)) {
	pInst.handlerOnConnect = fn
}
func (pInst *CWJWSServer) WSHandleClosed(fn func(CWSSocket)) {
	pInst.handlerOnClosed = fn
}
func (pInst *CWJWSServer) WSHandleMessage(fn func(CWSSocket, []byte)) {
	pInst.handlerOnMessage = fn
}
