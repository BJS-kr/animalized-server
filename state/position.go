package state

import (
	"animalized/message"
	"log/slog"
)

type Position struct {
	X int32
	Y int32
}

func (p *Position) determinePosition(direction message.Operation_Direction) {
	switch direction {
	case message.Operation_UP:
		if p.Y-1 < 0 {
			p.Y = MAP_SIZE - 1
		} else {
			p.Y--
		}
	case message.Operation_DOWN:
		if p.Y+1 >= MAP_SIZE {
			p.Y = 0
		} else {
			p.Y++
		}
	case message.Operation_LEFT:
		if p.X-1 < 0 {
			p.X = MAP_SIZE - 1
		} else {
			p.X--
		}
	case message.Operation_RIGHT:
		if p.X+1 >= MAP_SIZE {
			p.X = 0
		} else {
			p.X++
		}
	default:
		slog.Error("unknown direction detected")
	}
}
