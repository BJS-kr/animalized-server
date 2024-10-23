package room

import (
	"animalized/users"
	"errors"
)

func (rs *Rooms) Quit(roomName string, user *users.User) error {
	r, ok := rs.Rooms[RoomName(roomName)]

	if !ok {
		return errors.New("room does not exists")
	}

	remain := r.Quit(user)

	if remain <= 0 {
		delete(rs.Rooms, RoomName(roomName))
	}

	return nil
}

func (r *Room) Quit(user *users.User) int {
	return r.users.RemoveUser(user)
}
