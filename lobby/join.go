package lobby

import (
	"animalized/handler"
	"animalized/users"
)

func (l *Lobby) Join(user *users.User) error {
	err := l.Users.InsertUser(user)

	if err != nil {
		return err
	}

	handler.StartHandlers(l.Users, user, l.InputChannel)

	return nil
}
