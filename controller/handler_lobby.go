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
	case message.Lobby_CREATE_ROOM:
		u, err := c.Lobby.Session.FindUserById(input.UserId)

		if err != nil {
			return nil, err
		}

		r, err := c.Rooms.Create(lobbyInput.RoomName, int(lobbyInput.MaxUsers))

		if err != nil {
			return nil, err
		}

		u, err = c.Lobby.Quit(u.Id)

		if err != nil {
			return nil, err
		}

		_, err = c.Rooms.Join(lobbyInput.RoomName, u)

		if err != nil {
			c.Lobby.Join(u)
			c.Rooms.Remove(lobbyInput.RoomName)

			return nil, err
		}

		r.StartStreaming(c.roomHandler, ROOM_TICK_RATE)
		c.Lobby.SystemInput(c.MakeLobbyState(input.UserId))
		r.SystemDirectInput(c.MakeJoinInput(input.UserId, lobbyInput.RoomName))
		r.SystemInput(c.MakeRoomStateInput(input.UserId, lobbyInput.RoomName))

	case message.Lobby_JOIN_ROOM:
		u, err := c.Lobby.Session.FindUserById(input.UserId)

		if err != nil {
			return nil, err
		}

		u, err = c.Lobby.Quit(u.Id)

		if err != nil {
			return nil, err
		}

		r, err := c.Rooms.Join(lobbyInput.RoomName, u)

		if err != nil {
			c.Lobby.Join(u)

			return nil, err
		}

		r.SystemDirectInput(c.MakeJoinInput(input.UserId, lobbyInput.RoomName))
		r.SystemInput(c.MakeRoomStateInput(input.UserId, lobbyInput.RoomName))
		c.Lobby.SystemInput(c.MakeLobbyState(input.UserId))
	case message.Lobby_STATE:
		lobbyInput.RoomStates = c.Rooms.MakeRoomStates()
	default:
		return nil, errors.New("unknown lobby input type:" + lobbyInput.Type.String())
	}

	return input, nil
}
