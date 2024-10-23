package lobby

import (
	"animalized/message"
	"animalized/packet"
)

// start in main routine
func (l *Lobby) ReceiveLobbyInput(lobbyInputChannel <-chan *message.Input) {
	for input := range lobbyInputChannel {
		switch input.Type {
		case packet.LOBBY_STATUS:

		case packet.CREATE:

		case packet.JOIN:
		}
		l.Inputs.Enqueue(input)
	}
}
