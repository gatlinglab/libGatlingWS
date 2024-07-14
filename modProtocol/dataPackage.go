package modProtocol

import (
	"bytes"
)

const CPD_VERSION1_HEADLEN = 3
const c_max_datalength = C_P1_MAXDATALEN - CPD_VERSION1_HEADLEN - 1

func MP_PackageDataVersion1(msg []byte) ([]byte, error) {
	var iLen1 uint16 = uint16(len(msg))
	if iLen1 > c_max_datalength {
		//return nil, errors.New("data len > C_P1_MAXDATALEN")
		iLen1 = c_max_datalength
	}
	//fmt.Println("client write msg: ", string(msg))
	//return pInst.wsConn.WriteMessage(websocket.BinaryMessage, msg)
	data := new(bytes.Buffer) // = make([]byte, 0)
	//var datahead = make([]byte, 3)
	//datahead[0] = 0x2
	//datahead[1] = byte(len1 >> 8)
	//datahead[2] = byte(len1)
	data.WriteByte(0x2)
	data.WriteByte(byte(iLen1 >> 8))
	data.WriteByte(byte(iLen1))

	// if iLen1 > C_P1_MAXDATALEN {
	// 	//return nil, errors.New("data len > C_P1_MAXDATALEN")
	// 	iLen1 = C_P1_MAXDATALEN
	// 	data.Write(msg[:C_P1_MAXDATALEN-CPD_VERSION1_HEADLEN])
	// } else {
	// 	data.Write(msg)
	// }
	data.Write(msg[:iLen1])
	//fmt.Println("data last: ", len(data.Bytes()), data.Bytes())
	return data.Bytes(), nil
}
func MP_PackageDataVersion1HeadLen() int {
	return CPD_VERSION1_HEADLEN
}
