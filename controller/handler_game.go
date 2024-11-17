package controller

import (
	"animalized/common"
	"animalized/message"
	"animalized/rooms"
	"errors"
	"time"
)

var opInput *message.Operation

func (c *Controller) makeGameHandler(r *rooms.Room) common.Handler {
	var context, prevContext int64

	return func(input *message.Input) (*message.Input, error) {
		opKind, ok := input.Kind.(*message.Input_Operation)

		if !ok {
			return nil, errors.New("not operation input")
		}

		opInput = opKind.Operation

		switch opInput.Type {
		case message.Operation_MOVE:
			r.Game.State.UpdateUserPosition(input.UserId, opInput.Direction)
		case message.Operation_ATTACK:
		case message.Operation_HIT:
		case message.Operation_GAME_STATE:
		default:
			return nil, errors.New("unknown operation input type")
		}

		prevContext = context
		context = time.Now().UnixMilli()

		opInput.Context = context
		opInput.PrevContext = prevContext

		return input, nil
	}
}
