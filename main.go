package main

import (
	"animalized/lobby"
	"animalized/users"

	"log/slog"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9988")

	if err != nil {
		panic(err)
	}

	lobby := lobby.New(100)

	for {
		conn, err := listener.Accept()

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		go handle(conn, lobby)
	}
}

func handle(conn net.Conn, lobby *lobby.Lobby) error {
	u, err := users.Initialize(conn)

	if err != nil {
		return err
	}

	err = lobby.Join(u)

	if err != nil {
		return err
	}

	return nil
}
