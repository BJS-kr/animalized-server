package room

func (rs *Rooms) RemoveRoom(room *Room) {
	for k, v := range rs.NameMap {
		if v == room {
			room.StopStreaming()
			delete(rs.NameMap, k)
			return
		}
	}
}
