package game

import (
	"animalized/state"
	"animalized/users"
)

type Game struct {
	users.DistributableUsers
	State *state.GameState
}

func New(maxUsers int) *Game {
	g := new(Game)

	g.MakeWithUsers(maxUsers)
	g.State = state.New()

	return g
}
