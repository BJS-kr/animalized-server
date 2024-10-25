package game

import "animalized/state"

func New(maxUsers int) *Game {
	g := new(Game)

	g.MakeWithUsers(maxUsers)
	g.State = state.New()

	return g
}
