package lobby

import (
	"animalized/users"
)

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
