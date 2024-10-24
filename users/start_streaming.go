package users

import "animalized/common"

func (du *DistributableUsers) StartStreaming(handler common.Handler) {
	go du.Receive(handler)
	go du.Distribute()
}
