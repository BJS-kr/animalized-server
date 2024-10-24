package room

import (
	"animalized/users"
	"errors"
)

func (rs *Rooms) Join(roomName string, user *users.User) error {
	r, ok := rs.NameMap[RoomName(roomName)]

	if !ok {
		return errors.New("room not exists")
	}

	if err := r.Join(user); err != nil {
		return err
	}

	return nil
}

func (r *Room) Join(user *users.User) error {
	if err := r.Users.InsertUser(user); err != nil {
		return err
	}

	users.StartHandlers(r.Users, user, r.InputChannel)

	return nil
}
