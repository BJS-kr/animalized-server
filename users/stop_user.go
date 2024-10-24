package users

func (u *User) StopPacketHandlers() {
	close(u.Stop)
}
