package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/client"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

type GameState struct {
	LevelNum      int
	Level         *level.Level
	Players       []actor.Actor
	PlayerClients []client.UserClient
	Adversaries   []actor.Actor
}

// Creates a new game state with the players and adversaries moved to new positions
func (gs GameState) CreateUpdatedGameState(updatedPlayers []actor.Actor, updatedAdversaries []actor.Actor) *GameState {
	return &GameState{
		LevelNum:    gs.LevelNum,
		Level:       gs.Level,
		Players:     updatedPlayers,
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

// Adds an actor to the game state at the given position
// Adds an actor to the game state at the given position
func (gs *GameState) SpawnActor(actorToSpawn actor.Actor, initialPosition level.Position2D) {
	positionedActor := actorToSpawn.MoveActor(initialPosition)
	if actorToSpawn.Type == actor.PlayerType {
		gs.Players = append(gs.Players, positionedActor)
	} else {
		gs.Adversaries = append(gs.Adversaries, positionedActor)
	}
}

// Checks to see if the game has been won. If it has, returns true.
func (gs GameState) CheckVictory() bool {
	for _, player := range gs.Players {
		if tile := gs.Level.GetTile(player.Position); tile != nil && tile.Type == level.UnlockedExit {
			return true
		}
	}
	return false
}

// Changes all exits from locked exits to unlocked exits
func (gs *GameState) UnlockExits() {
	for _, exit := range gs.Level.Exits {
		exit.Type = level.UnlockedExit
	}
}

// Searches a gamestate for an actor with the given name (which functions as an id
func (gs GameState) GetActor(name string) *actor.Actor {

	findActor := func(actorList []actor.Actor) *actor.Actor {
		for _, player := range gs.Players {
			if player.Name == name {
				return &player
			}
		}
		return nil
	}

	if player := findActor(gs.Players); player != nil {
		return player
	}

	return findActor(gs.Adversaries)
}

/* ---------------------------- Internal Use Functions ------------------------------------- */

// Creates the initial game state. For internal use
func initGameState(firstLevel level.Level, players []actor.Actor, adversaries []actor.Actor) *GameState {
	gs := &GameState{
		Level: &firstLevel,
		//Adversaries: GenerateAdversaries(numPlayers),
	}
	placeActors(gs, players, getTopLeftUnoccupiedWalkable, level.NewPosition2D(0, 0))
	placeActors(gs, adversaries, getBottomRightUnoccupiedWalkable, firstLevel.Size)
	return gs
}

// places the players at the top left of the level
func placeActors(gameState *GameState, actors []actor.Actor,
	placementFunc func(GameState, level.Position2D) level.Position2D,
	placementStart level.Position2D) {
	for _, actor := range actors {
		gameState.SpawnActor(actor, placementFunc(*gameState, placementStart))
	}
}

// gets the top left most unoccupied tile relative to the start position
func getTopLeftUnoccupiedWalkable(gameState GameState, startPosn level.Position2D) level.Position2D {
	closestDistance := gameState.Level.Size.GetManhattanDistance(startPosn)
	closestTilePosn := gameState.Level.Size

	for x := startPosn.X; x < gameState.Level.Size.X && x-startPosn.X < closestDistance; x++ {
		for y := startPosn.Y; y < gameState.Level.Size.Y && y-startPosn.Y < closestDistance; y++ {
			currPos := level.NewPosition2D(x, y)
			if currPosTile := gameState.Level.GetTile(currPos); currPosTile != nil && currPosTile.Type == level.Walkable &&
				!isOccupiedByActor(gameState, currPos) &&
				currPos.GetManhattanDistance(startPosn) < closestDistance {
				closestTilePosn = currPos
				closestDistance = currPos.GetManhattanDistance(startPosn)
			}
		}
	}

	return closestTilePosn
}

// gets the bottom right most unoccupied tile relative to the start position
func getBottomRightUnoccupiedWalkable(gameState GameState, startPosn level.Position2D) level.Position2D {
	closestDistance := level.NewPosition2D(0, 0).GetManhattanDistance(startPosn)
	closestTilePosn := level.NewPosition2D(0, 0)

	for x := startPosn.X; x > 0 && x-startPosn.X < closestDistance; x-- {
		for y := startPosn.Y; y > 0 && y-startPosn.Y < closestDistance; y-- {
			currPos := level.NewPosition2D(x, y)
			if currPosTile := gameState.Level.GetTile(currPos); currPosTile != nil && currPosTile.Type == level.Walkable &&
				!isOccupiedByActor(gameState, currPos) &&
				currPos.GetManhattanDistance(startPosn) < closestDistance {
				closestTilePosn = currPos
				closestDistance = currPos.GetManhattanDistance(startPosn)
			}
		}
	}

	return closestTilePosn
}

// checks to see that a tile is occupied by an actor from the given list
func isOccupiedByActor(gameState GameState, posn level.Position2D) bool {
	occupied := false
	for i := 0; i < len(gameState.Players) && !occupied; i++ {
		occupied = occupied || gameState.Players[i].Position.Equals(posn)
	}
	for i := 0; i < len(gameState.Adversaries) && !occupied; i++ {
		occupied = occupied || gameState.Adversaries[i].Position.Equals(posn)
	}
	return occupied
}
