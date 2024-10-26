package rooms

func New() *Rooms {
	rs := new(Rooms)
	rs.NameMap = make(map[RoomName]*Room)

	return rs
}
