package users

import (
	"animalized/common"
	"animalized/message"
	"animalized/queue"
)

type DistributableSession struct {
	common.Actor
	*Session
}

func (ds *DistributableSession) Distribute() {
	for {
		select {
		case <-ds.Stop:
			return
		default:
			n := ds.Inputs.Dequeue()

			if n == nil {
				continue
			}

			for u := range ds.Session.LockedRange() {
				u.Inputs.Enqueue(n.Value)
			}
		}
	}
}

func (ds *DistributableSession) MakeWithSession(maxUsers int) {
	ds.Make()
	ds.Session = NewSession(maxUsers)
}

func (ds *DistributableSession) StartStreaming(handler common.Handler) {
	go ds.Receive(handler)
	go ds.Distribute()
}

func (ds *DistributableSession) StopStreaming() {
	close(ds.Stop)
	ds.Stop = make(chan common.Signal)
	ds.Inputs = queue.New[*message.Input]()
}
