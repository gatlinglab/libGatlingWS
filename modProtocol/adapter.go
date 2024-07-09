package modProtocol

import (
	"encoding/binary"
	"fmt"
)

type CBWJConnectedHandler func(*CWJSocket)
type CBWJClosedHandler func(*CWJSocket)
type CBWJMessageHandler func(*CWJSocket, uint32, []byte) // len, data;

const C_P1_MAXDATALEN = 256

type ProtocolAdapter struct {
	msgHandler CBWJMessageHandler
	msgcache   []byte
}

func NewProtocolAdapter() *ProtocolAdapter {
	return &ProtocolAdapter{msgHandler: emptyDefaultMessageHandler, msgcache: make([]byte, C_P1_MAXDATALEN)}
}

func (pInst *ProtocolAdapter) WsHandlerMessage(fn CBWJMessageHandler) {
	pInst.msgHandler = fn
}

func (pInst *ProtocolAdapter) OnMessage(s *CWJSocket, msg []byte) {
	switch msg[0] {
	case 0x2: // version 1;
		pInst.messageVersion1(s, msg)
	case string(rune(33))[0]:
		pInst.messageVersion0(s, msg)

	default:
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
func (pInst *ProtocolAdapter) messageVersion1(s *CWJSocket, msg []byte) {
	len1 := binary.BigEndian.Uint16(msg[1:2])
	if len1 > C_P1_MAXDATALEN {
		fmt.Println("protol1 error: len > C_P1_MAXDATALEN")
		return
	}
	if len(msg) < int(len1)+3 {
		fmt.Println("protol1 error: t < int(len)+3")
		return
	}
	var iLen uint32 = uint32(len1)
	copy(pInst.msgcache[:], msg[3:len1])

	pInst.msgHandler(s, iLen, pInst.msgcache)
}
func (pInst *ProtocolAdapter) messageVersion0(s *CWJSocket, msg []byte) {

	//copy(pInst.msgcache[:], msg[1:])
	iLen := len(msg) - 1

	pInst.msgHandler(s, uint32(iLen), msg[1:]) //pInst.msgcache[:iLen])
}

func emptyDefaultMessageHandler(s *CWJSocket, t uint32, msg []byte) {}
