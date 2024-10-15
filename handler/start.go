package handler

import (
	"animalized/message"
	"animalized/user"
	"bytes"
	"log/slog"
	"net"
)

func StartHandlers(users *user.Users, conn net.Conn, inputProduceChannel chan<- *message.Input) {
	defer conn.Close()

	buf, inputBuf := make([]byte, 0), bytes.NewBuffer(nil)
	u, err := initialize(users, conn, &buf, inputBuf)

	if err != nil {
		slog.Error(err.Error())
		users.RemoveUser(u)
		return
	}

	quit := make(chan struct{})

	go handleIncoming(users, u, &buf, inputBuf, inputProduceChannel, quit)
	go handleOutgoing(u, quit)
}
