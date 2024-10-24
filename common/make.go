package common

import (
	"animalized/message"
	"animalized/queue"
	"animalized/users"
)

func (b *Base) Make() {
	b.InputChannel = make(chan *message.Input)
	b.Inputs = queue.New[*message.Input]()
	b.Users = new(users.Users)
	b.Stop = make(chan Signal)
}
