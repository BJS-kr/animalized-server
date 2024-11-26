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
	if r.Status != message.RoomState_WAITING {
		return errors.New("room is not waiting")
	}

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

// Room struct 자체는 Name을 가지고 있지 않으므로 인자로 받는다.
func (r *Room) MakeRoomState(roomName string) *message.RoomState {
	rs := new(message.RoomState)

	rs.RoomName = roomName
	rs.MaxUsers = int32(r.Users.Max)
	rs.Status = r.Status
	rs.UserIds = r.Users.LockedIds()

	return rs
}
