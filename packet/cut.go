package packet

import (
	"bytes"
	"errors"
	"io"
)

func cutChunk(inputBuf *bytes.Buffer) (*[]byte, error) {
	chunk, err := inputBuf.ReadBytes(INPUT_PACKET_DELIMITER)

	if err != nil {
		if errors.Is(err, io.EOF) {
			inputBuf.Write(chunk)
			return &chunk, err
		}

		return &chunk, err
	}

	return &chunk, nil
}
