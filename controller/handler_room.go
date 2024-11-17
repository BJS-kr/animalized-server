package controller

import (
	"animalized/message"
	"animalized/rooms"
	"errors"
)

var roomInput *message.Room

func (c *Controller) roomHandler(input *message.Input) (*message.Input, error) {
	roomInputKind, ok := input.Kind.(*message.Input_Room)

	if !ok {
		return nil, errors.New("not room input")
	}

	roomInput = roomInputKind.Room

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
		r.StartStreaming(c.makeGameHandler(r))
		r.SetStatus(message.RoomState_PLAYING)
	case message.Room_QUIT:
		u, err := c.Rooms.Quit(roomInput.RoomName, input.UserId)

		if err != nil {
			return nil, err
		}

		err = c.Lobby.Join(u)

		if err != nil {
			return nil, err
		}

		if c.Rooms.NameMap[rooms.RoomName(roomInput.RoomName)] == nil {
			c.Lobby.SystemInput(c.MakeLobbyState())
		} else {
			r.SystemInput(r.MakeRoomStateInput(roomInput.RoomName))
		}
	default:
		return nil, errors.New("unknown room input type")
	}

	return input, nil
}
