package packet

import (
	"animalized/message"

	"animalized/queue"
	"animalized/users"
	"bytes"
	"errors"
	"net"
)

func Initialize(conn net.Conn) (*users.User, error) {
	buf, inputBuf := make([]byte, BUFFER_SIZE), bytes.NewBuffer(nil)
	initInput, err := ParseInput(conn, &buf, inputBuf)

	if err != nil {
		return nil, err
	}

	if !IsInitPacket(initInput) {
		return nil, errors.New("init packet type invalid")
	}

	userIdLen := len(initInput.UserId)
	if userIdLen == 0 || userIdLen > 10 {
		return nil, errors.New("empty or longer than 10 length id not allowed")
	}

	u := &users.User{
		Conn:   conn,
		Inputs: queue.New[*message.Input](),
		Id:     initInput.UserId,
	}

	return u, nil
}
