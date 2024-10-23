package lobby

import (
	"animalized/message"
	"animalized/queue"
	"animalized/room"
	"animalized/users"
)

type Lobby struct {
	Users  *users.Users
	Rooms  *room.Rooms
	Inputs *queue.Queue[*message.Input]
}
