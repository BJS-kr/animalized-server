package users

func (u *User) StartPacketHandlers(users *Users) {
	go u.handleIncoming(users)
	go u.handleOutgoing()
}
