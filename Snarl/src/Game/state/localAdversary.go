package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"math/rand"
)

// client that moves ghost-type enemies
type GhostClient struct {
	Name string
	LevelData level.Level
	MoveDistance int
	CurrentPosn level.Position2D
}

func (g *GhostClient) CalculateMove(playerPosns []level.Position2D, adversaryPosns []level.Position2D) Response {
	return Response{
		PlayerName: g.Name,
		Move:       level.Position2D{0, 0},
	}
}

func (g *GhostClient) UpdatePosition(d level.Position2D)  {
	g.CurrentPosn = d
}

func (g *GhostClient) GetName() string {
	return g.Name
}

// client that moves zombie-type enemies
type ZombieClient struct {
	Name string
	LevelData level.Level
	MoveDistance int
	CurrentPosn level.Position2D
}

func (z *ZombieClient) UpdatePosition(d level.Position2D)  {
	z.CurrentPosn = d
}

func (z *ZombieClient) GetName() string {
	return z.Name
}

func (z *ZombieClient) CalculateMove(playerPosns []level.Position2D, adversaryPosns []level.Position2D) Response {
	move := z.CurrentPosn
	roomHasPlayer := false
	validMoves := z.LevelData.GetWalkableTilePositions(z.CurrentPosn, z.MoveDistance)
	for _, posn := range playerPosns {
		if roomHasPlayer || z.LevelData.GetTile(posn).RoomId == z.LevelData.GetTile(z.CurrentPosn).RoomId {
			minDistance := z.LevelData.Size.Col * z.LevelData.Size.Row
			for _, pos := range validMoves {
				if pos.GetManhattanDistance(posn) < minDistance && z.LevelData.GetTile(pos).Type == level.Walkable {
					minDistance = pos.GetManhattanDistance(posn)
					move = pos
				}
			}
			break
		}
	}
	if !roomHasPlayer {
		// if there's no player in the room, pick a random move
		move = validMoves[rand.Intn(len(validMoves))]
	}
	return Response{
		PlayerName: z.Name,
		Move:       move,
	}
}

