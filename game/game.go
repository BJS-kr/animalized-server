package game

import (
	"animalized/state"
)

type Game struct {
	State *state.GameState
}

func New(maxUsers int) *Game {
	g := new(Game)

	g.State = state.New()

	return g
}
