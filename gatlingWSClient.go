package libGatlingWSServer

import (
	"github.com/gatlinglab/libGatlingWSServer/modClient"
	"github.com/gatlinglab/libGatlingWSServer/modProtocol"
)

type IWJWSClient interface {
	Initialize(serverurl, wsUpgradePath string) error /// format ip:port, address:port; path: /ws
	WSHandleConnected(fn modProtocol.CBWJConnectedHandler)
	WSHandleClosed(fn modProtocol.CBWJClosedHandler)
	WSHandleMessage(fn modProtocol.CBWJMessageHandler)
	Connect() error
}

func WWS_NewClient() IWJWSClient {
	return modClient.NewWSClient()
}
