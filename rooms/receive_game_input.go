package rooms

import (
	"animalized/message"
	"animalized/packet"
	"animalized/queue"
	"animalized/state"
	"time"
)

func (r *Room) ReceiveGameInput(inputQueue *queue.Queue[*message.Input], gameState *state.GameState, receiveChannel <-chan *message.Input) {
	var prevContext, context int64

	for input := range receiveChannel {
		switch input.Type {
		case packet.MOVE:
			gameState.UpdateUserPosition(input.UserId, *input.Direction)
		case packet.SERVER_STATE:
			prevContext = context
			context = time.Now().UnixMilli()

			input = &message.Input{
				Type:        packet.SERVER_STATE,
				UserId:      "system",
				PrevContext: &prevContext,
				UserStates:  gameState.GetUserStates(),
			}
		}

		input.Context = &context

		inputQueue.Enqueue(input)
	}
}
