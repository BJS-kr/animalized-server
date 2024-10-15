package state

import (
	"animalized/message"
	"animalized/packet"
	"errors"
	"time"
)

type UserID string
type ServerState struct {
	UserStates map[UserID]UserState
}

func (ss *ServerState) UpdateUserPosition(userId string, direction int32) {
	ss.UserStates[UserID(userId)].position.determinePosition(direction)
}

func (ss *ServerState) SignalServerState(inputProduceChannel chan<- *message.Input) {
	tick := time.Tick(time.Second)
	tickMessage := &message.Input{
		Type: packet.SERVER_STATE,
	}

	for range tick {
		inputProduceChannel <- tickMessage
	}
}

func (ss *ServerState) GetUserStates() []*message.UserState {
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

func (ss *ServerState) AddUserState(userId UserID) error {
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
