package users

func handleOutgoing(u *User) {
	for {
		select {
		case <-u.Stop:
			return
		default:
			Consume(u)
		}
	}
}
