package users

import (
	"animalized/packet"
	"log/slog"

	"google.golang.org/protobuf/proto"
)

func (u *User) handleOutgoing() {
	for {
		select {
		case <-u.Stop:
			return
		default:
			if u.Inputs.Len() == 0 {
				continue
			}

			n := u.Inputs.Dequeue()

			message, err := proto.Marshal(n.Value)

			if err != nil {
				slog.Error(err.Error())
				continue
			}

			_, err = u.Conn.Write(append(message, packet.INPUT_PACKET_DELIMITER))

			if err != nil {
				slog.Error(err.Error())
				return
			}
		}
	}
}
