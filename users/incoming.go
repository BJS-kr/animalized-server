package users

import (
	"log/slog"
)

func (u *User) handleIncoming(users *Users) {

	for {
		select {
		case <-u.Stop:
			return
		default:
			if u.Inputs.Len() != 0 {
				continue
			}

			input, err := u.ProduceInput()
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
