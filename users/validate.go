package users

import (
	"animalized/message"
	"errors"
)

func (u *User) validateInput(input *message.Input) error {
	if input.UserId != u.Id {
		return errors.New("user id not matched")
	}

	return nil
}
