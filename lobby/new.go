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

	rs.Rooms = make(map[room.RoomName]*room.Room)
	us.Max = max

	l.Users = us
	l.Rooms = rs
	l.Inputs = queue.New[*message.Input]()

	return l
}
