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
	gs.SpawnActor(testPlayer, level.NewPosition2D(1, 1))
	if IsGameEnd(*gs) {
		t.Fail()
	}

	movedPlayer := actor.NewWalkableActor("Luke", actor.PlayerType, 2)
	movedPlayer.Position = level.NewPosition2D(1, 2)
	newState := gs.CreateUpdatedGameState([]actor.Actor{movedPlayer}, gs.Adversaries)

	if !IsValidMove(*gs, *newState) {
		t.Fail()
	}
}

func TestIsGameEnd(t *testing.T) {
	// level setup
	gs := generateTestGameState()

	// the game should not start at the end state
	if IsGameEnd(*gs) {
		t.Fail()
	}
}

func TestIsLevelEnd(t *testing.T) {
	// level setup
	gs := generateTestGameState()

	// the game should not start at the end state
	if IsLevelEnd(*gs) {
		t.Fail()
	}
}
