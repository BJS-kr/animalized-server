package lobby

import (
	"animalized/users"
)

func (l *Lobby) Join(user *users.User) error {
	err := l.Users.InsertUser(user)

	if err != nil {
		return err
	}

	users.StartHandlers(l.Users, user, l.InputChannel)

	return nil
}
