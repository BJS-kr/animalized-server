package users

func (us *Users) LockedIds() []string {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	ids := make([]string, 0, len(us.list))

	for _, u := range us.list {
		ids = append(ids, u.Id)
	}

	return ids
}
