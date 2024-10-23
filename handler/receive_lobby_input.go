package handler

import (
	"animalized/message"
	"animalized/packet"
	"animalized/users"
)

// start in main routine
func ReceiveLobbyInput(lobby *users.Users, lobbyInputChannel <-chan *message.Input) {
	for input := range lobbyInputChannel {
		switch input.Type {
		case packet.CREATE:
		case packet.JOIN:
		case packet.QUIT:
		}
	}
}
