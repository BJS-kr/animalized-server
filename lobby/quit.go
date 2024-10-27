package lobby

import "animalized/users"

func (l *Lobby) Quit(userId string) (*users.User, error) {
	u, err := l.Users.FindUserById(userId)

	if err != nil {
		return nil, err
	}
	_, err = l.Users.Quit(u)

	return u, err
}
