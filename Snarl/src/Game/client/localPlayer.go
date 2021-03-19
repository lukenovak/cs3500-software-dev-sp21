package client

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

// struct for a local player
type LocalPlayer struct {
	Name string
}

// TODO: Include this functionality in a future Milestone
func (player LocalPlayer) SendPartialState([][]*level.Tile, []actor.Actor) error {
	return nil
}

// TODO: Include this functionality in a future Milestone
func (player LocalPlayer) SendMessage(string) error {
	return nil
}

// TODO: Include this functionality in a future Milestone
func (player LocalPlayer) GetInput() Response {
	return Response{}
}

func (player LocalPlayer) GetName() string {
	return player.Name
}
