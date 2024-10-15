package handler

import (
	"animalized/message"
	"animalized/queue"
	"animalized/user"
)

func Propagate(inputQueue *queue.Queue[*message.Input], users *user.Users) {
	for {
		n := inputQueue.Dequeue()

		if n == nil {
			return
		}

		for u := range users.LockedRange() {
			u.InputQueue.Enqueue(n.Value)
		}
	}
}
