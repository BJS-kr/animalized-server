package packet

import "bytes"

func NewStore() *PacketStore {
	ps := new(PacketStore)
	ps.incomingBuf = make([]byte, BUFFER_SIZE)
	ps.inputBuf = bytes.NewBuffer(nil)

	return ps
}
