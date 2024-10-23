package lobby

import (
	"animalized/message"
	"animalized/packet"
	"log/slog"
)

// start in main routine
func (l *Lobby) ReceiveLobbyInput(lobbyInputChannel <-chan *message.Input) {
	for input := range lobbyInputChannel {
		u := l.users.FindUserById(input.UserId)

		if u == nil {
			continue
		}

		switch input.Type {
		// 유저가 로그인 할 때
		// 누군가가 방에 JOIN하거나 QUIT할 때마다 통전송
		case packet.LOBBY_STATUS:

		case packet.CREATE:
			r, err := l.rooms.Create(input.RoomName, input.UsersLimit)

			if err != nil {
				slog.Error(err.Error())
				continue
			}

			err = l.rooms.Join(input.RoomName, u)

			if err != nil {
				slog.Error(err.Error())
				l.rooms.RemoveRoom(r)
				continue
			}

			l.users.RemoveUser(u)
		case packet.JOIN:
			err := l.rooms.Join(input.RoomName, u)

			if err != nil {
				slog.Error(err.Error())
				continue
			}

			l.users.RemoveUser(u)
		}

		l.inputs.Enqueue(input)
	}
}
