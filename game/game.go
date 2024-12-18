package game

import (
	"animalized/message"
	"animalized/state"
	"animalized/users"
	"strconv"

	"math/rand"
)

const TerrainsCount = 20

type Game struct {
	users.DistSession
	State *state.GameState
}

func New(maxUsers int) *Game {
	g := new(Game)

	g.State = state.New()
	g.Make(maxUsers)

	return g
}

func (g *Game) JoinGame(u *users.User) error {
	return g.Join(u, g.Receiver)
}

func (g *Game) InitTerrains() {
	// proto의 기본값을 제외하기 위해 0을 건너뛰기 위해 1 추가
	terrains := make([]*message.Terrain, TerrainsCount+1)
	reservedPositions := map[string]bool{
		"00": true,
	}
	generatedPositions := make([]*message.Position, 0, TerrainsCount)

	for len(generatedPositions) < TerrainsCount {
		x := rand.Int31n(state.MAP_SIZE)
		y := rand.Int31n(state.MAP_SIZE)
		xStr := strconv.Itoa(int(x))
		yStr := strconv.Itoa(int(y))

		if _, ok := reservedPositions[xStr+yStr]; !ok {
			generatedPositions = append(generatedPositions, &message.Position{
				X: x,
				Y: y,
			})

			reservedPositions[xStr+yStr] = true
		}
	}

	terrains[0] = &message.Terrain{
		Type:  message.TerrainType_ROCK,
		State: message.TerrainState_DESTROYED,
	}

	for i := 1; i <= TerrainsCount; i++ {
		terrains[i] = &message.Terrain{
			Type:     message.TerrainType_ROCK,
			State:    message.TerrainState_SOLID,
			Position: generatedPositions[i-1],
		}
	}

	g.State.Terrains = terrains
}
