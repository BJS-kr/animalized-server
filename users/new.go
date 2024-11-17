package users

import (
	"animalized/common"
	"animalized/message"
	"animalized/packet"
	"animalized/queue"
	"errors"
	"net"
)

func NewUsers(maxUsers int) *Users {
	us := new(Users)
	us.list = make([]*User, 0, maxUsers)
	us.Max = maxUsers

	return us
}

func NewUser(conn net.Conn, id string, packetStore *packet.PacketStore) (*User, error) {
	userIdLen := len(id)
	if userIdLen == 0 || userIdLen > 10 {
		return nil, errors.New("empty or longer than 10 length id not allowed")
	}
	u := new(User)
	u.Distributable = common.Distributable{
		Inputs: queue.New[*message.Input](),
	}
	u.Id = id
	u.Conn = conn
	u.packetStore = packetStore

	return u, nil

}
