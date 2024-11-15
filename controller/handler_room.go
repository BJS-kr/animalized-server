package controller

import (
	"animalized/game"
	"animalized/message"
	"animalized/rooms"
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
		return nil, errors.New("ROOM_STATUS not implemented")
	case message.Room_START:
		r.Game = game.New(r.Users.Max)

		for u := range r.Users.LockedRange() {
			r.Game.Users.Join(u, r.Game.InputChannel)
			r.Users.Quit(u)
		}

		r.Game.StartStreaming(c.makeHandleGame(r))
		r.Status = rooms.PLAYING
	case message.Room_QUIT:
		// TODO lobby status broadcast
		u, err := c.Rooms.Quit(roomInput.RoomName, input.UserId)

		if err != nil {
			return nil, err
		}

		err = c.Lobby.Join(u)
		// lobby에 join한 이후이므로 room handler는 유저에게 인풋을 전달할 수 없으므로 lobby input으로 보내준다.
		c.Lobby.Inputs.Enqueue(input)

		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	return input, nil
}
