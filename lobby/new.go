package lobby

func New(max int) *Lobby {
	l := new(Lobby)

	l.MakeWithUsers(max)

	return l
}
