package rooms

import (
	"animalized/message"
	"animalized/packet"
	"animalized/users"
	"errors"
)

func (rs *Rooms) Join(roomName string, user *users.User) error {
	r, ok := rs.NameMap[RoomName(roomName)]

	if !ok {
		return errors.New("room not exists")
	}

	if err := r.Join(user); err != nil {
		return err
	}

	r.Inputs.Enqueue(&message.Input{
		UserId:   user.Id,
		Type:     packet.JOIN,
		RoomName: &roomName,
	})

	// send when join or qui
	return nil
}

func (r *Room) Join(user *users.User) error {
	if err := r.Users.Join(user, r.InputChannel); err != nil {
		return err
	}

	return nil
}
