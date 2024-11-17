package state

import (
	"animalized/message"
	"errors"
	"time"
)

type UserID string
type GameState struct {
	UserStates map[UserID]UserState
}

func New() *GameState {
	gs := new(GameState)
	gs.UserStates = make(map[UserID]UserState)

	return gs
}

func (ss *GameState) UpdateUserPosition(userId string, direction message.Operation_Direction) {
	ss.UserStates[UserID(userId)].position.determinePosition(direction)
}

func (ss *GameState) SignalGameState(inputProduceChannel chan<- *message.Input) {
	tick := time.Tick(SERVER_STATE_SIGNAL_INTERVAL)
	tickMessage := &message.Input{
		Kind: &message.Input_Operation{
			Operation: &message.Operation{
				Type: message.Operation_GAME_STATE,
			},
		},
	}

	for range tick {
		inputProduceChannel <- tickMessage
	}
}

func (ss *GameState) GetUserStates() *message.UserStates {
	userStates := new(message.UserStates)

	for _, us := range ss.UserStates {
		mus := &message.UserStates_UserState{
			Position: &message.UserStates_Position{
				X: us.position.X,
				Y: us.position.Y,
			},
			Score: us.score,
		}

		userStates.UserStates = append(userStates.UserStates, mus)
	}

	return userStates
}

func (ss *GameState) AddUserState(userId UserID) error {
	if _, ok := ss.UserStates[userId]; ok {
		return errors.New("user id already exists")
	}

	us := UserState{}
	us.position = &Position{
		X: 0,
		Y: 0,
	}

	ss.UserStates[userId] = us

	return nil
}
