package users

func (du *DistributableUsers) MakeWithUsers(maxUsers int) {
	du.Make()
	du.Users = New(maxUsers)
}
