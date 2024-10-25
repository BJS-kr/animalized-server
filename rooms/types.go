package rooms

import (
	"animalized/users"
)

type RoomName string
type RoomStatus int

type Room struct {
	users.DistributableUsers
	status RoomStatus
}

type Rooms struct {
	NameMap map[RoomName]*Room
}
