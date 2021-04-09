package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"testing"
)

const lukeActorName = "Luke"
const willActorName = "Will"
const ghostActorName = "Casper"

func TestIsValidMove(t *testing.T) {

	// level setup
	gs := generateTestGameState()
	testPlayer := actor.NewPlayerActor(lukeActorName, actor.PlayerType, 2)
	secondTestPlayer := actor.NewPlayerActor(willActorName, actor.PlayerType, 2)
	testAdversary := actor.NewPlayerActor(ghostActorName, actor.GhostType, 1)

	gs.SpawnActor(testPlayer, level.NewPosition2D(1, 1))

	// testing a valid move
	if !IsValidMove(*gs, lukeActorName, level.NewPosition2D(2, 0)) {
		t.Fail()
	}

	// testing a valid move over a player
	gs.SpawnActor(secondTestPlayer, level.NewPosition2D(2, 1))

	if !IsValidMove(*gs, lukeActorName, level.NewPosition2D(2, 0)) {
		t.Fail()
	}

	// testing an invalid move into a wall
	if IsValidMove(*gs, lukeActorName, level.NewPosition2D(-1, 0)) {
		t.Fail()
	}

	// testing a valid move to a door
	gs.MoveActorAbsolute(lukeActorName, level.NewPosition2D(3, 1))
	if !IsValidMove(*gs, lukeActorName, level.NewPosition2D(1, 0)) {
		t.Fail()
	}

	// testing an invalid move over a wall
	gs.MoveActorAbsolute(lukeActorName, level.NewPosition2D(3, 4))
	if IsValidMove(*gs, lukeActorName, level.NewPosition2D(0, 2)) {
		t.Fail()
	}

	gs.SpawnActor(testAdversary, level.NewPosition2D(2, 2))

	// testing a valid move of an adversary
	if !IsValidMove(*gs, ghostActorName, level.NewPosition2D(-1, 0)) {
		t.Fail()
	}

	// testing an invalid move of an adversary
	if IsValidMove(*gs, ghostActorName, level.NewPosition2D(1, 1)) {
		t.Fail()
	}
}

func TestIsGameEnd(t *testing.T) {
	// level setup
	gs := generateTestGameState()

	gs.SpawnActor(actor.NewPlayerActor(lukeActorName, actor.PlayerType, 2), level.NewPosition2D(1, 1))

	// the game should not start at the end state
	if IsGameEnd(*gs, 5) {
		t.Fail()
	}
}

func TestIsLevelEnd(t *testing.T) {
	// level setup
	gs := generateTestGameState()
	testPlayer := actor.NewPlayerActor(lukeActorName, actor.PlayerType, 2)

	// the game should not end when the player is on a locked door
	gs.SpawnActor(testPlayer, level.NewPosition2D(12, 14))

	if IsLevelEnd(*gs) {
		t.Fail()
	}

	// with the door unlocked the level should be at a valid end since the player is there
	gs.UnlockExits()

	if !IsLevelEnd(*gs) {
		t.Fail()
	}

}
