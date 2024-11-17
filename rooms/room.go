package rooms

import (
	"animalized/game"
	"animalized/message"
	"animalized/users"
	"errors"
)

type Room struct {
	users.DistributableUsers
	Status message.RoomState_RoomStatusType
	Game   *game.Game
}

func (r *Room) Join(user *users.User) error {
	if err := r.Users.Join(user, r.InputChannel); err != nil {
		return err
	}

	return nil
}

func (r *Room) SetStatus(targetStatus message.RoomState_RoomStatusType) error {
	if targetStatus == message.RoomState_PLAYING && r.Status != message.RoomState_WAITING {
		return errors.New("cannot set room status as PLAYING")
	}

	r.Status = targetStatus

	return nil
}

func (r *Room) Quit(user *users.User) (int, error) {
	return r.Users.Quit(user)
}
