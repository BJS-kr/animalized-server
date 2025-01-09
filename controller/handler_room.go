package controller

import (
	"animalized/message"
	"animalized/rooms"
	"animalized/state"
	"errors"
)

func (c *Controller) roomHandler(input *message.Input) (*message.Input, error) {
	roomInputKind, ok := input.Kind.(*message.Input_Room)

	if !ok {
		return nil, errors.New("not room input")
	}

	roomInput := roomInputKind.Room

	if roomInput.RoomName == "" {
		return nil, errors.New("room name not provided in room handler")
	}

	r, ok := c.Rooms.NameMap[rooms.RoomName(roomInput.RoomName)]

	if !ok {
		return nil, errors.New("room not found in room handler")
	}

	switch roomInput.Type {
	case message.Room_STATE:
		roomInput.RoomState = r.MakeRoomState(roomInput.RoomName)
	case message.Room_START:
		r.StopStreaming()
		characterTypes := r.PickCharacterRandomTypes()
		userIds := r.Session.LockedIds()

		if len(userIds) <= 1 {
			return nil, errors.New("not enough users to start game")
		}

		for _, userId := range userIds {
			r.Game.State.AddUserState(state.UserID(userId))
			u, err := r.Session.FindUserById(userId)

			if err != nil {
				return nil, err
			}

			err = r.Game.JoinGame(u)

			if err != nil {
				return nil, err
			}

			_, err = r.Quit(u)

			if err != nil {
				return nil, err
			}
		}

		c.Rooms.Remove(roomInput.RoomName)
		r.Game.Init(c.makeGameHandler(r, roomInput.RoomName), GAME_TICK_RATE)
		r.Game.SystemDirectInput(c.MakeGameStartInput(input.UserId, roomInput.RoomName, characterTypes, r.Game.State.Terrains))
	case message.Room_QUIT:
		u, err := c.Rooms.Quit(roomInput.RoomName, input.UserId)

		if err != nil {
			return nil, err
		}

		err = c.Lobby.Join(u)

		if err != nil {
			return nil, err
		}

		c.Lobby.SystemInput(c.MakeLobbyState(input.UserId))

		if c.Rooms.NameMap[rooms.RoomName(roomInput.RoomName)] != nil {
			r.SystemDirectInput(c.MakeRoomStateDirectInput(input.UserId, roomInput.RoomName, r))
		}

		c.Lobby.SystemDirectInput(c.MakeQuitRoomInput(input.UserId, roomInput.RoomName))
	default:
		return nil, errors.New("unknown room input type")
	}

	return input, nil
}
