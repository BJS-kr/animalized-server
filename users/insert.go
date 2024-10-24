package users

import (
	"animalized/common"
	"errors"
)

func (us *Users) InsertUser(u *User) error {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	if len(us.list) >= us.Max {
		return errors.New("users max capacity reached")
	}

	u.Stop = make(chan common.Signal)
	us.list = append(us.list, u)

	return nil
}
