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

func (pInst *ProtocolAdapter) OnMessage(s *CWJSocket, t int, msg []byte) {
	switch msg[0] {
	case 0x2: // version 1;
		pInst.messageVersion1(s, t, msg)
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
func (pInst *ProtocolAdapter) messageVersion1(s *CWJSocket, t int, msg []byte) {
	len := binary.BigEndian.Uint16(msg[1:2])
	if len > C_P1_MAXDATALEN {
		fmt.Println("protol1 error: len > C_P1_MAXDATALEN")
		return
	}
	if t < int(len)+3 {
		fmt.Println("protol1 error: t < int(len)+3")
		return
	}
	var iLen uint32 = uint32(len)
	copy(pInst.msgcache[:], msg[3:len])

	pInst.msgHandler(s, iLen, pInst.msgcache)
}

func emptyDefaultMessageHandler(s *CWJSocket, t uint32, msg []byte) {}
