package actor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"golang.org/x/image/colornames"
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
	RenderedObj		fyne.CanvasObject
}

// Returns a new actor at the new position
func (actor Actor) MoveActor(newPos level.Position2D) Actor {
	return Actor{Type: actor.Type,
		Position:        newPos,
		Name:            actor.Name,
		CanOccupyTile:   actor.CanOccupyTile,
		MaxMoveDistance: actor.MaxMoveDistance,
		RenderedObj:     actor.RenderedObj,
	}
}

// Constructs a new actor with the default behavior (can occupy walkable tiles)
func NewPlayerActor(name string, actorType int, moveDistance int) Actor {
	return Actor{
		Type:            actorType,
		Position:        level.NewPosition2D(-1, -1),
		Name:            name,
		CanOccupyTile:   canOccupyWalkable,
		MaxMoveDistance: moveDistance,
		RenderedObj: 	 canvas.NewCircle(colornames.Cornflowerblue),
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
		RenderedObj: 	 canvas.NewCircle(colornames.Crimson),
	}
}

/* ------------ Tile occupancy functions --------------- */

// generic behavor lambda
func canOccupyWalkable(currTile *level.Tile) bool {
	return currTile != nil && (currTile.Type == level.Walkable || currTile.Type == level.Door)
}

func adversaryOccupy(currTile *level.Tile) bool {
	return currTile != nil && currTile.Type == level.Walkable
}
