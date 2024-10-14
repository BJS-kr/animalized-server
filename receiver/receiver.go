package receiver

import "animalized/message"

func Receive(inputProduceChannel <-chan *message.Input) {
	for input := range inputProduceChannel {
		Inputs.Enqueue(input)
	}
}
