package users

import (
	"animalized/message"
	"animalized/packet"
	"bytes"
)

func StartHandlers(users *Users, user *User, inputProduceChannel chan<- *message.Input) {
	buf, inputBuf := make([]byte, packet.BUFFER_SIZE), bytes.NewBuffer(nil)

	go handleIncoming(users, user, &buf, inputBuf, inputProduceChannel)
	go handleOutgoing(user)
}
