package packet

import "bytes"

type PacketStore struct {
	incomingBuf []byte
	inputBuf    *bytes.Buffer
}
