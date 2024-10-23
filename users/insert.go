package users

import "errors"

func (us *Users) InsertUser(u *User) error {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	if len(us.users) >= us.Max {
		return errors.New("users max capacity reached")
	}

	us.users = append(us.users, u)

	return nil
}
