package game

import (
	"animalized/common"
	"animalized/message"
	"animalized/state"
	"animalized/users"
	"time"

	"strconv"

	"math/rand"
)

const TERRAINS_COUNT = 40

type Game struct {
	users.DistSession
	AttackDedup map[int32]bool
	State       *state.GameState
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

func (g *Game) Init(handler common.Handler, tickRate time.Duration) {
	g.initTerrains()
	g.AttackDedup = make(map[int32]bool)
	g.StartStreaming(handler, tickRate)
}

func (g *Game) initTerrains() {
	// proto의 기본값을 제외하기 위해 0을 건너뛰기 위해 1 추가
	terrains := make([]*message.Terrain, TERRAINS_COUNT+1)
	reservedPositions := map[string]bool{
		"00": true,
	}
	generatedPositions := make([]*message.Position, 0, TERRAINS_COUNT)

	for len(generatedPositions) < TERRAINS_COUNT {
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
		Type:  message.Terrain_ROCK,
		State: message.Terrain_DESTROYED,
	}

	for i := 1; i <= TERRAINS_COUNT; i++ {
		pos := generatedPositions[i-1]
		pos.X = pos.X * state.CLIENT_CELL_SIZE
		pos.Y = pos.Y * state.CLIENT_CELL_SIZE

		terrains[i] = &message.Terrain{
			Type:     message.Terrain_ROCK,
			State:    message.Terrain_SOLID,
			Position: pos,
		}
	}

	g.State.Terrains = terrains
}
