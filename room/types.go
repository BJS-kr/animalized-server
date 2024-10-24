package room

import (
	"animalized/common"
)

type RoomName string
type RoomStatus int

type Room struct {
	common.Base
	status RoomStatus
}

type Rooms struct {
	NameMap map[RoomName]*Room
}
