package users

func New(usersLimit int) *Users {
	us := new(Users)
	us.list = make([]*User, 0)
	us.Max = usersLimit

	return us
}
