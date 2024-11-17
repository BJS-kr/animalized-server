package users

import (
	"animalized/packet"
	"errors"
	"net"
)

func Initialize(conn net.Conn) (*User, error) {
	personalPacketStore := packet.NewStore()
	input, err := personalPacketStore.ParseInput(conn)

	if err != nil {
		return nil, err
	}

	if ok := packet.IsInit(input); !ok {
		return nil, errors.New("invalid init packet type")
	}

	user, err := NewUser(conn, input.UserId, personalPacketStore)

	if err != nil {
		return nil, err
	}

	return user, nil
}
