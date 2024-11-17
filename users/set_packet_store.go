package users

import "animalized/packet"

func (u *User) SetPacketStore(packetStore *packet.PacketStore) {
	u.packetStore = packetStore
}
