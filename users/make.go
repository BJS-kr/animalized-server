package users

func (du *DistributableUsers) MakeWithUsers(usersLimit int) {
	du.Make()
	du.Users = New(usersLimit)
}
