package room

import "errors"

func (rs *Rooms) Create(roomName string, particiPantsLImit int) error {
	r := new(Room)

	if particiPantsLImit > 8 {
		return errors.New("room participant limit has exceeded")
	}

	r.participantsLimit = particiPantsLImit
	r.status = READY

	rs.Rooms[RoomName(roomName)] = r

	return nil
}
