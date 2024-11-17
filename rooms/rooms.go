package rooms

import (
	"animalized/message"
	"animalized/users"
	"errors"
)

type Rooms struct {
	NameMap map[RoomName]*Room
}

func New() *Rooms {
	rs := new(Rooms)
	rs.NameMap = make(map[RoomName]*Room)

	return rs
}

func (rs *Rooms) Create(roomName string, maxUsers int) (*Room, error) {
	if roomName == "" {
		return nil, errors.New("room name not provided when creating room")
	}

	if maxUsers <= 0 || maxUsers > MAX_USERS_LIMIT {
		return nil, errors.New("max users not in valid range")
	}

	if r, ok := rs.NameMap[RoomName(roomName)]; ok {
		return r, errors.New("room already exists")
	}

	r := new(Room)

	if maxUsers > MAX_USERS_LIMIT {
		return r, errors.New("room users limit has exceeded")
	}

	r.MakeWithUsers(maxUsers)
	r.SetStatus(message.RoomState_WAITING)
	rs.NameMap[RoomName(roomName)] = r

	return r, nil
}

func (rs *Rooms) Join(roomName string, user *users.User) (*Room, error) {
	r, ok := rs.NameMap[RoomName(roomName)]

	if !ok {
		return nil, errors.New("room not exists")
	}

	if err := r.Join(user); err != nil {
		return nil, err
	}

	return r, nil
}

func (rs *Rooms) RemoveRoom(room *Room) {
	for k, v := range rs.NameMap {
		if v == room {
			room.StopStreaming()
			delete(rs.NameMap, k)
			return
		}
	}
}

func (rs *Rooms) Quit(roomName string, userName string) (*users.User, error) {
	r, ok := rs.NameMap[RoomName(roomName)]

	if !ok {
		return nil, errors.New("room does not exists")
	}

	user, err := r.Users.FindUserById(userName)

	if err != nil {
		return nil, err
	}

	remain, err := r.Quit(user)

	if err != nil {
		return nil, err
	}

	if remain <= 0 {
		delete(rs.NameMap, RoomName(roomName))
	}

	return user, nil
}
