package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

type GameState struct {
	LevelNum int
	Level *level.Level
	Players []actor.Actor
	Adversaries []actor.Actor
}

// Creates a new game state with the players and adversaries moved to new positions
func (gs GameState) CreateUpdatedGameState(updatedPlayers []actor.Actor, updatedAdversaries []actor.Actor) *GameState {
	return &GameState{
		LevelNum: gs.LevelNum,
		Level: gs.Level,
		Players: updatedPlayers,
		Adversaries: updatedAdversaries,
	}
}

// Generates a new level
func (gs GameState) GenerateLevel(size level.Position2D) error {
	newLevel, err := level.NewEmptyLevel(size.X, size.Y)
	if err != nil {
		return err
	}
	gs.Level = &newLevel
	return nil
}

func (gs *GameState) SpawnActor(actorToSpawn actor.Actor, initialPosition level.Position2D) {
	positionedActor := actorToSpawn.MoveActor(initialPosition)
	if actorToSpawn.Type == actor.PlayerType {
		gs.Players = append(gs.Players, positionedActor)
	} else {
		gs.Players = append(gs.Adversaries, positionedActor)
	}
}

func (gs GameState) CheckVictory() bool {
	for _, player := range gs.Players {
		if tile := gs.Level.GetTile(player.Position); tile != nil && tile.Type == level.UnlockedExit {
			return true
		}
	}
	return false
}

func (gs *GameState) UnlockExits() {
	for _, exit := range gs.Level.Exits {
		exit.Type = level.UnlockedExit
	}
}

func initGameState(firstLevel level.Level, players []actor.Actor) *GameState {
	gs := &GameState{
		Level:      &firstLevel,
		//Adversaries: GenerateAdversaries(numPlayers),
	}
	placePlayers(gs, players)
	return gs
}

func placePlayers(gameState *GameState, players []actor.Actor)  {
	for _, player := range players {
		gameState.SpawnActor(player, getTopLeftUnoccupiedWalkable(*gameState, level.NewPosition2D(0, 0)))
	}
}

func getTopLeftUnoccupiedWalkable(gameState GameState, startPosn level.Position2D) level.Position2D {
	closestDistance := gameState.Level.Size.GetManhattanDistance(startPosn)
	closestTilePosn := gameState.Level.Size

	for x := startPosn.X; x < gameState.Level.Size.X && x - startPosn.X < closestDistance; x++ {
		for y := startPosn.Y; y < gameState.Level.Size.Y && x - startPosn.X < closestDistance; y++ {
			currPos := level.NewPosition2D(x, y)
			if currPosTile := gameState.Level.GetTile(currPos);
			currPosTile != nil && currPosTile.Type == level.Walkable && !isOccupiedByPlayer(gameState, currPos) &&
				currPos.GetManhattanDistance(startPosn) < closestDistance {
				closestTilePosn = currPos
				closestDistance = currPos.GetManhattanDistance(startPosn)
			}
		}
	}

	return closestTilePosn
}

func isOccupiedByPlayer(state GameState, posn level.Position2D) bool {
	occupied := false
	for i := 0; i < len(state.Players) && !occupied; i++ {
		occupied = occupied || state.Players[i].Position.Equals(posn)
	}
	return occupied
}
