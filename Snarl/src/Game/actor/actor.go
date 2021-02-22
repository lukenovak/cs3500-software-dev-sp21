package actor

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"io"
)

// actor type constants
const (
	PlayerType = 0
	ghostType = 1
	zombieType = 2
)


type Actor struct {
	Type int
	Id int
	Position level.Position2D
	Input io.Reader
	Output io.Writer
}

func (actor Actor) MoveActor(action int) error {
	return fmt.Errorf("Not yet implemented")
}

// generates a new list of player actors
func NewPlayerList() []Actor {
	return nil
}
