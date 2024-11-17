package controller

import (
	"animalized/message"
	"errors"
)

var lobbyInput *message.Lobby

func (c *Controller) lobbyHandler(input *message.Input) (*message.Input, error) {
	lobbyInputKind, ok := input.Kind.(*message.Input_Lobby)

	if !ok {
		return nil, errors.New("not lobby input")
	}

	lobbyInput = lobbyInputKind.Lobby

	switch lobbyInput.Type {
	case message.Lobby_CREATE:
		u, err := c.FindUserById(input.UserId)

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

			return nil, err
		}

		c.Lobby.SystemInput(c.MakeLobbyState())
		r.StartStreaming(c.roomHandler)
	case message.Lobby_JOIN:
		u, err := c.FindUserById(input.UserId)

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

		c.Lobby.SystemInput(c.MakeLobbyState())
		r.SystemInput(r.MakeRoomStateInput(lobbyInput.RoomName))
	default:
		return nil, errors.New("unknown lobby input type")

	}

	return input, nil
}
