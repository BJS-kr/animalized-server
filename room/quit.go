package room

import (
	"animalized/users"
	"errors"
)

func (rs *Rooms) Quit(roomName string, user *users.User) error {
	r, ok := rs.NameMap[RoomName(roomName)]

	if !ok {
		return errors.New("room does not exists")
	}

	remain := r.Quit(user)

	if remain <= 0 {
		delete(rs.NameMap, RoomName(roomName))
	}

	return nil
}

func (r *Room) Quit(user *users.User) int {
	return r.Users.Quit(user)
}
