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

		validTiles := currState.Level.GetWalkableTilePositions(movingActor.Position, movingActor.MaxMoveDistance)

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

func IsLevelEnd(state GameState) bool {
	var isEnd = true
	for _, player := range state.Players {
		position_tile := state.Level.GetTile(player.Position)
		if position_tile == nil {
			return false
		}
		if position_tile.Item.Type == level.UnlockedExit {
			isEnd = true && isEnd
		} else {
			isEnd = false
		}
	}
	return isEnd
}

func IsGameEnd(state GameState, maxLevel int) bool {
	return IsLevelEnd(state) && state.LevelNum == maxLevel
}
