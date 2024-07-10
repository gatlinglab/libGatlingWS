package modClient

import (
	"net/url"

	"github.com/gatlinglab/libGatlingWSServer/modProtocol"
	"github.com/gorilla/websocket"
)

type CGatlingWSClient struct {
	wsUrl          *url.URL
	wsConn         *websocket.Conn
	wsSocket       *CWJSocketClient
	handlerClose   modProtocol.CBWJClosedHandler
	handlerConnect modProtocol.CBWJConnectedHandler
	//handlerMessage modProtocol.CBWJMessageHandler
	adapter *modProtocol.ProtocolAdapter
}

func NewWSClient() *CGatlingWSClient {
	inst := &CGatlingWSClient{adapter: modProtocol.NewProtocolAdapter()}
	return inst
}

func (pInst *CGatlingWSClient) Initialize(serverurl, wsUpgradePath string) error {
	if wsUpgradePath == "" {
		wsUpgradePath = "/ws"
	}
	url1 := url.URL{Scheme: "ws", Host: serverurl, Path: wsUpgradePath}
	pInst.wsUrl = &url1

	return nil
}

func (pInst *CGatlingWSClient) Connect() error {
	connHandler, _, err := websocket.DefaultDialer.Dial(pInst.wsUrl.String(), nil)
	if err != nil {
		return err
	}

	pInst.wsConn = connHandler

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			t, message, err := connHandler.ReadMessage()
			if err != nil {
				pInst.handlerClose(pInst.wsSocket)
				return
			}
			switch t {
			case websocket.TextMessage:
				pInst.adapter.OnMessage(pInst.wsSocket, message)
			case websocket.BinaryMessage:
				pInst.adapter.OnMessageBinary(pInst.wsSocket, message)
			}
		}
	}()

	socket := NewCWJSessionServer(connHandler)
	pInst.wsSocket = socket

	pInst.handlerConnect(socket)

	return nil
}

func (pInst *CGatlingWSClient) WSHandleClosed(fn modProtocol.CBWJClosedHandler) {
	pInst.handlerClose = fn
}

func (pInst *CGatlingWSClient) WSHandleConnected(fn modProtocol.CBWJConnectedHandler) {
	pInst.handlerConnect = fn
}

func (pInst *CGatlingWSClient) WSHandleMessage(fn modProtocol.CBWJMessageHandler) {
	pInst.adapter.WsHandlerMessage(fn)
}

func (pInst *CGatlingWSClient) WSHandleMessageBinary(fn modProtocol.CBWJMessageBinaryHandler) {
	pInst.adapter.WsHandlerMessageBinary(fn)
}
