package controller

import (
	"animalized/game"
	"animalized/message"
	"animalized/packet"
	"animalized/rooms"
	"errors"
)

func (c *Controller) roomHandler(input *message.Input) (*message.Input, error) {
	r, ok := c.Rooms.NameMap[rooms.RoomName(*input.RoomName)]

	if !ok {
		return nil, errors.New("room not found")
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
		u, err := r.Users.FindUserById(input.UserId)

		if err != nil {
			return nil, err
		}

		err = c.Lobby.Users.Join(u, c.Lobby.InputChannel)

		if err != nil {
			return nil, err
		}

		_, err = r.Users.Quit(u)

		if err != nil {
			return nil, err
		}
	}

	return input, nil
}
