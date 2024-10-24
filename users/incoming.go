package users

import (
	"animalized/message"

	"bytes"
	"log/slog"
)

func handleIncoming(users *Users, u *User, buf *[]byte, inputBuf *bytes.Buffer, inputProduceChannel chan<- *message.Input) {
	for {
		select {
		case <-u.Stop:
			return
		default:
			if err := ProduceInput(u, buf, inputBuf, inputProduceChannel); err != nil {
				slog.Error(err.Error())
				users.Quit(u)
				u.Conn.Close()
				close(u.Stop)
				return
			}
		}
	}
}
