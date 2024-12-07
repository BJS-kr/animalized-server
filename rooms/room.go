package rooms

import (
	"animalized/game"
	"animalized/message"
	"animalized/users"
	"errors"
	"math/rand"
)

type Room struct {
	users.DistSession
	Status message.RoomState_RoomStatusType
	Game   *game.Game
}

type UserCharacterTypes map[string]message.Room_CharacterType

func (r *Room) Join(user *users.User) error {
	if r.Status != message.RoomState_WAITING {
		return errors.New("room is not waiting")
	}

	if err := r.Session.Join(user, r.Receiver); err != nil {
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
	return r.Session.Quit(user)
}

// Room struct 자체는 Name을 가지고 있지 않으므로 인자로 받는다.
func (r *Room) MakeRoomState(roomName string) *message.RoomState {
	rs := new(message.RoomState)

	rs.RoomName = roomName
	rs.MaxUsers = int32(r.Session.Max)
	rs.Status = r.Status
	rs.UserIds = r.Session.LockedIds()

	return rs
}

func (r *Room) PickCharacterRandomTypes() UserCharacterTypes {
	ids := r.Session.LockedIds()
	count := len(ids)
	remain := 0
	typeMap := make(UserCharacterTypes)

	if count%CHARACTER_TYPES_COUNT == 0 {
		remain = count / CHARACTER_TYPES_COUNT
	} else {
		remain = count/CHARACTER_TYPES_COUNT + 1
	}

	characterRemains := [CHARACTER_TYPES_COUNT]int{remain, remain, remain}
	picked := make([]message.Room_CharacterType, 0, count)

	for {
		pickedI := rand.Intn(CHARACTER_TYPES_COUNT)

		biggestRemain := characterRemains[pickedI]

		for _, remain := range characterRemains {
			if remain > biggestRemain {
				biggestRemain = remain
			}
		}

		if characterRemains[pickedI] != biggestRemain {
			continue
		}

		if characterRemains[pickedI] > 0 {
			picked = append(picked, message.Room_CharacterType(pickedI+1))
			characterRemains[pickedI]--
			count--
		}

		if count == 0 {
			break
		}
	}

	for i, id := range ids {
		typeMap[id] = picked[i]
	}

	return typeMap
}
