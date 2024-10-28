package users

import (
	"animalized/packet"

	"bytes"
	"log/slog"
)

func (u *User) handleIncoming(users *Users) {
	buf, inputBuf := make([]byte, packet.BUFFER_SIZE), bytes.NewBuffer(nil)

	for {
		select {
		case <-u.Stop:
			return
		default:
			if u.Inputs.Len() != 0 {
				continue
			}

			input, err := u.ProduceInput(buf, inputBuf)
			if err != nil {
				slog.Error(err.Error())
				users.Quit(u)
				u.Conn.Close()
				return
			}
			u.produceChannel <- input
		}
	}
}
