package main

import (
	"animalized/handler"
	"animalized/message"
	"animalized/queue"
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

	go handler.Receive(mainInputs, inputProduceChannel)
	go handler.Propagate(mainInputs, users)

	for {
		conn, err := listener.Accept()

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		handler.StartHandlers(users, conn, inputProduceChannel)
	}
}
