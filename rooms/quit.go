package rooms

import (
	"animalized/users"
	"errors"
)

func (rs *Rooms) Quit(roomName string, userName string) (*users.User, error) {
	r, ok := rs.NameMap[RoomName(roomName)]

	if !ok {
		return nil, errors.New("room does not exists")
	}

	user, err := r.Users.FindUserById(userName)

	if err != nil {
		return nil, err
	}

	remain, err := r.Quit(user)

	if err != nil {
		return nil, err
	}

	if remain <= 0 {
		delete(rs.NameMap, RoomName(roomName))
	}

	return user, nil
}

func (r *Room) Quit(user *users.User) (int, error) {
	return r.Users.Quit(user)
}
