package handler

import (
	"animalized/message"
	"animalized/producer"
	"animalized/users"
	"bytes"
	"log/slog"
)

func handleIncoming(users *users.Users, u *users.User, buf *[]byte, inputBuf *bytes.Buffer, inputProduceChannel chan<- *message.Input) {
	for {
		select {
		case <-u.Quit:
			return
		default:
			if err := producer.ProduceInput(u, buf, inputBuf, inputProduceChannel); err != nil {
				slog.Error(err.Error())
				users.RemoveUser(u)
				u.Conn.Close()
				close(u.Quit)
				return
			}
		}

	}
}
