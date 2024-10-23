package main

import (
	"animalized/lobby"
	"animalized/message"

	"log/slog"
	"net"
)

func main() {
	lobby := lobby.New(100)
	lobbyInputChannel := make(chan *message.Input, 100)

	listener, err := net.Listen("tcp", "127.0.0.1:9988")

	if err != nil {
		panic(err)
	}

	go lobby.ReceiveLobbyInput(lobbyInputChannel)
	lobby.Propagate()

	for {
		conn, err := listener.Accept()

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		go lobby.JoinLobby(conn, lobbyInputChannel)
	}
}
