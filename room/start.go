package room

import "errors"

func (rs *Rooms) Start(roomName string) error {
	r, ok := rs.Rooms[RoomName(roomName)]

	if !ok {
		return errors.New("room does not exists")
	}

	r.status = PLAYING

	return nil
}
