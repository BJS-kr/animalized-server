package controller

import (
	"animalized/message"
	"animalized/packet"
	"errors"
)

func (c *Controller) handler(input *message.Input) (*message.Input, error) {
	u := c.Users.FindUserById(input.UserId)

	if u == nil {
		return nil, errors.New("user not found")
	}

	switch input.Type {
	// 유저가 로그인 할 때
	// 누군가가 방에 JOIN하거나 QUIT할 때마다 통전송
	case packet.LOBBY_STATUS:

	case packet.CREATE:
		r, err := c.Rooms.Create(*input.RoomName, int(*input.UsersLimit))

		if err != nil {
			return nil, err
		}

		err = c.Rooms.Join(*input.RoomName, u)

		if err != nil {

			c.Rooms.RemoveRoom(r)
			return nil, err
		}

		c.Lobby.Users.Quit(u)
	case packet.JOIN:
		err := c.Rooms.Join(*input.RoomName, u)

		if err != nil {
			return nil, err
		}

		c.Lobby.Users.Quit(u)
	}

	return input, nil
}
