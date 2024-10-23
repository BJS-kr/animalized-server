package room

import (
	"animalized/handler"
	"animalized/message"
	"animalized/queue"
	"animalized/state"
	"errors"
)

func (rs *Rooms) Start(roomName string) error {
	r, ok := rs.Rooms[RoomName(roomName)]

	if !ok {
		return errors.New("room does not exists")
	}

	r.status = PLAYING

	mainInputs := queue.New[*message.Input]()
	inputProduceChannel := make(chan *message.Input, 100)
	gameState := state.NewGameState()

	for u := range r.users.LockedRange() {
		gameState.AddUserState(state.UserID(u.Id))
		handler.StartHandlers(r.users, u, gameState, inputProduceChannel)
	}

	go handler.Receive(mainInputs, gameState, inputProduceChannel)
	go handler.Propagate(mainInputs, r.users)
	go gameState.SignalGameState(inputProduceChannel)

	return nil
}
