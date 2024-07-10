package main

import (
	"fmt"
	"net/http"

	"github.com/gatlinglab/libGatlingWSServer"
	"github.com/gatlinglab/libGatlingWSServer/modProtocol"
)

func main() {

	pInst := libGatlingWSServer.WWS_NewServer()

	err := pInst.Initialize(8080)
	if err != nil {
		fmt.Println("ws init error: ", err)
	}

	pInst.HttpHandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.html")
	})

	pInst.WSHandleConnected(onConnect)
	pInst.WSHandleClosed(onClose)
	pInst.WSHandleMessage(onMessage)
	pInst.WSHandleMessageBinary(onMessageBinary)

	err = pInst.Start()
	if err != nil {
		fmt.Println("http start error: ", err)
	}
}

func onConnect(sock modProtocol.IWJSocket) {
	fmt.Println("onConnect")
}
func onClose(sock modProtocol.IWJSocket) {
	fmt.Println("onClose")
}
func onMessage(sock modProtocol.IWJSocket, len1 uint32, msg []byte) {
	fmt.Println("onMessage: ", string(msg), "!!!")
	reply := "nice: " + string(msg)
	sock.Write([]byte(reply))
}

func onMessageBinary(sock modProtocol.IWJSocket, len1 uint32, msg []byte) {
	data1 := msg[:len1]
	fmt.Println("onMessageBinary: ", string(data1))
	reply := "nice: " + string(data1)
	sock.WriteBinary([]byte(reply))
}
