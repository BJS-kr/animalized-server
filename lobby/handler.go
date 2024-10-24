package lobby

import (
	"animalized/message"
	"animalized/packet"
	"errors"
)

func (l *Lobby) handler(input *message.Input) (*message.Input, error) {
	u := l.Users.FindUserById(input.UserId)

	if u == nil {
		return nil, errors.New("user not found")
	}

	switch input.Type {
	// 유저가 로그인 할 때
	// 누군가가 방에 JOIN하거나 QUIT할 때마다 통전송
	case packet.LOBBY_STATUS:

	case packet.CREATE:
		r, err := l.rooms.Create(*input.RoomName, int(*input.UsersLimit))

		if err != nil {
			return nil, err
		}

		err = l.rooms.Join(*input.RoomName, u)

		if err != nil {

			l.rooms.RemoveRoom(r)
			return nil, err
		}

		l.Users.Quit(u)
	case packet.JOIN:
		err := l.rooms.Join(*input.RoomName, u)

		if err != nil {
			return nil, err
		}

		l.Users.Quit(u)
	}

	return input, nil
}
