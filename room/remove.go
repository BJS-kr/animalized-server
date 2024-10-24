package room

func (rs *Rooms) RemoveRoom(room *Room) {
	for k, v := range rs.NameMap {
		if v == room {
			close(room.Stop)
			delete(rs.NameMap, k)
			return
		}
	}
}
