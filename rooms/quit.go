package rooms

import (
	"animalized/users"
	"errors"
)

func (rs *Rooms) Quit(roomName string, user *users.User) error {
	r, ok := rs.NameMap[RoomName(roomName)]

	if !ok {
		return errors.New("room does not exists")
	}

	remain, err := r.Quit(user)

	if err != nil {
		return err
	}

	if remain <= 0 {
		delete(rs.NameMap, RoomName(roomName))
	}

	return nil
}

func (r *Room) Quit(user *users.User) (int, error) {
	return r.Users.Quit(user)
}
