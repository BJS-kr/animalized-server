package room

import (
	"animalized/users"
	"errors"
)

func (rs *Rooms) Join(roomName string, user *users.User) error {
	r, ok := rs.Rooms[RoomName(roomName)]

	if !ok {
		return errors.New("room not exists")
	}

	if err := r.Join(user); err != nil {
		return err
	}

	return nil
}

func (r *Room) Join(user *users.User) error {
	if len(r.participants) >= r.participantsLimit {
		return errors.New("room full")
	}

	r.participants = append(r.participants, user)

	return nil
}
