package packet

import (
	"errors"
	"io"
)

func (ps *PacketStore) cutChunk() ([]byte, error) {
	chunk, err := ps.inputBuf.ReadBytes(INPUT_PACKET_DELIMITER)

	if err != nil {
		if errors.Is(err, io.EOF) {
			ps.inputBuf.Write(chunk)
		}

		return chunk, err
	}

	return chunk, nil
}
