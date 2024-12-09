package controller

import "time"

const (
	LOBBY_TICK_RATE = 100 * time.Millisecond
	ROOM_TICK_RATE  = 100 * time.Millisecond
	GAME_TICK_RATE  = 3 * time.Millisecond
)
