package state

import (
	"animalized/message"
	"log/slog"
)

func determinePosition(position *message.Position, direction message.Operation_Direction) {
	switch direction {
	case message.Operation_UP:
		if position.Y-CLIENT_CELL_SIZE < 0 {
			position.Y = MAX_SPACE
		} else {
			position.Y -= CLIENT_CELL_SIZE
		}
	case message.Operation_DOWN:
		if position.Y+CLIENT_CELL_SIZE >= MAX_SPACE {
			position.Y = 0
		} else {
			position.Y += CLIENT_CELL_SIZE
		}
	case message.Operation_LEFT:
		if position.X-CLIENT_CELL_SIZE < 0 {
			position.X = MAX_SPACE
		} else {
			position.X -= CLIENT_CELL_SIZE
		}
	case message.Operation_RIGHT:
		if position.X+CLIENT_CELL_SIZE >= MAX_SPACE {
			position.X = 0
		} else {
			position.X += CLIENT_CELL_SIZE
		}
	default:
		slog.Error("unknown direction detected")
	}
}

func IsHit(pos *message.Position, hitRange *message.Operation_HitRange) bool {
	return pos.X >= hitRange.LeftBottom.GetX() && pos.X <= hitRange.RightTop.GetX() && pos.Y >= hitRange.LeftBottom.GetY() && pos.Y <= hitRange.RightTop.GetY()
}
