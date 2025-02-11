package modDataPackage

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gatlinglab/libGatlingWS/internal/honorMelody"
	"github.com/gatlinglab/libGatlingWS/modProtocol"
)

type CGatlingWSServer struct {
	wsServer    *honorMelody.Melody
	serverPort  int
	router      *cWJWSRouter
	dataReciver *CWJMelodyDataHandler
}

func NewWSServer() *CGatlingWSServer {
	melody := honorMelody.New()
	server := &CGatlingWSServer{wsServer: melody}
	reciver := newMelodyDataHandler()
	router := newWJWSRouter(server)
	server.router = router
	server.dataReciver = reciver
	return server
}

func (pInst *CGatlingWSServer) Initialize(port int) error {
	if port < 10 || port > 65535 {
		return errors.New("port error")
	}
	m := honorMelody.New()
	pInst.wsServer = m
	pInst.serverPort = port

	pInst.wsServer.HandleConnect(pInst.dataReciver.OnConnect)
	pInst.wsServer.HandleClose(pInst.dataReciver.OnClose)
	pInst.wsServer.HandleMessage(pInst.dataReciver.OnMessage)
	pInst.wsServer.HandleMessageBinary(pInst.dataReciver.OnMessageBinary)

	return nil
}

func (pInst *CGatlingWSServer) Start() error {
	listenStr := fmt.Sprintf(":%d", pInst.serverPort)
	fmt.Println("ws server start listen: ", listenStr)
	err := http.ListenAndServe(listenStr, pInst.router)
	return err
}
func (pInst *CGatlingWSServer) Stop() {
	// can not stop now;
}

func (pInst *CGatlingWSServer) HttpHandleFunc(pattern string, fn http.HandlerFunc) {
	pInst.router.HandlerFunc(pattern, fn)
}

func (pInst *CGatlingWSServer) Upgrade(w http.ResponseWriter, r *http.Request) error {
	return pInst.wsServer.HandleRequest(w, r)
}

func (pInst *CGatlingWSServer) WSHandleClosed(fn modProtocol.CBWJClosedHandler) {
	pInst.dataReciver.WsHandlerClose(fn)
}

func (pInst *CGatlingWSServer) WSHandleConnected(fn modProtocol.CBWJConnectedHandler) {
	pInst.dataReciver.WsHandlerConnect(fn)
}

func (pInst *CGatlingWSServer) WSHandleMessage(fn modProtocol.CBWJMessageHandler) {
	pInst.dataReciver.WsHandlerMessage(fn)
}

func (pInst *CGatlingWSServer) WSHandleMessageBinary(fn modProtocol.CBWJMessageBinaryHandler) {
	pInst.dataReciver.WsHandlerMessageBinary(fn)
}
