package packet

import (
	"errors"
	"io"
	"net"
)

func (ps *PacketStore) makeChunk(conn net.Conn) ([]byte, error) {
	for {
		chunk, err := ps.cutChunk()

		if err == nil {
			return chunk, nil
		}

		if !errors.Is(err, io.EOF) {
			return chunk, err
		}

		size, err := ps.readInput(conn)

		if err != nil {
			return ps.incomingBuf, err
		}

		if err := ps.writeInput(size); err != nil {
			return ps.incomingBuf, err
		}
	}
}
