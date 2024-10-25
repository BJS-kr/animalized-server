package controller

import (
	"animalized/lobby"
	"animalized/rooms"
)

func New(usersLimit int) *Controller {
	c := new(Controller)

	c.Lobby = lobby.New(usersLimit)
	c.Rooms = rooms.New()

	c.Lobby.StartStreaming(c.lobbyHandler)

	return c
}
