package users

import (
	"animalized/common"
	"animalized/message"
	"errors"
)

func (us *Users) Join(u *User, inputProduceChannel chan<- *message.Input) error {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	if len(us.list) >= us.Max {
		return errors.New("users max capacity reached")
	}

	u.Stop = make(chan common.Signal)
	u.SetProduceChannel(inputProduceChannel)
	us.list = append(us.list, u)

	return nil
}
