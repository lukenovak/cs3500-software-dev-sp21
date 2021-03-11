package actor

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

// actor type constants
const (
	PlayerType = 0
	ghostType = 1
	zombieType = 2
)

type Actor struct {
	Type 			int
	Name 			string
	Position 		level.Position2D
	CanOccupyTile 	func(*level.Tile) bool
	MaxMoveDistance int
}


// Returns a new actor at the new position
func (actor Actor) MoveActor(newPos level.Position2D) Actor {
	return Actor{Type: actor.Type,
		Position: newPos,
		Name: actor.Name,
		CanOccupyTile: actor.CanOccupyTile,
		MaxMoveDistance: actor.MaxMoveDistance,
	}
}

// Used to generate a new actor with the default behavior (can occupy walkable tiles)
func NewWalkableActor(name string, actorType int, moveDistance int) Actor {
	return Actor{
		Type:            actorType,
		Position:        level.NewPosition2D(0, 0),
		Name:            name,
		CanOccupyTile:   canOccupyWalkable,
		MaxMoveDistance: moveDistance,
	}
}
/* ------------ Tile occupancy functions --------------- */

func canOccupyWalkable(currTile *level.Tile) bool {
	if currTile != nil && currTile.Type == level.Walkable {
		return true
	}
	return false
}