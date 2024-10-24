package users

import (
	"animalized/packet"

	"google.golang.org/protobuf/proto"
)

func Consume(u *User) error {
	n := u.Inputs.Dequeue()

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
