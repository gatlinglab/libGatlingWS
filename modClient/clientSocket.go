package modClient

import (
	"bytes"
	"fmt"

	"github.com/gatlinglab/libGatlingWSServer/modProtocol"
	"github.com/gorilla/websocket"
)

type CWJSocketClient struct {
	wsConn *websocket.Conn
}

func NewCWJSessionServer(conn *websocket.Conn) *CWJSocketClient {
	return &CWJSocketClient{wsConn: conn}
}

func (pInst *CWJSocketClient) Write(msg []byte) error {
	//return pInst.wsConn.WriteMessage(websocket.TextMessage, msg)

	fmt.Println("client write msg: ", string(msg))
	len1 := len(msg)
	data := new(bytes.Buffer) // = make([]byte, 0)
	var datahead = make([]byte, 3)
	datahead[0] = 0x2
	datahead[1] = byte(len1 >> 8)
	datahead[2] = byte(len1)
	data.WriteByte(0x2)
	data.WriteByte(byte(len1 >> 8))
	data.WriteByte(byte(len1))
	data.Write(msg)
	fmt.Println("data last: ", len(data.Bytes()), data.Bytes())

	return pInst.wsConn.WriteMessage(websocket.TextMessage, data.Bytes())
}
func (pInst *CWJSocketClient) WriteBinary(msg []byte) error {
	data1, err := modProtocol.MP_PackageDataVersion1(msg)
	if err != nil {
		return err
	}

	return pInst.wsConn.WriteMessage(websocket.BinaryMessage, data1)
}
