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

// Creates a deep copy of the game state
func (gs GameState) CopyGameState() GameState {

	copiedPlayers := make([]actor.Actor, 0)
	copy(copiedPlayers, gs.Players)
	copiedAdversaries := make([]actor.Actor, 0)
	copy(copiedAdversaries, gs.Adversaries)

	return GameState{
		LevelNum:    gs.LevelNum,
		Level:       gs.Level,
		Players:     copiedPlayers,
		Adversaries: copiedAdversaries,
	}
}

// Generates a new level
func (gs GameState) GenerateLevel(size level.Position2D) error {
	newLevel, err := level.NewEmptyLevel(size.Row, size.Col)
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

func (gs *GameState) MoveActorRelative(name string, relativeMove level.Position2D) {
	moveActorifExists := func(actorList []actor.Actor) {
		for i := range actorList {
			if actorList[i].Name == name {
				oldPos := actorList[i].Position
				newPos := oldPos.AddPosition(relativeMove)
				actorList[i] = actorList[i].MoveActor(newPos)
			}
		}
	}

	moveActorifExists(gs.Players)
	moveActorifExists(gs.Adversaries)
}

func (gs *GameState) MoveActorAbsolute(name string, newPosition level.Position2D) {

	moveActorIfExists := func(actorList []actor.Actor) {
		for i := range actorList {
			if actorList[i].Name == name {
				actorList[i] = actorList[i].MoveActor(newPosition)
			}
		}
	}

	moveActorIfExists(gs.Players)
	moveActorIfExists(gs.Adversaries)
}

// Searches a gamestate for an actor with the given name (which functions as an id)
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

// Mutates the current game state and removes the actor from the array
func (gs *GameState) RemoveActor(name string) {

	removeActor := func(actorList []actor.Actor) []actor.Actor {
		for i := range actorList {
			if actorList[i].Name == name {
				// reorders the array, but order doesn't matter
				actorList[len(actorList)-1], actorList[i] = actorList[i], actorList[len(actorList)-1]
				if len(actorList) == 1 {
					return []actor.Actor{}
				} else {
					return actorList[:len(actorList)-1]
				}
			}
		}
		return actorList
	}

	gs.Players = removeActor(gs.Players)
	gs.Adversaries = removeActor(gs.Adversaries)
}

// Generates a "partial game state" showing all tiles with in an n x n square as well as all actors in that square
func (gs GameState) GeneratePartialState(position level.Position2D, viewDistance int) ([][]*level.Tile, []actor.Actor) {
	// allocation
	visibleTiles := make([][]*level.Tile, viewDistance*2+1)
	for i := range visibleTiles {
		visibleTiles[i] = make([]*level.Tile, viewDistance*2+1)
	}

	var visibleActors []actor.Actor

	for partialX := 0; partialX < viewDistance; partialX++ {
		for partialY := 0; partialY < viewDistance; partialY++ {
			tilePos := level.NewPosition2D(position.Row-viewDistance+partialX, position.Col-viewDistance+partialY)
			// add the tile to the new state
			visibleTiles[partialX][partialY] = gs.Level.GetTile(tilePos)
			// if there's an adversary here, generate a new one with the relative location and move it there
			tileActor := GetActorAtPosition(gs.Players, tilePos)
			if tileActor == nil {
				tileActor = GetActorAtPosition(gs.Adversaries, tilePos)
			}
			if tileActor != nil {
				visibleActors = append(visibleActors, tileActor.MoveActor(level.NewPosition2D(partialX, partialY)))
			}
		}
	}

	return visibleTiles, nil
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

	for x := startPosn.Row; x < gameState.Level.Size.Row && x-startPosn.Row < closestDistance; x++ {
		for y := startPosn.Col; y < gameState.Level.Size.Col && y-startPosn.Col < closestDistance; y++ {
			currPos := level.NewPosition2D(x, y)
			if currPosTile := gameState.Level.GetTile(currPos); currPosTile != nil && currPosTile.Type == level.Walkable &&
				!ActorsOccupyPosition(gameState.Players, currPos) &&
				!ActorsOccupyPosition(gameState.Adversaries, currPos) &&
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

	for x := startPosn.Row; x > 0 && x-startPosn.Row < closestDistance; x-- {
		for y := startPosn.Col; y > 0 && y-startPosn.Col < closestDistance; y-- {
			currPos := level.NewPosition2D(x, y)
			if currPosTile := gameState.Level.GetTile(currPos); currPosTile != nil && currPosTile.Type == level.Walkable &&
				!ActorsOccupyPosition(gameState.Players, currPos) &&
				!ActorsOccupyPosition(gameState.Adversaries, currPos) &&
				currPos.GetManhattanDistance(startPosn) < closestDistance {
				closestTilePosn = currPos
				closestDistance = currPos.GetManhattanDistance(startPosn)
			}
		}
	}

	return closestTilePosn
}

// function to check if all actors don't occupy a position
func ActorsOccupyPosition(actors []actor.Actor, pos level.Position2D) bool {
	for _, actr := range actors {
		if actr.Position.Equals(pos) {
			return true
		}
	}
	return false
}

// gets an actor at a position if it exists. else return nil
func GetActorAtPosition(actors []actor.Actor, pos level.Position2D) *actor.Actor {
	for _, actr := range actors {
		if actr.Position.Equals(pos) {
			return &actr
		}
	}
	return nil
}
