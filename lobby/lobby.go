package lobby

import (
	"animalized/handler"
	"animalized/message"
	"animalized/packet"
	"net"
)

func (l *Lobby) JoinLobby(conn net.Conn, lobbyInputChannel chan<- *message.Input) error {
	u, err := packet.Initialize(conn)

	if err != nil {
		return err
	}

	err = l.users.InsertUser(u)

	if err != nil {
		return err
	}

	handler.StartHandlers(l.users, u, lobbyInputChannel)

	return nil
}
