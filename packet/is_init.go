package packet

import "animalized/message"

func IsInitPacket(input *message.Input) bool {
	return input.GetType() == INIT
}
