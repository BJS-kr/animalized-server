package users

func (us *Users) LockedLen() int {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	return len(us.list)
}
