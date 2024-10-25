package rooms

import (
	"animalized/game"
	"animalized/users"
)

type RoomName string
type RoomStatus int

type Room struct {
	users.DistributableUsers
	Status RoomStatus
	Game   *game.Game
}

type Rooms struct {
	NameMap map[RoomName]*Room
}
