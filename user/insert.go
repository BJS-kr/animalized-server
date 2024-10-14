package user

func (us *Users) InsertUser(u *User) {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	us.users = append(us.users, u)
}
