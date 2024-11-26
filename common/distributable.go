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

// 해당하는 컨텍스트의 핸들러를 거치지 않는 input
func (d *Distributable) SystemDirectInput(message *message.Input) {
	d.Inputs.Enqueue(message)
}

// 해당하는 컨텍스트의 핸들러를 거치는 input
// ex) Room state를 handle하는 과정은 roomHandler에 있으므로 별도의 로직을 작성해 분산시키지 않고 handler에 보낸다.
func (d *Distributable) SystemInput(message *message.Input) {
	d.InputChannel <- message
}
