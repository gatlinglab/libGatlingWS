package modProtocol

import (
	"bytes"

	"github.com/gatlinglab/libGatlingWSServer/internal/honorMelody"
)

type IWJSocket interface {
	Write(msg []byte) error
	WriteBinary(msg []byte) error
}

type CWJSocketServer struct {
	melodySession *honorMelody.Session
}

func NewCWJSessionServer(melodySession *honorMelody.Session) *CWJSocketServer {
	return &CWJSocketServer{melodySession: melodySession}
}

func (pInst *CWJSocketServer) Write(msg []byte) error {
	return pInst.melodySession.Write(msg)
}
func (pInst *CWJSocketServer) WriteBinary(msg []byte) error {
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

	return pInst.melodySession.WriteBinary(data.Bytes())
}
