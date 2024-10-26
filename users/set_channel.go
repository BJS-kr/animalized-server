package users

import "animalized/message"

func (u *User) SetProduceChannel(ch chan<- *message.Input) {
	u.produceChannel = ch
}
