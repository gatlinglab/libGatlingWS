package libGatlingWS

import (
	"github.com/gatlinglab/libGatlingWS/modClient"
	"github.com/gatlinglab/libGatlingWS/modProtocol"
)

type IWJWSClient interface {
	Initialize(serverurl, wsUpgradePath string) error /// format ip:port, address:port; path: /ws
	WSHandleConnected(fn modProtocol.CBWJConnectedHandler)
	WSHandleClosed(fn modProtocol.CBWJClosedHandler)
	WSHandleMessage(fn modProtocol.CBWJMessageHandler)
	WSHandleMessageBinary(fn modProtocol.CBWJMessageBinaryHandler)
	Connect() error
}

func WWS_NewClient() IWJWSClient {
	return modClient.NewWSClient()
}
