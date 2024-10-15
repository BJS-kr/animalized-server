package handler

import (
	"animalized/message"
	"animalized/packet"
	"animalized/state"
	"animalized/user"
	"bytes"
	"log/slog"
	"net"
)

func StartHandlers(users *user.Users, serverState *state.ServerState, conn net.Conn, inputProduceChannel chan<- *message.Input) {
	buf, inputBuf := make([]byte, packet.BUFFER_SIZE), bytes.NewBuffer(nil)
	u, err := initialize(users, conn, &buf, inputBuf)

	if err != nil {
		slog.Error(err.Error())
		return
	}

	err = serverState.AddUserState(state.UserID(u.Id))

	if err != nil {
		slog.Error(err.Error())
		return
	}

	users.InsertUser(u)

	quit := make(chan struct{})

	go handleIncoming(users, u, &buf, inputBuf, inputProduceChannel, quit)
	go handleOutgoing(u, quit)
}
