package State

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	testJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
)

func TestUpdateState(initialState GameStateObject, name string, newPos testJson.LevelTestPoint) []interface{} {
	// convert the position to our position
	newPos2D := newPos.ToPosition2D()

	// get actors arrays
	var players []actor.Actor
	var adversaries []actor.Actor

	convertToActor := func(actorObject ActorPositionObject) actor.Actor {
		var actorType int
		switch actorObject.Type {
		case "player":
			actorType = actor.PlayerType
		case "ghost":
			actorType = actor.GhostType
		case "zombie":
			actorType = actor.ZombieType
		}
		newActor := actor.NewWalkableActor(actorObject.Name, actorType, 2)
		newActor = newActor.MoveActor(actorObject.Position.ToPosition2D())
		return newActor
	}

	for _, player := range initialState.Players {
		players = append(players, convertToActor(player))
	}

	for _, adversary := range initialState.Adversaries {
		adversaries = append(adversaries, convertToActor(adversary))
	}

	// convert the level to our level representation
	gameLevel := initialState.Level.ToGameLevel()

	// convert game state json to actual game state
	initialGameState := state.GameState{
		Players:       players,
		Adversaries:   adversaries,
		Level:         &gameLevel,
		LevelNum:      1,
		PlayerClients: nil,
	}

	if !initialState.ExitLocked {
		initialGameState.Level.UnlockExits()
	}

	// handle player not part of input
	namedPlayer := initialGameState.GetActor(name)
	if namedPlayer == nil {
		return generateInvalidPlayerMessage(name)
	}

	// handling invalid destination tiles
	if !namedPlayer.CanOccupyTile(gameLevel.GetTile(newPos2D)) ||
		(state.ActorsOccupyPosition(players, newPos2D) && state.GetActorAtPosition(players, newPos2D).Name != name) {
		return generateInvalidDestinationMessage(newPos2D)
	}

	// move the requested player
	initialGameState.MoveActorAbsolute(name, newPos2D)
	updatedActor := initialGameState.GetActor(name)

	// call our rule checking sub-functions

	// successful exit
	if gameLevel.GetTile(updatedActor.Position).Type == level.UnlockedExit {
		initialGameState.RemoveActor(name)
		return generatePlayerLeaveMessage(name, initialGameState, initialState.Level, " exited.")
	} else if state.GetActorAtPosition(adversaries, newPos2D) != nil { // successful ejection
		initialGameState.RemoveActor(name)
		return generatePlayerLeaveMessage(name, initialGameState, initialState.Level, " was ejected.")
	} else { // successful move
		// check for key and remove it from test level
		if item := initialGameState.Level.GetTile(newPos2D).Item; item != nil && item.Type == level.KeyID {
			initialGameState.Level.ClearItem(newPos2D)
			initialGameState.Level.UnlockExits()
			initialState.Level.UnlockExits()
		}
		return generateGoodMoveMessage(initialGameState, initialState.Level)
	}
}

/* ------------------------- Message Generation Functions ----------------------------- */

func generateInvalidDestinationMessage(pos level.Position2D) []interface{} {
	var msgArray []interface{}
	msgArray = append(msgArray, "Failure", "The destination position ", testJson.NewTestPointFromPosition2D(pos), " is invalid.")
	return msgArray
}

func generateInvalidPlayerMessage(name string) []interface{} {
	var msgArray []interface{}
	msgArray = append(msgArray, "Failure", "Player ", name, " is not a part of the game.")
	return msgArray
}

func generatePlayerLeaveMessage(name string, gameState state.GameState, testLevel testJson.TestLevelObject, messageEnd string) []interface{} {
	outputState := gameStateObjectFromGameState(gameState, testLevel)
	var msgArray []interface{}
	msgArray = append(msgArray, "Success", "Player ", name, messageEnd, outputState)
	return msgArray
}

func generateGoodMoveMessage(gameState state.GameState, testLevel testJson.TestLevelObject) []interface{} {
	outputState := gameStateObjectFromGameState(gameState, testLevel)
	var msgArray []interface{}
	msgArray = append(msgArray, "Success", outputState)
	return msgArray
}
