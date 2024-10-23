package handler

import (
	"animalized/message"
	"animalized/users"
	"net"
)

func JoinLobby(lobby *users.Users, conn net.Conn, lobbyInputChannel chan<- *message.Input) error {
	u, err := initialize(conn)

	if err != nil {
		return err
	}

	err = lobby.InsertUser(u)

	if err != nil {
		return err
	}

	go StartHandlers(lobby, u, lobbyInputChannel)

	return nil
}
