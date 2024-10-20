package main

import (
	"animalized/handler"
	"animalized/message"
	"animalized/queue"
	"animalized/state"
	"animalized/user"

	"log/slog"
	"net"
)

var users = new(user.Users)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9988")

	if err != nil {
		panic(err)
	}

	mainInputs := queue.New[*message.Input]()
	inputProduceChannel := make(chan *message.Input, 100)
	serverState := state.New()

	go handler.Receive(mainInputs, serverState, inputProduceChannel)
	go handler.Propagate(mainInputs, users)
	go serverState.SignalServerState(inputProduceChannel)

	for {
		conn, err := listener.Accept()

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		go handler.StartHandlers(users, serverState, conn, inputProduceChannel)
	}
}
