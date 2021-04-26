package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

// main rule checking function
func IsValidMove(currState GameState, movingActorName string, relativeMove level.Position2D) bool {
	validMove := true
	movingActor := currState.GetActor(movingActorName)
	if movingActor == nil {
		return false
	} else {

		validTiles := currState.Level.GetTraversableTilePositions(movingActor.Position, movingActor.MaxMoveDistance)

		// generic contains functions for posns (no generics)
		posnListContains := func(posnList []level.Position2D, searchPosn level.Position2D) bool {
			for _, listPosn := range posnList {
				if listPosn.Equals(searchPosn) {
					return true
				}
			}
			return false
		}

		newPosition := movingActor.Position.AddPosition(relativeMove)

		validMove = validMove &&
			movingActor.CanOccupyTile(currState.Level.GetTile(newPosition)) &&
			posnListContains(validTiles, newPosition) &&
			!(ActorsOccupyPosition(currState.Players, newPosition) && GetActorAtPosition(currState.Players, newPosition).Name != movingActorName)
	}
	return validMove
}

// Returns true if a player is standing on an unlocked exit and is the only player remaining
func IsLevelEnd(state GameState) bool {
	var playerTile *level.Tile
	if len(state.Players) != 0 {
		playerTile = state.Level.GetTile(state.Players[0].Position)
	}
	return len(state.Players) == 0 || (len(state.Players) == 1 && playerTile.Item != nil && playerTile.Item.Type == level.UnlockedExit)
}

func IsGameEnd(state GameState, maxLevel int) bool {
	return IsLevelEnd(state) && state.LevelNum == maxLevel
}
