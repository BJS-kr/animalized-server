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

func (gs *GameState) UpdateUserPosition(userId UserID, direction message.Operation_Direction) error {
	us, ok := gs.UserStates[userId]

	if !ok {
		return errors.New("user state not found")
	}

	us.Position.determinePosition(direction)

	return nil
}

func (gs *GameState) SignalGameState(inputProduceChannel chan<- *message.Input) {
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

func (gs *GameState) GetGameState() *message.Operation_GameState {
	gameState := new(message.Operation_GameState)

	for _, us := range gs.UserStates {
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

func (gs *GameState) AddUserState(userId UserID) error {
	if _, ok := gs.UserStates[userId]; ok {
		return errors.New("user id already exists")
	}

	us := UserState{}
	us.Position = &Position{
		X: 0,
		Y: 0,
	}

	gs.UserStates[userId] = &us

	return nil
}