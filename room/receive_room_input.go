package room

import (
	"animalized/message"
	"animalized/packet"
)

func (r *Room) ReceiveRoomInput(roomInputChannel <-chan *message.Input) {
	for input := range roomInputChannel {
		switch input.Type {
		//join 직후 방 상태를 받는 경로
		case packet.ROOM_STATUS:
		// END는 없다. 방 상태가 게임 종료시점에 도달하면 서버가 알아서 end패킷을 룸 유저들에게 쏴줄 것
		case packet.START:

		case packet.QUIT:
		}
	}
}
