package handler

import (
	"animalized/message"
	"animalized/queue"
	"animalized/users"
)

func Propagate(inputQueue *queue.Queue[*message.Input], users *users.Users) {
	for {
		n := inputQueue.Dequeue()

		if n == nil {
			continue
		}

		for u := range users.LockedRange() {
			u.InputQueue.Enqueue(n.Value)
		}
	}
}
