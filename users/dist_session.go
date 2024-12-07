package users

import (
	"animalized/common"
	"animalized/message"
	"log/slog"
	"time"
)

type DistSession struct {
	Receiver chan *message.Input
	// Dispatcher: SystemDirectInput과 Distribute에 의해 사용되는 채널
	Dispatcher chan *message.Input
	Stop       chan common.Signal
	*Session
}

func (ds *DistSession) Make(maxUsers int) {
	ds.Receiver = make(chan *message.Input, 100)
	ds.Dispatcher = make(chan *message.Input, 100)
	ds.Stop = make(chan common.Signal)
	ds.Session = NewSession(maxUsers)
}

func (ds *DistSession) Receive(handler common.Handler) {
	for input := range ds.Receiver {
		input, err := handler(input)

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		if input == nil {
			continue
		}

		select {
		case ds.Dispatcher <- input:
		case <-ds.Stop:
			return
		case <-time.After(1 * time.Second):
			slog.Error("timeout: failed to dispatch input")
			continue
		}
	}
}

// 해당하는 컨텍스트의 핸들러를 거치지 않는 input
func (ds *DistSession) SystemDirectInput(message *message.Input) {
	ds.Dispatcher <- message
}

// 해당하는 컨텍스트의 핸들러를 거치는 input
// ex) Room state를 handle하는 과정은 roomHandler에 있으므로 별도의 로직을 작성해 분산시키지 않고 handler에 보낸다.
func (ds *DistSession) SystemInput(message *message.Input) {
	ds.Receiver <- message
}

func (ds *DistSession) Distribute() {
	for {
		select {
		case <-ds.Stop:
			return
		case input := <-ds.Dispatcher:
			for u := range ds.Session.LockedRange() {
				u.outgoingQueue.Enqueue(input)
			}
		}
	}
}

func (ds *DistSession) StartStreaming(handler common.Handler) {
	go ds.Receive(handler)
	go ds.Distribute()
}

func (ds *DistSession) StopStreaming() {
	close(ds.Stop)
	ds.Stop = make(chan common.Signal)
	ds.Dispatcher = make(chan *message.Input, 100)
}
