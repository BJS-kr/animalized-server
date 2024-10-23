package room

import (
	// vscode의 go tools에서 문제가 생겨 명칭 변경. 아래의 DeleteFunc에서 user를 os/user로 인식해서 에러를 내는 버그 발생
	"animalized/users"
	"errors"
	"slices"
)

func (rs *Rooms) Quit(roomName string, user *users.User) error {
	r, ok := rs.Rooms[RoomName(roomName)]

	if !ok {
		return errors.New("room does not exists")
	}

	r.Quit(user)

	if len(r.participants) <= 0 {
		delete(rs.Rooms, RoomName(roomName))
	}

	return nil
}

func (r *Room) Quit(user *users.User) {
	r.participants = slices.DeleteFunc(r.participants, func(u *users.User) bool {
		return user == u
	})
}
