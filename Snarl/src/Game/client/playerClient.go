package client

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

type UserClient interface {
	// Sends a new state to the player
	SendPartialState([][]*level.Tile, []actor.Actor) error

	// Sends a message to the player (used for invalid moves);
	SendMessage(string) error

	// Waits for a player input then returns the player's action after input
	GetInput() []ClientResponse
}

type ClientResponse struct {
	PlayerId   int
	PlayerName string
	Move       level.Position2D
	Actions    map[string]interface{}
}
