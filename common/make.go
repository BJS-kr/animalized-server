package common

import (
	"animalized/message"
	"animalized/queue"
)

func (d *Distributable) Make() {
	d.InputChannel = make(chan *message.Input)
	d.Inputs = queue.New[*message.Input]()
	d.Stop = make(chan Signal)
}
