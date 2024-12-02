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

	g.State = state.New()
	g.MakeWithUsers(maxUsers)

	return g
}

func (g *Game) JoinGame(u *users.User) error {
	return g.Join(u, g.InputChannel)
}
