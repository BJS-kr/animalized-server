package main

import (
	"animalized/handler"
	"animalized/message"
	"animalized/users"

	"log/slog"
	"net"
)

func main() {
	lobby := new(users.Users)
	lobby.Max = 100
	lobbyInputChannel := make(chan *message.Input, 100)

	listener, err := net.Listen("tcp", "127.0.0.1:9988")

	if err != nil {
		panic(err)
	}

	go handler.ReceiveLobbyInput(lobby, lobbyInputChannel)

	for {
		conn, err := listener.Accept()

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		go handler.JoinLobby(lobby, conn, lobbyInputChannel)
	}
}
