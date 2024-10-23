package lobby

import (
	"animalized/message"
	"animalized/queue"
	"animalized/room"
	"animalized/users"
)

type Lobby struct {
	users  *users.Users
	rooms  *room.Rooms
	inputs *queue.Queue[*message.Input]
}
