package main

import (
	"fmt"
	"net/http"

	"github.com/gatlinglab/libGatlingWSServer"
	"github.com/gatlinglab/libGatlingWSServer/modProtocol"
)

func main() {

	pInst := libGatlingWSServer.WWS_NewServer()

	pInst.Initialize(8080)

	pInst.HttpHandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("home")
	})

	pInst.WSHandleConnected(onConnect)
	pInst.WSHandleClosed(onClose)
	pInst.WSHandleMessage(onMessage)

	select {}
}

func onConnect(sock *modProtocol.CWJSocket) {
	fmt.Println("onConnect")
}
func onClose(sock *modProtocol.CWJSocket) {
	fmt.Println("onClose")
}
func onMessage(sock *modProtocol.CWJSocket, len uint32, msg []byte) {
	fmt.Println("onMessage: ", string(msg))
}
