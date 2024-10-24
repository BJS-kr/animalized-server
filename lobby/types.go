package lobby

import (
	"animalized/common"
	"animalized/room"
)

type Lobby struct {
	common.Base
	rooms *room.Rooms
}
