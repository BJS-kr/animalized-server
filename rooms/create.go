package rooms

import (
	"errors"
)

func (rs *Rooms) Create(roomName string, usersLimit int) (*Room, error) {
	if r, ok := rs.NameMap[RoomName(roomName)]; ok {
		return r, errors.New("room already exists")
	}

	r := new(Room)

	if usersLimit > 8 {
		return r, errors.New("room users limit has exceeded")
	}

	r.MakeWithUsers(usersLimit)
	r.Status = READY
	rs.NameMap[RoomName(roomName)] = r

	return r, nil
}
