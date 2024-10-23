package room

import "animalized/users"

type RoomName string
type RoomStatus int

type Room struct {
	users  *users.Users
	status RoomStatus
}

type Rooms struct {
	RoomMap map[RoomName]*Room
}
