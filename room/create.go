package room

import (
	"animalized/users"
	"errors"
)

func (rs *Rooms) Create(roomName string, usersLimit int) (*Room, error) {
	if r, ok := rs.RoomMap[RoomName(roomName)]; ok {
		return r, errors.New("room already exists")
	}

	r := new(Room)

	if usersLimit > 8 {
		return r, errors.New("room users limit has exceeded")
	}

	roomUsers := new(users.Users)
	roomUsers.Max = usersLimit

	r.users = roomUsers
	r.status = READY

	rs.RoomMap[RoomName(roomName)] = r

	return r, nil
}
