package users

import (
	"animalized/message"
	"animalized/packet"
	"slices"
)

func (us *Users) RemoveUser(u *User, inputProduceChannel chan<- *message.Input) {
	us.mtx.Lock()
	defer us.mtx.Unlock()

	us.users = slices.DeleteFunc(us.users, func(eu *User) bool {
		return eu.Id == u.Id
	})

	inputProduceChannel <- &message.Input{
		Type:   packet.QUIT,
		UserId: u.Id,
	}
}
