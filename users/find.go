package users

import "errors"

func (us *Users) FindUserById(userId string) (*User, error) {
	for u := range us.LockedRange() {
		if u.Id == userId {
			return u, nil
		}
	}

	return nil, errors.New("user not found")
}
