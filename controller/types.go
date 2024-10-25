package controller

import (
	"animalized/lobby"
	"animalized/rooms"
)

type Controller struct {
	*lobby.Lobby
	*rooms.Rooms
}
