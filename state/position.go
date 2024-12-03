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
		if p.Y-CLIENT_CELL_SIZE < 0 {
			p.Y = MAX_SPACE
		} else {
			p.Y -= CLIENT_CELL_SIZE
		}
	case message.Operation_DOWN:
		if p.Y+CLIENT_CELL_SIZE >= MAX_SPACE {
			p.Y = 0
		} else {
			p.Y += CLIENT_CELL_SIZE
		}
	case message.Operation_LEFT:
		if p.X-CLIENT_CELL_SIZE < 0 {
			p.X = MAX_SPACE
		} else {
			p.X -= CLIENT_CELL_SIZE
		}
	case message.Operation_RIGHT:
		if p.X+CLIENT_CELL_SIZE >= MAX_SPACE {
			p.X = 0
		} else {
			p.X += CLIENT_CELL_SIZE
		}
	default:
		slog.Error("unknown direction detected")
	}
}

func (p *Position) IsHit(hitRange *message.Operation_HitRange) bool {
	return p.X >= hitRange.LeftBottom.GetX() && p.X <= hitRange.RightTop.GetX() && p.Y >= hitRange.LeftBottom.GetY() && p.Y <= hitRange.RightTop.GetY()
}
