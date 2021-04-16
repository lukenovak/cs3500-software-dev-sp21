package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

// UserClient represents an interactable user client
type UserClient interface {
	// RegisterClient Used to initialize the client
	RegisterClient() (actor.Actor, error)

	// SendPartialState Sends a new state to the player
	SendPartialState(layout [][]*level.Tile, actors []actor.Actor, pos level.Position2D) error

	// SendMessage Sends a message to the player, used to acknowledge player moves
	SendMessage(message string, pos level.Position2D) error

	// GetInput Waits for a player input then returns the player's action after input
	GetInput() Response

	// GetName returns the unique name Id for this client
	GetName() string
}

// Response represents an internal response to a prompt to get user input.
type Response struct {
	PlayerName string
	Move       level.Position2D
	Actions    map[string]interface{}
}
