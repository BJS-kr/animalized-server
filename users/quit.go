package users

import (
	"slices"
)

func (us *Users) Quit(user *User) int {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	us.list = slices.DeleteFunc(us.list, func(u *User) bool {
		return u.Id == user.Id
	})

	user.StopPacketHandlers()

	return len(us.list)
}
