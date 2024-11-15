package controller

import (
	"animalized/message"
	"errors"
)

func (c *Controller) lobbyHandler(input *message.Input) (*message.Input, error) {
	lobbyInputKind, ok := input.Kind.(*message.Input_Lobby)

	if !ok {
		return nil, errors.New("not lobby input")
	}

	lobbyInput := lobbyInputKind.Lobby

	switch lobbyInput.Type {
	// 유저가 로그인 할 때
	// 누군가가 방에 JOIN하거나 QUIT할 때마다 통전송
	case message.Lobby_STATE:
		lobbyInput.RoomStates = make([]*message.RoomState, 0, len(c.Rooms.NameMap))

		for name, room := range c.Rooms.NameMap {
			lobbyInput.RoomStates = append(lobbyInput.RoomStates, &message.RoomState{
				RoomName: string(name),
				MaxUsers: int32(room.Users.Max),
				UserIds:  room.Users.LockedIds(),
			})
		}
	case message.Lobby_CREATE:
		// TODO lobby status input 추가
		r, err := c.Rooms.Create(lobbyInput.RoomName, int(lobbyInput.MaxUsers))

		if err != nil {
			return nil, err
		}

		r.StartStreaming(c.roomHandler)
	case message.Lobby_JOIN:
		// TODO lobby status input 추가
		// TODO room에 status input broadcast

	}

	return input, nil
}
