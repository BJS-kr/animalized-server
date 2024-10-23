package handler

import (
	"animalized/consumer"
	"animalized/users"
)

func handleOutgoing(u *users.User, quit <-chan struct{}) {
	for {
		select {
		case <-quit:
			return
		default:
			consumer.Consume(u)
		}
	}
}
