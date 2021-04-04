package state

import "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"

// client controller that determines an Adversary's moves
type AdversaryClient interface {
	// Returns a Response with the best relative move
	CalculateMove(playerPosns []level.Position2D, adversaryPositions []level.Position2D) Response

	// Sends an update to the AdversaryClient if its adversary has been moved to a new position
	UpdatePosition(d level.Position2D)

	// Gets the name to search for the adversary
	GetName() string
}

type ExampleAdvClient struct {
	Name     	  	string
	LevelData  		level.Level
	MoveDistance 	int
	CurrentPosn  	level.Position2D
}
