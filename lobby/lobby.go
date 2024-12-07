package lobby

import (
	"animalized/users"
)

type Lobby struct {
	users.DistSession
}

func New(max int) *Lobby {
	l := new(Lobby)

	l.Make(max)

	return l
}

func (l *Lobby) InitialJoin(user *users.User) error {
	err := l.Session.Join(user, l.Receiver)

	if err != nil {
		return err
	}

	user.StartPacketHandlers(l.Session)

	return nil
}

func (l *Lobby) Join(user *users.User) error {
	return l.Session.Join(user, l.Receiver)
}

func (l *Lobby) Quit(userId string) (*users.User, error) {
	u, err := l.Session.FindUserById(userId)

	if err != nil {
		return nil, err
	}
	_, err = l.Session.Quit(u)

	return u, err
}
