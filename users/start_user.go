package users

import (
	"animalized/message"
)

func (u *User) StartPacketHandlers(users *Users, inputProduceChannel chan<- *message.Input) {
	go u.handleIncoming(users, inputProduceChannel)
	go u.handleOutgoing()
}
