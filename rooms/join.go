package rooms

import (
	"animalized/users"
	"errors"
)

func (rs *Rooms) Join(roomName string, user *users.User) (*Room, error) {
	r, ok := rs.NameMap[RoomName(roomName)]

	if !ok {
		return nil, errors.New("room not exists")
	}

	if err := r.Join(user); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Room) Join(user *users.User) error {
	if err := r.Users.Join(user, r.InputChannel); err != nil {
		return err
	}

	return nil
}
