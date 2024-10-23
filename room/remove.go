package room

func (rs *Rooms) RemoveRoom(room *Room) {
	for k, v := range rs.RoomMap {
		if v == room {
			delete(rs.RoomMap, k)
			return
		}
	}
}
