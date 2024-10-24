package users

import (
	"animalized/message"
	"animalized/packet"

	"bytes"
	"log/slog"
)

func (u *User) handleIncoming(users *Users, inputProduceChannel chan<- *message.Input) {
	buf, inputBuf := make([]byte, packet.BUFFER_SIZE), bytes.NewBuffer(nil)

	for {
		select {
		case <-u.Stop:
			return
		default:
			if err := u.ProduceInput(&buf, inputBuf, inputProduceChannel); err != nil {
				slog.Error(err.Error())
				users.Quit(u)
				u.Conn.Close()
				return
			}
		}
	}
}
