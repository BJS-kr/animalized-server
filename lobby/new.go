package lobby

func New(max int) *Lobby {
	l := new(Lobby)

	l.Make()
	l.Users.Max = max

	return l
}
