package consumer

import (
	"animalized/packet"
	"animalized/users"

	"google.golang.org/protobuf/proto"
)

func Consume(u *users.User) error {
	n := u.InputQueue.Dequeue()

	if n == nil {
		return nil
	}

	message, err := proto.Marshal(n.Value)

	if err != nil {
		return err
	}

	u.Conn.Write(append(message, packet.INPUT_PACKET_DELIMITER))

	return nil
}
