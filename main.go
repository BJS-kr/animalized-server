package main

import (
	"animalized/controller"
	"animalized/lobby"
	"animalized/users"

	"log/slog"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9988")

	if err != nil {
		panic(err)
	}

	println("Listening")

	c := controller.New(100)

	for {
		conn, err := listener.Accept()

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		go handle(conn, c.Lobby)
	}
}

func handle(conn net.Conn, lobby *lobby.Lobby) error {
	u, err := users.Initialize(conn)

	if err != nil {
		return err
	}

	err = lobby.InitialJoin(u)

	if err != nil {
		return err
	}

	return nil
}
