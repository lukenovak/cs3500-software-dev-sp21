package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

// GameState contains all of the current information necessary to track the state of one game level.
type GameState struct {
	LevelNum       int            // The level number of the active level
	Level          *level.Level   // A pointer to the current active level (also contains items)
	Players        []actor.Actor  // The current list of Player actors (also contains positions)
	Adversaries    []actor.Actor  // The current list of Adversary actors (also contains positions)
}

// SpawnActor adds the given actor to the game state at the given position
func (gs *GameState) SpawnActor(actorToSpawn actor.Actor, initialPosition level.Position2D) {
	positionedActor := actorToSpawn.MoveActor(initialPosition)
	if actorToSpawn.Type == actor.PlayerType {
		gs.Players = append(gs.Players, positionedActor)
	} else {
		gs.Adversaries = append(gs.Adversaries, positionedActor)
	}
}

// UnlockExits changes all exits from locked exits to unlocked exits
func (gs *GameState) UnlockExits() {
	gs.Level.UnlockExits()
}

// MoveActorRelative moves the actor to the space given a cooordinate represented as the new position relative to the actor's current
// position
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

// MoveActorAbsolute moves the actor with the given name to the given Position2D, in absolute coordinates from the
// level origin
func (gs *GameState) MoveActorAbsolute(name string, newPosition level.Position2D) {

	// local function searches for an actor, and if one matches the given name, it moves the actor
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

// GetActor Searches a gamestate for an actor with the given name (which functions as an id)
func (gs GameState) GetActor(name string) *actor.Actor {

	findActor := func(actorList []actor.Actor) *actor.Actor {
		for _, player := range actorList {
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

// RemoveActor Mutates the current game state and removes the actor from the array
func (gs *GameState) RemoveActor(name string) {

	removeActor := func(actorList []actor.Actor) []actor.Actor {
		var newList []actor.Actor
		for i := range actorList {
			if actorList[i].Name != name {
				newList = append(newList, actorList[i])
			}
		}
		return newList
	}

	gs.Players = removeActor(gs.Players)
	gs.Adversaries = removeActor(gs.Adversaries)
}

// GeneratePartialState returns a "partial game state" in the form of a tuple with a Tile layout, a list of Actors in
// that are in that space, and a Position2D containing the position it is being called from
func (gs GameState) GeneratePartialState(position level.Position2D, viewDistance int) ([][]*level.Tile, []actor.Actor, level.Position2D) {
	// allocation
	visibleTiles := make([][]*level.Tile, viewDistance*2+1)
	for i := range visibleTiles {
		visibleTiles[i] = make([]*level.Tile, viewDistance*2+1)
	}

	var visibleActors []actor.Actor

	for partialX := 0; partialX < viewDistance*2+1; partialX++ {
		for partialY := 0; partialY < viewDistance*2+1; partialY++ {
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

	return visibleTiles, visibleActors, position
}

/* ---------------------------- External Use Functions ------------------------------------- */

// ActorsOccupyPosition checks all actors to see if one occupies the position. If it does, returns true. Else returns false
func ActorsOccupyPosition(actors []actor.Actor, pos level.Position2D) bool {
	for _, actr := range actors {
		if actr.Position.Equals(pos) {
			return true
		}
	}
	return false
}

// GetActorAtPosition returns an actor at a position if it exists. Otherwise it returns nil
func GetActorAtPosition(actors []actor.Actor, pos level.Position2D) *actor.Actor {
	for _, actr := range actors {
		if actr.Position.Equals(pos) {
			return &actr
		}
	}
	return nil
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

// places the players at the top left of the level if their position is invalid. Otherwise, uses the current
// position
func placeActors(gameState *GameState, actors []actor.Actor,
	placementFunc func(GameState, level.Position2D) level.Position2D,
	placementStart level.Position2D) {
	for _, actor := range actors {
		if actor.Position.Row < 0 || actor.Position.Col < 0 {
			gameState.SpawnActor(actor, placementFunc(*gameState, placementStart))
		} else {
			gameState.SpawnActor(actor, actor.Position)
		}
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
