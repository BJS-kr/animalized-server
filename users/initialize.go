package users

import (
	"animalized/common"
	"animalized/message"
	"animalized/packet"
	"animalized/queue"
	"bytes"
	"errors"
	"net"
)

func Initialize(conn net.Conn) (*User, error) {
	buf, inputBuf := make([]byte, packet.BUFFER_SIZE), bytes.NewBuffer(nil)
	input, err := packet.ParseInput(conn, buf, inputBuf)

	if err != nil {
		return nil, err
	}

	if ok := packet.IsInit(input); !ok {
		return nil, errors.New("invalid init packet type")
	}

	userIdLen := len(input.UserId)

	if userIdLen == 0 || userIdLen > 10 {
		return nil, errors.New("empty or longer than 10 length id not allowed")
	}

	u := &User{
		Conn: conn,
		Id:   input.UserId,
		Distributable: common.Distributable{
			Inputs: queue.New[*message.Input](),
		},
	}

	return u, nil
}
