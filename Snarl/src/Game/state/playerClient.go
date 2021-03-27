package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

type UserClient interface {
	// Used to initialize the client
	RegisterClient() (actor.Actor, error)

	// Sends a new state to the player
	SendPartialState(tiles [][]*level.Tile, actors []actor.Actor) error

	// Sends a message to the player (used for invalid moves);
	SendMessage(message string) error

	// Waits for a player input then returns the player's action after input
	GetInput() Response

	// Used on startup- gets the unique name Id for this client
	GetName() string
}

type Response struct {
	PlayerName string
	Move       level.Position2D
	Actions    map[string]interface{}
}
