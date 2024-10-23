package room

import "animalized/users"

type RoomName string
type RoomStatus int

type Room struct {
	participantsLimit int
	participants      []*users.User
	status            RoomStatus
}

type Rooms struct {
	Rooms map[RoomName]*Room
}
