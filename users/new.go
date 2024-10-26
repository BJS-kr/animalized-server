package users

func New(maxUsers int) *Users {
	us := new(Users)
	us.list = make([]*User, 0, maxUsers)
	us.Max = maxUsers

	return us
}
