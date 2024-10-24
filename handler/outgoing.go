package handler

import (
	"animalized/consumer"
	"animalized/users"
)

func handleOutgoing(u *users.User) {
	for {
		select {
		case <-u.Stop:
			return
		default:
			consumer.Consume(u)
		}
	}
}
