package iWSServer

import (
	"errors"
	"fmt"
	"net/http"

	//gatlingWS "github.com/gatlinglab/libGatlingWSServer"
	"github.com/gatlinglab/libGatlingWSServer/internal/honorMelody"
)

type CWJWSServer struct {
	serverPort int
	wsServer   *honorMelody.Melody
	router     *cWJWSRouter
	//handlerOnConnect func(CWSSocket)
	//handlerOnClosed  func(CWSSocket)
	//handlerOnMessage func(CWSSocket, []byte)
}

func newWSServer() *CWJWSServer {
	melody := honorMelody.New()
	server := &CWJWSServer{wsServer: melody}
	router := newWJWSRouter(server)
	server.router = router
	return server
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
	listenStr := fmt.Sprintf(":%d", pInst.serverPort)
	err := http.ListenAndServe(listenStr, pInst.router)
	return err
}

func (pInst *CWJWSServer) HttpHandleFunc(pattern string, fn http.HandlerFunc) {
	pInst.router.HandlerFunc(pattern, fn)
}
func (pInst *CWJWSServer) Upgrade(w http.ResponseWriter, r *http.Request) error {
	return nil
}

/*func (pInst *CWJWSServer) WSHandleConnected(fn func(CWSSocket)) {
	pInst.handlerOnConnect = fn
}
func (pInst *CWJWSServer) WSHandleClosed(fn func(CWSSocket)) {
	pInst.handlerOnClosed = fn
}
func (pInst *CWJWSServer) WSHandleMessage(fn func(CWSSocket, []byte)) {
	pInst.handlerOnMessage = fn
}*/
