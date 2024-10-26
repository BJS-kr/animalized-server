package users

import (
	"errors"
	"slices"
)

func (us *Users) Quit(user *User) (int, error) {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	var found *User

	us.list = slices.DeleteFunc(us.list, func(u *User) bool {
		if u == user {
			found = u

			return true
		}

		return false
	})

	if found == nil {
		return len(us.list), errors.New("failed to quit. user not found")
	}

	return len(us.list), nil
}
