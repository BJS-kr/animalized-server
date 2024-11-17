package lobby

import (
	"animalized/users"
)

type Lobby struct {
	users.DistributableUsers
}

func New(max int) *Lobby {
	l := new(Lobby)

	l.MakeWithUsers(max)

	return l
}

func (l *Lobby) InitialJoin(user *users.User) error {
	err := l.Users.Join(user, l.InputChannel)

	if err != nil {
		return err
	}

	user.StartPacketHandlers(l.Users)

	return nil
}

func (l *Lobby) Join(user *users.User) error {
	return l.Users.Join(user, l.InputChannel)
}

func (l *Lobby) Quit(userId string) (*users.User, error) {
	u, err := l.Users.FindUserById(userId)

	if err != nil {
		return nil, err
	}
	_, err = l.Users.Quit(u)

	return u, err
}
