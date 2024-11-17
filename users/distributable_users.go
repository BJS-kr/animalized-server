package users

import "animalized/common"

type DistributableUsers struct {
	common.Distributable
	*Users
}

func (du *DistributableUsers) Distribute() {
	for {
		select {
		case <-du.Stop:
			return
		default:
			n := du.Inputs.Dequeue()

			if n == nil {
				continue
			}

			for u := range du.Users.LockedRange() {
				u.Inputs.Enqueue(n.Value)
			}
		}
	}
}

func (du *DistributableUsers) MakeWithUsers(maxUsers int) {
	du.Make()
	du.Users = NewUsers(maxUsers)
}

func (du *DistributableUsers) StartStreaming(handler common.Handler) {
	go du.Receive(handler)
	go du.Distribute()
}

func (du *DistributableUsers) StopStreaming() {
	close(du.Stop)
}
