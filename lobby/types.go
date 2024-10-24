package lobby

import (
	"animalized/room"
	"animalized/users"
)

type Lobby struct {
	users.DistributableUsers
	rooms *room.Rooms
}
