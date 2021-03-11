package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"testing"
)

func TestIsValidMove(t *testing.T) {

	// level setup
	gs := generateTestGameState()
	testPlayer := actor.NewWalkableActor("Luke", actor.PlayerType, 2)
	secondTestPlayer := actor.NewWalkableActor("Will", actor.PlayerType, 2)

	gs.SpawnActor(testPlayer, level.NewPosition2D(1, 1))

	movedPlayer := actor.NewWalkableActor("Luke", actor.PlayerType, 2)
	movedPlayer.Position = level.NewPosition2D(1, 2)
	newState := *gs.CreateUpdatedGameState([]actor.Actor{movedPlayer}, gs.Adversaries)


	// testing a valid move
	if !IsValidMove(*gs, newState) {
		t.Fail()
	}

	// testing a valid move over a player
	newState = gs.CopyGameState()
	newState.SpawnActor(secondTestPlayer, level.NewPosition2D(1, 2))
	newState.MoveActor("Luke", level.NewPosition2D(1, 3))

	if !IsValidMove(*gs, newState) {
		t.Fail()
	}

	// TODO: testing a valid move to a door

	// TODO: testing an invalid move over a wall

	// testing an invalid move into a wall
	newState = *newState.CreateUpdatedGameState([]actor.Actor{
		newState.Players[0].MoveActor(level.NewPosition2D(2, 0)),
	}, gs.Adversaries)

	if IsValidMove(*gs, newState) {
		t.Fail()
	}

	// TODO: testing a valid move of an adversary

	// TODO: testing an invalid move of an adversary

}

func TestIsGameEnd(t *testing.T) {
	// level setup
	gs := generateTestGameState()

	// the game should not start at the end state
	if IsGameEnd(*gs, 5) {
		t.Fail()
	}
}

func TestIsLevelEnd(t *testing.T) {
	// level setup
	gs := generateTestGameState()
	testPlayer := actor.NewWalkableActor("Luke", actor.PlayerType, 2)

	// the game should not start at the end state
	if IsLevelEnd(*gs) {
		t.Fail()
	}

	// the game should not end when the player is on a locked door
	gs.SpawnActor(testPlayer, level.NewPosition2D(12, 14))

	if IsLevelEnd(*gs) {
		t.Fail()
	}

	// with the door unlocked the level should end
	gs.UnlockExits()

	if !IsLevelEnd(*gs) {
		t.Fail()
	}

}
