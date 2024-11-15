package rooms

import (
	"errors"
)

func (rs *Rooms) Create(roomName string, maxUsers int) (*Room, error) {
	if roomName == "" {
		return nil, errors.New("room name not provided when creating room")
	}

	if maxUsers <= 0 || maxUsers > 8 {
		return nil, errors.New("max users not in valid range")
	}

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
