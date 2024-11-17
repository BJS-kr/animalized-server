package rooms

import (
	"animalized/message"
	"errors"
)

func (r *Room) SetStatus(targetStatus message.RoomState_RoomStatusType) error {
	if targetStatus == message.RoomState_PLAYING && r.Status != message.RoomState_WAITING {
		return errors.New("cannot set room status as PLAYING")
	}

	r.Status = targetStatus

	return nil
}
