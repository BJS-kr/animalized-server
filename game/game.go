package game

import (
	"animalized/state"
	"animalized/users"
)

type Game struct {
	users.DistributableSession
	State *state.GameState
}

func New(maxUsers int) *Game {
	g := new(Game)

	g.State = state.New()
	g.MakeWithSession(maxUsers)

	return g
}

func (g *Game) JoinGame(u *users.User) error {
	return g.Join(u, g.InputChannel)
}
