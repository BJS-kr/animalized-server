package state

import (
	"animalized/packet"
	"log/slog"
)

type Position struct {
	X int32
	Y int32
}

func (p *Position) determinePosition(direction int32) {
	switch direction {
	case packet.UP:
		if p.Y-1 < 0 {
			p.Y = MAP_SIZE - 1
		} else {
			p.Y--
		}
	case packet.DOWN:
		if p.Y+1 >= MAP_SIZE {
			p.Y = 0
		} else {
			p.Y++
		}
	case packet.LEFT:
		if p.X-1 < 0 {
			p.X = MAP_SIZE - 1
		} else {
			p.X--
		}
	case packet.RIGHT:
		if p.X+1 >= MAP_SIZE {
			p.X = 0
		} else {
			p.X++
		}
	default:
		slog.Error("unknown direction detected")
	}
}
