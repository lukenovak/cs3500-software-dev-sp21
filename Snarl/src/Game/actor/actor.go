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
	Type int
	Id int // must be unique
	Position level.Position2D
}

// Returns a new actor at the new position
func (actor Actor) MoveActor(newPos level.Position2D) Actor {
	return Actor{Type: actor.Type, Id: actor.Id, Position: newPos}
}

// generates a new list of player actors
func NewPlayerList() []Actor {
	return nil
}
