package packet

import (
	"bytes"
	"errors"
	"io"
	"net"
)

func makeChunk(conn *net.TCPConn,buf []byte, inputBuf *bytes.Buffer) ([]byte, error) {
		for {
			size, err := readInput(buf, conn)

			if err != nil {
				return buf, err
			}

			if err := writeInput(buf[:size], inputBuf); err != nil {
				return buf, err
			}

			chunk, err := cutChunk(inputBuf)

			if err != nil {
				if errors.Is(err, io.EOF) {
					continue
				}

				return chunk, err
			}

			return chunk, nil
	}
}