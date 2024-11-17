package controller

import (
	"animalized/lobby"
	"animalized/message"
	"animalized/rooms"
)

type Controller struct {
	*lobby.Lobby
	*rooms.Rooms
}

func New(maxUsers int) *Controller {
	c := new(Controller)

	c.Lobby = lobby.New(maxUsers)
	c.Rooms = rooms.New()

	c.Lobby.StartStreaming(c.lobbyHandler)

	return c
}

func (c *Controller) MakeLobbyState() *message.Input {
	return &message.Input{
		Kind: &message.Input_Lobby{
			Lobby: &message.Lobby{
				Type:       message.Lobby_STATE,
				RoomStates: c.MakeRoomStates(),
			},
		},
	}
}
