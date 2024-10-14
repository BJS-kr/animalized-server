package packet

import (
	"animalized/message"
	"bytes"
	"net"
)

func ParseInput(conn net.Conn, buf *[]byte, inputBuf *bytes.Buffer) (*message.Input, error) {
	chunk, err := makeChunk(conn, buf, inputBuf)

	if err != nil {
		return nil, err
	}

	err = stripDelimiter(chunk)

	if err != nil {
		return nil, err
	}

	input := new(message.Input)
	err = into(input, chunk)

	if err != nil {
		return nil, err
	}

	return input, nil
}
