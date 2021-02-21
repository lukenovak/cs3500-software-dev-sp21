package actor

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"io"
)

const noAction = 0

type Actor struct {
	Type int
	Position level.Position2D
	Input io.Reader
	Output io.Writer
}

func (actor Actor) MoveActor(action int) error {
	return fmt.Errorf("Not yet implemented")
}