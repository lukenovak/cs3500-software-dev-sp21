package actor

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

// actor type constants
const (
	PlayerType = 0
	GhostType  = 1
	ZombieType = 2
)

type Actor struct {
	Type            int
	Name            string
	Position        level.Position2D
	CanOccupyTile   func(*level.Tile) bool
	MaxMoveDistance int
}

// Returns a new actor at the new position
func (actor Actor) MoveActor(newPos level.Position2D) Actor {
	return Actor{Type: actor.Type,
		Position:        newPos,
		Name:            actor.Name,
		CanOccupyTile:   actor.CanOccupyTile,
		MaxMoveDistance: actor.MaxMoveDistance,
	}
}

// Constructs a new actor with the default behavior (can occupy walkable tiles)
func NewWalkableActor(name string, actorType int, moveDistance int) Actor {
	return Actor{
		Type:            actorType,
		Position:        level.NewPosition2D(-1, -1),
		Name:            name,
		CanOccupyTile:   canOccupyWalkable,
		MaxMoveDistance: moveDistance,
	}
}

// Constructs an actor with the adversary behavior (cannot walk on doors)
func NewAdversaryActor(adversaryType int, name string, moveDistance int) Actor {
	return Actor{
		Type:            adversaryType,
		Name:            name,
		Position:        level.NewPosition2D(-1, -1),
		CanOccupyTile:   adversaryOccupy,
		MaxMoveDistance: moveDistance,
	}
}

/* ------------ Tile occupancy functions --------------- */

// generic behavor lambda
func canOccupyWalkable(currTile *level.Tile) bool {
	return currTile != nil && (currTile.Type == level.Walkable || currTile.Type == level.Door || currTile.Type == level.LockedExit || currTile.Type == level.UnlockedExit)
}

func adversaryOccupy(currTile *level.Tile) bool {
	return currTile != nil && currTile.Type == level.Walkable
}
