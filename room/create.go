package room

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

	r.Make()
	r.Users.Max = usersLimit
	r.status = READY
	rs.NameMap[RoomName(roomName)] = r

	go r.Receive(r.handler)
	go r.Distribute()

	return r, nil
}
