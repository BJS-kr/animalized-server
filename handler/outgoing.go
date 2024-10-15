package handler

import (
	"animalized/consumer"
	"animalized/user"
)

func handleOutgoing(u *user.User, quit <-chan struct{}) {
	for {
		select {
		case <-quit:
			return
		default:
			consumer.Consume(u)
		}
	}
}
