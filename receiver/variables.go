package receiver

import (
	"animalized/message"
	"animalized/queue"
)

var Inputs = queue.New[*message.Input]()
