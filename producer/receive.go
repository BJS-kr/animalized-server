package producer

import (
	"animalized/message"
	"animalized/queue"
)

func Receive(inputProduceChannel <-chan *message.Input, inputs *queue.Queue[*message.Input]) {
	for input := range inputProduceChannel {
		inputs.Enqueue(input)
	}
}
