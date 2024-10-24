package lobby

import (
	"animalized/users"
)

func (l *Lobby) Join(user *users.User) error {
	err := l.Users.Join(user, l.InputChannel)

	if err != nil {
		return err
	}

	return nil
}
