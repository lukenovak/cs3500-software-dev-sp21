package state

import "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"

// main rule checking function
func IsValidMove(oldState GameState, newState GameState) bool {
	validMove := true
	for _, p := range newState.Players {
		oldPos := oldState.GetActor(p.Name).Position
		validTiles := oldState.Level.GetWalkableTilePositions(oldPos, p.MaxMoveDistance)

		// generic contains functions for posns (no generics)
		posnListContains := func(posnList []level.Position2D, searchPosn level.Position2D) bool {
			for _, listPosn := range posnList {
				if listPosn.Equals(searchPosn) {
					return true
				}
			}
			return false
		}

		// TODO: MAKE SURE PLAYERS DON'T OVERLAP
		validMove = validMove &&
			p.CanOccupyTile(newState.Level.GetTile(p.Position)) &&
			posnListContains(validTiles, p.Position)
	}

	return validMove
}

func IsLevelEnd(state GameState) bool {
	return false
}

func IsGameEnd(state GameState) bool {
	return false
}
