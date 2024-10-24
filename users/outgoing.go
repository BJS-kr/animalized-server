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
			n := u.Inputs.Dequeue()

			if n == nil {
				continue
			}

			message, err := proto.Marshal(n.Value)

			if err != nil {
				slog.Error(err.Error())
				continue
			}

			u.Conn.Write(append(message, packet.INPUT_PACKET_DELIMITER))
		}
	}
}
