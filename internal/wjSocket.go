package iWSServer

type CWSSocket struct {
}

func (pInst *CWSSocket) Write(msg []byte) error {
	return nil
}
func (pInst *CWSSocket) WriteBinary(msg []byte) error {
	return nil

}
func (pInst *CWSSocket) Close() error {
	return nil

}
func (pInst *CWSSocket) IsClosed() bool {
	return true
}
func (pInst *CWSSocket) LocalAddr() {

}
func (pInst *CWSSocket) RemoteAddr() {

}
