package lobby

import (
	"animalized/room"
)

func New(max int) *Lobby {
	l := new(Lobby)
	rs := new(room.Rooms)

	rs.NameMap = make(map[room.RoomName]*room.Room)

	l.Make()
	l.Users.Max = max
	l.rooms = rs

	l.StartStreaming(l.handler)

	return l
}
