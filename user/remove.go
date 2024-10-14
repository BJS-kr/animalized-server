package user

import "slices"

func (us *Users) RemoveUser(u *User) {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	us.users = slices.DeleteFunc(us.users, func(eu *User) bool {
		return eu.Id == u.Id
	})
}
