package handler

import (
	"animalized/consumer"
	"animalized/users"
)

func handleOutgoing(u *users.User) {
	for {
		select {
		case <-u.Quit:
			return
		default:
			consumer.Consume(u)
		}
	}
}
