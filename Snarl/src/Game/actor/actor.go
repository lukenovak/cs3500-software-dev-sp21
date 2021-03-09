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

// generates a new list of player actors
func NewPlayerList() []Actor {
	return nil
}


/* ------------ Tile occupancy functions --------------- */

func canOccupyWalkable2Steps(currTile *level.Tile, newPosition level.Position2D) bool {
	if currTile != nil && currTile.Type == level.Walkable {
		return true
	}
	return false
}