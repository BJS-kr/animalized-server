package handler

import (
	"animalized/message"
	"animalized/packet"
	"animalized/queue"
	"animalized/users"
	"bytes"
	"errors"
	"net"
)

func initialize(conn net.Conn) (*users.User, error) {
	buf, inputBuf := make([]byte, packet.BUFFER_SIZE), bytes.NewBuffer(nil)
	initInput, err := packet.ParseInput(conn, &buf, inputBuf)

	if err != nil {
		return nil, err
	}

	if !packet.IsInitPacket(initInput) {
		return nil, errors.New("init packet type invalid")
	}

	userIdLen := len(initInput.UserId)
	if userIdLen == 0 || userIdLen > 10 {
		return nil, errors.New("empty or longer than 10 length id not allowed")
	}

	u := &users.User{
		Conn:       conn,
		InputQueue: queue.New[*message.Input](),
		Id:         initInput.UserId,
	}

	return u, nil
}
