package handler

import (
	"animalized/message"
	"animalized/producer"
	"animalized/user"
	"bytes"
	"log/slog"
)

func handleIncoming(users *user.Users, u *user.User, buf *[]byte, inputBuf *bytes.Buffer, inputProduceChannel chan<- *message.Input, quit chan<- struct{}) {
	for {
		if err := producer.ProduceInput(u, buf, inputBuf, inputProduceChannel); err != nil {
			slog.Error(err.Error())
			users.RemoveUser(u)
			u.Conn.Close()
			close(quit)
			return
		}
	}
}
