package common

import (
	"animalized/message"
	"animalized/queue"
)

func (b *Base) Make() {
	b.InputChannel = make(chan *message.Input)
	b.Inputs = queue.New[*message.Input]()
	b.Stop = make(chan Signal)
}
