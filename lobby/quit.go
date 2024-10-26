package lobby

import "animalized/users"

func (l *Lobby) Quit(user *users.User) error {
	_, err := l.Users.Quit(user)

	return err
}
