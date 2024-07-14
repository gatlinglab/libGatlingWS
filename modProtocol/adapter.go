package modProtocol

import (
	"fmt"
)

const C_P1_MAXDATALEN = 1024

type IWJSocket interface {
	Write(msg []byte) error
	WriteBinary(msg []byte) error
	PutSocketData(interface{}) // save custom data to socket for next time using;
	GetSocketData() interface{}
}
type CBWJConnectedHandler func(IWJSocket)
type CBWJClosedHandler func(IWJSocket)
type CBWJMessageHandler func(IWJSocket, uint32, []byte)       // len, data;
type CBWJMessageBinaryHandler func(IWJSocket, uint32, []byte) // len, data;

type ProtocolAdapter struct {
	msgHandler       CBWJMessageHandler
	msgBinaryHandler CBWJMessageBinaryHandler
	msgcache         []byte
}

func NewProtocolAdapter() *ProtocolAdapter {
	return &ProtocolAdapter{msgHandler: emptyDefaultMessageHandler,
		msgBinaryHandler: emptyDefaultMessageHandler,
		msgcache:         make([]byte, C_P1_MAXDATALEN)}
}

func (pInst *ProtocolAdapter) WsHandlerMessage(fn CBWJMessageHandler) {
	pInst.msgHandler = fn
}

func (pInst *ProtocolAdapter) WsHandlerMessageBinary(fn CBWJMessageBinaryHandler) {
	pInst.msgBinaryHandler = fn
}

func (pInst *ProtocolAdapter) OnMessage(s IWJSocket, msg []byte) {
	//fmt.Println("server protocol get message", len(msg))
	switch msg[0] {
	case 33: //string(rune(33))[0]:
		pInst.messageVersion0(s, msg)

	default:
		fmt.Println("protol error: ", msg, len(msg))
		//error

	}
}

func (pInst *ProtocolAdapter) OnMessageBinary(s IWJSocket, msg []byte) {
	//fmt.Println("server protocol get message", len(msg))
	switch msg[0] {
	case 0x2: // version 1;
		pInst.messageVersion1(s, msg)

	default:
		fmt.Println("protol error: ", msg, len(msg))
		//error

	}
}

/*
numBytes := []byte{0xFF, 0x10}

	    u := binary.BigEndian.Uint16(numBytes)
	    fmt.Printf("%#X %[1]v\n", u) // 0XFF10 65296
	}

and see inside binary.BigEndian.Uint16(b []byte):

	func (bigEndian) Uint16(b []byte) uint16 {
	    _ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
	    return uint16(b[1]) | uint16(b[0])<<8
	}
*/
func (pInst *ProtocolAdapter) messageVersion1(s IWJSocket, msg []byte) {
	//fmt.Println("binary protol1", string(msg), len(msg))
	var len1 int16 = int16(msg[1])<<8 | int16(msg[2])
	//fmt.Println("binary protol1 data len: ", pLen)
	//var lendata []byte
	//lendata = append(lendata, msg[1:3]...)
	//len1 := pLen //binary.BigEndian.Uint16(lendata)
	if len1 > C_P1_MAXDATALEN {
		fmt.Println("protol1 error: len > C_P1_MAXDATALEN", len1, C_P1_MAXDATALEN)
		return
	}
	//fmt.Println("binary protol1 -2", string(msg))
	if len(msg) < int(len1)+3 {
		fmt.Println("protol1 error: t < int(len)+3")
		return
	}
	//fmt.Println("binary protol1-3", string(msg))
	var iLen uint32 = uint32(len1)
	copy(pInst.msgcache[:], msg[3:3+len1])

	//fmt.Println("binary protol1-4", string(pInst.msgcache))

	pInst.msgBinaryHandler(s, iLen, pInst.msgcache)
}

func (pInst *ProtocolAdapter) messageVersion0(s IWJSocket, msg []byte) {

	//copy(pInst.msgcache[:], msg[1:])
	iLen := len(msg) - 1

	pInst.msgHandler(s, uint32(iLen), msg[1:]) //pInst.msgcache[:iLen])
}

func emptyDefaultMessageHandler(s IWJSocket, t uint32, msg []byte) {}
