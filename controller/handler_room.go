package controller

import (
	"animalized/game"
	"animalized/message"
	"animalized/packet"
	"animalized/rooms"
	"errors"
)

func (c *Controller) roomHandler(input *message.Input) (*message.Input, error) {
	if input.RoomName == nil {
		return nil, errors.New("room name not provided in room handler")
	}

	r, ok := c.Rooms.NameMap[rooms.RoomName(*input.RoomName)]

	if !ok {
		return nil, errors.New("room not found in room handler")
	}

	switch input.Type {
	case packet.ROOM_STATUS:
		return nil, errors.New("ROOM_STATUS not implemented")
	case packet.START:
		r.Game = game.New(r.Users.Max)

		for u := range r.Users.LockedRange() {
			r.Game.Users.Join(u, r.Game.InputChannel)
			r.Users.Quit(u)
		}

		r.Game.StartStreaming(c.makeHandleGame(r))
		r.Status = rooms.PLAYING
	case packet.QUIT:
		u, err := c.Rooms.Quit(*input.RoomName, input.UserId)

		if err != nil {
			return nil, err
		}

		err = c.Lobby.Join(u)
		// lobby에 join한 이후이므로 room handler는 유저에게 인풋을 전달할 수 없으므로 lobby input으로 보내준다.
		c.Lobby.InputChannel <- input

		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	return input, nil
}
