package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

// client that moves ghost-type enemies
type GhostClient struct {
	Name string
}

func (g GhostClient) CalculateMove(playerPosns []level.Position2D) Response {
	return Response{
		PlayerName: g.Name,
		Move:       level.Position2D{0, 0},
	}
}

// client that moves zombie-type enemies
type ZombieClient struct {
	Name string
	LevelData level.Level
	MoveDistance int
	CurrentPosn level.Position2D
}

func (z ZombieClient) CalculateMove(playerPosns []level.Position2D) Response {
	return Response{
		PlayerName: z.Name,
		Move:       level.Position2D{0, 0},
	}
}

