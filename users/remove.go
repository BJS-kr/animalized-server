package users

import (
	"slices"
)

func (us *Users) RemoveUser(u *User) int {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	us.users = slices.DeleteFunc(us.users, func(eu *User) bool {
		return eu.Id == u.Id
	})

	return len(us.users)
}
