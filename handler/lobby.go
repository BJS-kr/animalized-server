package handler

import (
	"animalized/users"
	"log/slog"
	"net"
)

func JoinLobby(lobby *users.Users, conn net.Conn) {
	u, err := initialize(conn)

	if err != nil {
		slog.Error(err.Error())
		return
	}

	lobby.InsertUser(u)
}
