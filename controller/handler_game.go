package controller

import (
	"animalized/common"
	"animalized/message"
	"animalized/rooms"
	"animalized/state"
	"errors"
	"time"
)

func (c *Controller) makeGameHandler(r *rooms.Room) common.Handler {
	var (
		context, prevContext int64
		userState            *state.UserState
	)

	return func(input *message.Input) (*message.Input, error) {
		opKind, ok := input.Kind.(*message.Input_Op)

		if !ok {
			return nil, errors.New("not operation input")
		}

		opInput := opKind.Op

		switch opInput.Type {
		case message.Operation_MOVE:
			if err := r.Game.State.UpdateUserPosition(state.UserID(input.UserId), opInput.Direction); err != nil {
				return nil, err
			}
		case message.Operation_ATTACK:
		case message.Operation_HIT:
			if opInput.TargetUserId == "" {
				return nil, errors.New("target user id not provided in hit operation. not fatal")
			}

			userState = r.Game.State.UserStates[state.UserID(opInput.TargetUserId)]

			if !userState.Position.IsHit(opInput.HitRange) {
				return nil, errors.New("target user is not in hit range")
			}

			userState.IncreaseUserScore(1)
		case message.Operation_GAME_STATE:
			opInput.GameState = r.Game.State.GetGameState()
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
