package handler

import (
	"animalized/message"
	"animalized/packet"
	"animalized/state"
	"animalized/users"
	"bytes"
)

func StartHandlers(users *users.Users, user *users.User, gameState *state.GameState, inputProduceChannel chan<- *message.Input) {
	buf, inputBuf := make([]byte, packet.BUFFER_SIZE), bytes.NewBuffer(nil)

	quit := make(chan struct{})

	go handleIncoming(users, user, &buf, inputBuf, inputProduceChannel, quit)
	go handleOutgoing(user, quit)
}
