package modDataPackage

import (
	"github.com/gatlinglab/libGatlingWS/internal/honorMelody"
	"github.com/gatlinglab/libGatlingWS/modProtocol"
)

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
	data1, err := modProtocol.MP_PackageDataVersion1(msg)
	if err != nil {
		return err
	}
	// fmt.Println("server write binary msg: ", string(msg))
	// len1 := len(msg)
	// data := new(bytes.Buffer)
	// data.WriteByte(0x2)
	// data.WriteByte(byte(len1 >> 8))
	// data.WriteByte(byte(len1))
	// data.Write(msg)
	// fmt.Println("server data last: ", len(data.Bytes()), data.Bytes())

	return pInst.melodySession.WriteBinary(data1)
}
