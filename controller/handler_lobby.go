package controller

import (
	"animalized/message"
	"animalized/packet"
	"errors"
)

func (c *Controller) lobbyHandler(input *message.Input) (*message.Input, error) {
	switch input.Type {
	// 유저가 로그인 할 때
	// 누군가가 방에 JOIN하거나 QUIT할 때마다 통전송
	case packet.LOBBY_STATUS:
		rss := make([]*message.RoomStatus, len(c.Rooms.NameMap))
		for name, room := range c.Rooms.NameMap {
			rss = append(rss, &message.RoomStatus{
				Name:       string(name),
				MaxUsers:   int32(room.Users.Max),
				UsersCount: int32(room.Users.LockedLen()),
			})
		}

		input.LobbyStatus = &message.LobbyStatus{
			RoomStatuses: rss,
		}
	case packet.CREATE:
		if input.RoomName == nil {
			return nil, errors.New("room name not provided when creating room")
		}
		if input.MaxUsers == nil {
			return nil, errors.New("max users not provided when creating room")
		}

		r, err := c.Rooms.Create(*input.RoomName, int(*input.MaxUsers))

		if err != nil {
			return nil, err
		}

		r.StartStreaming(c.roomHandler)
	case packet.JOIN:
		if input.RoomName == nil {
			return nil, errors.New("room name not provided when join room")
		}

		u, err := c.Lobby.Quit(input.UserId)

		if err != nil {
			return nil, err
		}

		err = c.Rooms.Join(*input.RoomName, u)

		if err != nil {
			c.Lobby.Join(u)
			return nil, err
		}
	}

	return input, nil
}
