package modProtocol

import (
	"github.com/gatlinglab/libGatlingWSServer/internal/honorMelody"
)

type CWJSocket struct {
	melodySession *honorMelody.Session
}

func NewCWJSession(melodySession *honorMelody.Session) *CWJSocket {
	return &CWJSocket{melodySession: melodySession}
}

func (pInst *CWJSocket) Write(msg []byte) error {
	return pInst.melodySession.Write(msg)
}
func (pInst *CWJSocket) WriteBinary(msg []byte) error {
	return pInst.melodySession.WriteBinary(msg)
}
