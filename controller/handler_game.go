package controller

import (
	"animalized/common"
	"animalized/message"
	"animalized/rooms"
	"animalized/state"
	"animalized/users"
	"errors"
	"time"
)

func (c *Controller) makeGameHandler(r *rooms.Room, roomName string) common.Handler {
	var context, prevContext int64

	return func(input *message.Input) (*message.Input, error) {
		opKind, ok := input.Kind.(*message.Input_Op)

		if !ok {
			return nil, errors.New("not operation input")
		}

		if r.Game == nil {
			return nil, errors.New("game not found. maybe game is over")
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

			targetUserState, ok := r.Game.State.UserStates[state.UserID(opInput.TargetUserId)]

			if !ok {
				return nil, errors.New("target user state not found")
			}

			if !targetUserState.Position.IsHit(opInput.HitRange) {
				return nil, errors.New("target user is not in hit range")
			}

			userState, ok := r.Game.State.UserStates[state.UserID(input.UserId)]

			if !ok {
				return nil, errors.New("user state not found")
			}

			userState.IncreaseUserScore(1)

			if userState.IsWinner() {
				users := make([]*users.User, 0)

				for u := range r.Game.LockedRange() {
					users = append(users, u)
				}

				for _, user := range users {
					r.Game.Quit(user)
					c.Lobby.Join(user)
				}

				r.Game.StopStreaming()
				c.Lobby.SystemDirectInput(c.MakeGameResultInput(input.UserId, roomName))
				c.Lobby.SystemInput(c.MakeLobbyState(input.UserId))

				return nil, nil
			}
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
