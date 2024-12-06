package common

import (
	"animalized/message"
	"animalized/queue"
	"log/slog"
)

type Actor struct {
	Inputs       *queue.Queue[*message.Input]
	InputChannel chan *message.Input
	Stop         chan Signal
}

func (a *Actor) Make() {
	a.InputChannel = make(chan *message.Input, 100)
	a.Inputs = queue.New[*message.Input]()
	a.Stop = make(chan Signal)
}

func (a *Actor) Receive(handler Handler) {
	for input := range a.InputChannel {
		input, err := handler(input)

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		if input == nil {
			continue
		}

		a.Inputs.Enqueue(input)

	}
}

// 해당하는 컨텍스트의 핸들러를 거치지 않는 input
func (a *Actor) SystemDirectInput(message *message.Input) {
	a.Inputs.Enqueue(message)
}

// 해당하는 컨텍스트의 핸들러를 거치는 input
// ex) Room state를 handle하는 과정은 roomHandler에 있으므로 별도의 로직을 작성해 분산시키지 않고 handler에 보낸다.
func (a *Actor) SystemInput(message *message.Input) {
	a.InputChannel <- message
}
