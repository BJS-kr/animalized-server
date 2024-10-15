package handler

import (
	"animalized/message"
	"animalized/queue"
)

func Receive(inputQueue *queue.Queue[*message.Input], receiveChannel <-chan *message.Input) {
	for input := range receiveChannel {
		inputQueue.Enqueue(input)
	}
}
