package controller

import (
	"animalized/lobby"
	"animalized/rooms"
)

func New(maxUsers int) *Controller {
	c := new(Controller)

	c.Lobby = lobby.New(maxUsers)
	c.Rooms = rooms.New()

	c.Lobby.StartStreaming(c.lobbyHandler)

	return c
}
