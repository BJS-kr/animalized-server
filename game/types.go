package game

import (
	"animalized/state"
	"animalized/users"
)

type Game struct {
	users.DistributableUsers
	State *state.GameState
}
