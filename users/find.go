package users

func (us *Users) FindUserById(userId string) *User {
	for u := range us.LockedRange() {
		if u.Id == userId {
			return u
		}
	}

	return nil
}
