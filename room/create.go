package room

import (
	"animalized/users"
	"errors"
)

func (rs *Rooms) Create(roomName string, usersLimit int) error {
	r := new(Room)

	if usersLimit > 8 {
		return errors.New("room users limit has exceeded")
	}

	roomUsers := new(users.Users)
	roomUsers.Max = usersLimit

	r.users = roomUsers
	r.status = READY

	rs.Rooms[RoomName(roomName)] = r

	return nil
}
