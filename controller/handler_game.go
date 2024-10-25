package controller

import (
	"animalized/common"
	"animalized/message"
	"animalized/packet"
	"animalized/rooms"
	"errors"
	"time"
)

func (c *Controller) makeHandleGame(r *rooms.Room) common.Handler {
	var context, prevContext int64
	return func(input *message.Input) (*message.Input, error) {
		switch input.Type {
		case packet.MOVE:
			r.Game.State.UpdateUserPosition(input.UserId, *input.Direction)
		case packet.SERVER_STATE:
			prevContext = context
			context = time.Now().UnixMilli()

			input = &message.Input{
				Type:        packet.SERVER_STATE,
				UserId:      "system",
				PrevContext: &prevContext,
				UserStates:  r.Game.State.GetUserStates(),
			}
		case packet.FINISH:
			return nil, errors.New("FINISH not implemented")

		}

		input.Context = &context

		return input, nil
	}
}
