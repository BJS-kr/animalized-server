package common

import (
	"animalized/message"
	"animalized/queue"
	"log/slog"
)

type Distributable struct {
	Inputs       *queue.Queue[*message.Input]
	InputChannel chan *message.Input
	Stop         chan Signal
}

func (d *Distributable) Make() {
	d.InputChannel = make(chan *message.Input, 100)
	d.Inputs = queue.New[*message.Input]()
	d.Stop = make(chan Signal)
}

func (d *Distributable) Receive(handler Handler) {
	for input := range d.InputChannel {
		select {
		case <-d.Stop:
			return
		default:
			input, err := handler(input)

			if err != nil {
				slog.Error(err.Error())
				continue
			}

			if input == nil {
				continue
			}

			d.Inputs.Enqueue(input)
		}
	}
}

func (d *Distributable) SystemInput(message *message.Input) {
	d.Inputs.Enqueue(message)
}
