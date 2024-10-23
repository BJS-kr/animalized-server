package handler

import (
	"animalized/message"
	"animalized/packet"
	"animalized/users"
	"bytes"
)

func StartHandlers(users *users.Users, user *users.User, inputProduceChannel chan<- *message.Input) {
	buf, inputBuf := make([]byte, packet.BUFFER_SIZE), bytes.NewBuffer(nil)

	go handleIncoming(users, user, &buf, inputBuf, inputProduceChannel)
	go handleOutgoing(user)
}
