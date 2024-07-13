package main

import (
	"fmt"

	"github.com/gatlinglab/libGatlingWS"
	"github.com/gatlinglab/libGatlingWS/modProtocol"
)

var g_socket modProtocol.IWJSocket = nil

func main() {

	pInst := libGatlingWS.WWS_NewClient()

	err := pInst.Initialize("ws", "127.0.0.1:8080", "/ws")
	if err != nil {
		fmt.Println("ws connect error: ", err)
		return
	}

	pInst.WSHandleConnected(onConnect)
	pInst.WSHandleClosed(onClose)
	pInst.WSHandleMessage(onMessage)
	pInst.WSHandleMessageBinary(onMessageBinary)

	err = pInst.Connect()
	if err != nil {
		fmt.Println("http start error: ", err)
	}

	commandRun()
}

func onConnect(sock modProtocol.IWJSocket) {
	fmt.Println("client Connected")
	g_socket = sock
}
func onClose(sock modProtocol.IWJSocket) {
	fmt.Println("client Closed")
}
func onMessage(sock modProtocol.IWJSocket, len1 uint32, msg []byte) {
	fmt.Println("onMessage: ", string(msg))
}

func onMessageBinary(sock modProtocol.IWJSocket, len1 uint32, msg []byte) {
	fmt.Println("onMessage binary: ", string(msg))
}
