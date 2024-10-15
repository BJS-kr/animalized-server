package handler

import (
	"animalized/message"
	"animalized/packet"
	"animalized/queue"
	"animalized/user"
	"bytes"
	"errors"
	"net"
)

func initialize(conn net.Conn, buf *[]byte, inputBuf *bytes.Buffer) (*user.User, error) {
	initInput, err := packet.ParseInput(conn, buf, inputBuf)

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

	u := &user.User{
		Conn:       conn,
		InputQueue: queue.New[*message.Input](),
		Id:         initInput.UserId,
	}

	users.InsertUser(u)

	return u, nil
}
