package packet

import (
	"animalized/message"
	"net"
)

func (ps *PacketStore) ParseInput(conn net.Conn) (*message.Input, error) {
	chunk, err := ps.makeChunk(conn)

	if err != nil {
		return nil, err
	}

	stripped, err := stripDelimiter(chunk)

	if err != nil {
		return nil, err
	}

	input := new(message.Input)
	err = into(input, stripped)

	if err != nil {
		return nil, err
	}

	return input, nil
}
