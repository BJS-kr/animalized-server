package common

import (
	"animalized/message"
	"animalized/queue"
)

type Distributable struct {
	Inputs       *queue.Queue[*message.Input]
	InputChannel chan *message.Input
	Stop         chan Signal
}

type Handler func(*message.Input) (*message.Input, error)

type Signal struct{}
