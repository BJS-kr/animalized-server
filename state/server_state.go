package state

import (
	"animalized/message"
	"animalized/packet"
	"errors"
	"time"
)

type UserID string
type GameState struct {
	UserStates map[UserID]UserState
}

func NewGameState() *GameState {
	ss := new(GameState)
	ss.UserStates = make(map[UserID]UserState)

	return ss
}

func (ss *GameState) UpdateUserPosition(userId string, direction int32) {
	ss.UserStates[UserID(userId)].position.determinePosition(direction)
}

func (ss *GameState) SignalGameState(inputProduceChannel chan<- *message.Input) {
	tick := time.Tick(SERVER_STATE_SIGNAL_INTERVAL)
	tickMessage := &message.Input{
		Type: packet.SERVER_STATE,
	}

	for range tick {
		inputProduceChannel <- tickMessage
	}
}

func (ss *GameState) GetUserStates() []*message.UserState {
	userStates := make([]*message.UserState, 0)

	for _, us := range ss.UserStates {
		mus := &message.UserState{
			Position: &message.Position{
				X: us.position.X,
				Y: us.position.Y,
			},
			Score: us.score,
		}

		userStates = append(userStates, mus)
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
