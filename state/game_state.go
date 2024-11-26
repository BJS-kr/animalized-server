package state

import (
	"animalized/message"
	"errors"
	"time"
)

type UserID string
type GameState struct {
	UserStates map[UserID]*UserState
}

func New() *GameState {
	gs := new(GameState)
	gs.UserStates = make(map[UserID]*UserState)

	return gs
}

func (ss *GameState) UpdateUserPosition(userId string, direction message.Operation_Direction) {
	ss.UserStates[UserID(userId)].Position.determinePosition(direction)
}

func (ss *GameState) SignalGameState(inputProduceChannel chan<- *message.Input) {
	tick := time.Tick(SERVER_STATE_SIGNAL_INTERVAL)
	tickMessage := &message.Input{
		Kind: &message.Input_Op{
			Op: &message.Operation{
				Type: message.Operation_GAME_STATE,
			},
		},
	}

	for range tick {
		inputProduceChannel <- tickMessage
	}
}

func (ss *GameState) GetGameState() *message.Operation_GameState {
	gameState := new(message.Operation_GameState)

	for _, us := range ss.UserStates {
		mus := &message.Operation_GameState_UserState{
			Position: &message.Position{
				X: us.Position.X,
				Y: us.Position.Y,
			},
			Score: us.score,
		}

		gameState.UserStates = append(gameState.UserStates, mus)
	}

	return gameState
}

func (ss *GameState) AddUserState(userId UserID) error {
	if _, ok := ss.UserStates[userId]; ok {
		return errors.New("user id already exists")
	}

	us := UserState{}
	us.Position = &Position{
		X: 0,
		Y: 0,
	}

	ss.UserStates[userId] = &us

	return nil
}
