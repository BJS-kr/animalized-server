package packet

import "animalized/message"

func IsInit(input *message.Input) bool {
	_, ok := input.Kind.(*message.Input_Init)

	return ok
}
