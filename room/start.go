package room

import (
	"animalized/handler"
	"animalized/message"
	"animalized/queue"
	"animalized/state"
	"errors"
)

// TODO user conn에 대한 lobby goroutines를 종료시켜야 한다.
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
		close(u.Quit)                // lobby goroutines를 종료
		u.Quit = make(chan struct{}) // game을 빠져 나올 때 쓸 채널
		gameState.AddUserState(state.UserID(u.Id))
		handler.StartHandlers(r.users, u, inputProduceChannel)
	}

	go handler.ReceiveGameInput(mainInputs, gameState, inputProduceChannel)
	go handler.Propagate(mainInputs, r.users)
	go gameState.SignalGameState(inputProduceChannel)

	return nil
}
