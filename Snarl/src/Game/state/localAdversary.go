package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"math/rand"
)

// client that moves ghost-type enemies
type GhostClient struct {
	Name         string
	LevelData    level.Level
	MoveDistance int
	CurrentPosn  level.Position2D
}

// calulates a move for a ghost type adversary
func (g *GhostClient) CalculateMove(playerPosns []level.Position2D, adversaryPosns []level.Position2D) Response {
	// find the move that results in the minimum distance to a player in the same room
	move := g.CurrentPosn
	roomHasPlayer := false
	walkableMoves := g.LevelData.GetWalkableTilePositions(g.CurrentPosn, g.MoveDistance)
	minDistance := g.LevelData.Size.Col * g.LevelData.Size.Row
	for _, playerPosn := range playerPosns {
		if g.LevelData.GetTile(playerPosn).RoomId == g.LevelData.GetTile(g.CurrentPosn).RoomId {
			for _, pos := range walkableMoves {
				if testDistance := pos.GetManhattanDistance(playerPosn); testDistance < minDistance {
					move = pos
					minDistance = testDistance
				}
			}
			roomHasPlayer = true
		}
	}

	// return the move found if in the same room as a player
	if roomHasPlayer {
		return Response{
			PlayerName: g.Name,
			Move:       move,
		}
	}

	// find if a teleport is possible
	hasAdjacentWall := false
	if leftTile := g.LevelData.GetTile(level.NewPosition2D(g.CurrentPosn.Row-1, g.CurrentPosn.Col)); leftTile != nil && leftTile.Type == level.Wall {
		hasAdjacentWall = true
	}
	if rightTile := g.LevelData.GetTile(level.NewPosition2D(g.CurrentPosn.Row+1, g.CurrentPosn.Col)); rightTile != nil && rightTile.Type == level.Wall {
		hasAdjacentWall = true
	}
	if upTile := g.LevelData.GetTile(level.NewPosition2D(g.CurrentPosn.Row, g.CurrentPosn.Col+1)); upTile != nil && upTile.Type == level.Wall {
		hasAdjacentWall = true
	}
	if downTile := g.LevelData.GetTile(level.NewPosition2D(g.CurrentPosn.Row, g.CurrentPosn.Col-1)); downTile != nil && downTile.Type == level.Wall {
		hasAdjacentWall = true
	}

	// pick a random wall tile to teleport to
	if hasAdjacentWall {
		var randomWallPos level.Position2D
		wallTileCount := 0
		for i, row := range g.LevelData.Tiles {
			for j, tile := range row {
				if tile != nil && tile.Type == level.Wall {
					if wallTileCount == 0 {
						randomWallPos = level.NewPosition2D(i, j)
						wallTileCount++
					} else {
						wallTileCount++
						if rand.Intn(wallTileCount) == 0 {
							randomWallPos = level.NewPosition2D(i, j)
						}
					}
				}
			}
		}

		if wallTileCount != 0 {
			return Response{
				PlayerName: g.Name,
				Move:       randomWallPos,
			}
		}
	}

	// pick a random move as a last resort
	if len(walkableMoves) > 0 {
		move = walkableMoves[rand.Intn(len(walkableMoves))]
	}

	return Response{
		PlayerName: g.Name,
		Move:       move,
	}
}

func (g *GhostClient) UpdatePosition(d level.Position2D) {
	g.CurrentPosn = d
}

func (g *GhostClient) GetName() string {
	return g.Name
}

// client that moves zombie-type enemies
type ZombieClient struct {
	Name         string
	LevelData    level.Level
	MoveDistance int
	CurrentPosn  level.Position2D
}

func (z *ZombieClient) UpdatePosition(d level.Position2D) {
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
		if z.LevelData.GetTile(posn).RoomId == z.LevelData.GetTile(z.CurrentPosn).RoomId {
			roomHasPlayer = true
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
