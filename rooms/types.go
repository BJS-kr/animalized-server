package rooms

import (
	"animalized/game"
	"animalized/message"
	"animalized/users"
)

type RoomName string
type RoomStatus int

type Room struct {
	users.DistributableUsers
	Status message.RoomState_RoomStatusType
	Game   *game.Game
}

type Rooms struct {
	NameMap map[RoomName]*Room
}
