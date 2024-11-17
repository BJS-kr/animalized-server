package common

import (
	"animalized/message"
)

type Handler func(*message.Input) (*message.Input, error)

type Signal struct{}
