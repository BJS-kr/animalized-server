package packet

import (
	"bytes"
	"errors"
	"io"
	"net"
)

func makeChunk(conn net.Conn, buf *[]byte, inputBuf *bytes.Buffer) (*[]byte, error) {
	for {
		chunk, err := cutChunk(inputBuf)

		if err == nil {
			return chunk, nil
		}

		if !errors.Is(err, io.EOF) {
			return chunk, err
		}

		size, err := readInput(buf, conn)

		if err != nil {
			return buf, err
		}

		if err := writeInput((*buf)[:size], inputBuf); err != nil {
			return buf, err
		}
	}
}
