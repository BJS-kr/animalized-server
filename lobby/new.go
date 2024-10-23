package lobby

import (
	"animalized/message"
	"animalized/queue"
	"animalized/room"
	"animalized/users"
)

func New(max int) *Lobby {
	l := new(Lobby)
	rs := new(room.Rooms)
	us := new(users.Users)

	rs.RoomMap = make(map[room.RoomName]*room.Room)
	us.Max = max

	l.users = us
	l.rooms = rs
	l.inputs = queue.New[*message.Input]()

	return l
}
