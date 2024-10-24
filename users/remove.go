package users

import (
	"slices"
)

func (us *Users) Quit(user *User) int {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	us.users = slices.DeleteFunc(us.users, func(u *User) bool {
		return u.Id == user.Id
	})

	close(user.Stop)

	return len(us.users)
}
