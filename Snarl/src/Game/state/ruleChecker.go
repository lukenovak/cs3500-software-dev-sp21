package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

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

		// local function to check if all actors don't occupy a position
		actorsOccupyPosition := func(actors []actor.Actor, pos level.Position2D) bool {
			for _, actr := range actors {
				if actr.Position.Equals(pos) {
					return true
				}
			}
			return false
		}

		// TODO: MAKE SURE PLAYERS DON'T OVERLAP
		validMove = validMove &&
			p.CanOccupyTile(newState.Level.GetTile(p.Position)) &&
			posnListContains(validTiles, p.Position) &&
			!actorsOccupyPosition(oldState.Players, p.Position)
	}

	return validMove
}

func IsLevelEnd(state GameState) bool {
	for _, player := range state.Players {
		position_tile := state.Level.GetTile(player.Position)
		if position_tile.Type == level.UnlockedExit {
			return true
		}
	}
	return false
}

func IsGameEnd(state GameState) bool {
	return IsLevelEnd(state)
}
