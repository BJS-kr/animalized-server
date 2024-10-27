package users

import (
	"animalized/message"
	"animalized/packet"
	"errors"
)

func (u *User) validateInput(input *message.Input) error {
	if input.UserId != u.Id {
		return errors.New("user id not matched")
	}

	if input.Type == packet.INIT {
		return errors.New("init packet not allowed in producer")
	}

	if input.Type == packet.MOVE && input.Direction == nil {
		return errors.New("move packet did not include direction")
	}

	return nil
}
