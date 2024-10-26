package rooms

import (
	"errors"
)

func (rs *Rooms) Create(roomName string, maxUsers int) (*Room, error) {
	if r, ok := rs.NameMap[RoomName(roomName)]; ok {
		return r, errors.New("room already exists")
	}

	r := new(Room)

	if maxUsers > MAX_USERS_LIMIT {
		return r, errors.New("room users limit has exceeded")
	}

	r.MakeWithUsers(maxUsers)
	r.Status = READY
	rs.NameMap[RoomName(roomName)] = r

	return r, nil
}
