package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"testing"
)

func TestGameState_CreateUpdatedGameState(t *testing.T) {
	// game state generation
	testGameState := generateTestGameState()
	if len(testGameState.Players) > 0 {
		t.Fatal()
	}
	testGameState.SpawnActor(actor.Actor{Type: actor.PlayerType}, level.NewPosition2D(1, 1))
	testGameState.SpawnActor(actor.Actor{Type: 1}, level.NewPosition2D(1, 2))

	// create the updated (intermediate) game state
	newPlayers := []actor.Actor{{Type: actor.PlayerType, Position: level.NewPosition2D(2, 1)}}
	newAdversaries := []actor.Actor{{Type: 1, Position: level.NewPosition2D(1, 3)}}
	newGameState := testGameState.CreateUpdatedGameState(newPlayers, newAdversaries)

	// test to see that the game state has been updated properly

	// players in right (new) positions
	if len(newGameState.Players) != 1 || !newGameState.Players[0].Position.Equals(level.NewPosition2D(2, 1)) {
		t.Fail()
	}

	// adversaries in right (new) positions
	if len(newGameState.Adversaries) != 1 || !newGameState.Adversaries[0].Position.Equals(level.NewPosition2D(1, 3)) {
		t.Fail()
	}

	// Level is the same (pointer)
	if newGameState.Level != testGameState.Level {
		t.Fail()
	}

	// Level Num has not changed
	if newGameState.LevelNum != testGameState.LevelNum {
		t.Fail()
	}
}

func TestGameState_SpawnActor(t *testing.T) {
	// game state generation
	testGameState := generateTestGameState()
	if len(testGameState.Players) > 0 {
		t.Fatal()
	}

	// test spawning a player actor
	testActor := actor.Actor{Type: actor.PlayerType}
	testGameState.SpawnActor(testActor, level.NewPosition2D(1, 1))
	if len(testGameState.Players) != 1 || !testGameState.Players[0].Position.Equals(level.NewPosition2D(1, 1)) {
		t.Fail()
	}

	// test spawning an adversary actor
	testActor = actor.Actor{Type: 1}
	testGameState.SpawnActor(testActor, level.NewPosition2D(1, 2))
	if len(testGameState.Players) != 1 || len(testGameState.Adversaries) != 1 ||
		!testGameState.Adversaries[0].Position.Equals(level.NewPosition2D(1, 2)) {
		t.Fail()
	}

}

func TestGameState_CheckVictory(t *testing.T) {
	// test a level without victory
	testGameState := generateTestGameState()
	if testGameState.CheckVictory() {
		t.Fail()
	}

	// test a level where victory has been achieved
	testGameState.UnlockExits()
	testGameState.SpawnActor(actor.Actor{Type: actor.PlayerType}, level.NewPosition2D(12, 14))

	if !testGameState.CheckVictory() {
		t.Fail()
	}
}

func TestGameState_UnlockExits(t *testing.T) {
	// test a level with locked exits
	testGameState := generateTestGameState()
	testGameState.UnlockExits()
	if testGameState.Level.GetTile(level.NewPosition2D(12, 14)).Type != level.UnlockedExit {
		t.Fail()
	}
}

func TestGameState_MoveActor(t *testing.T) {
	testGameState := generateTestGameState()
	testGameState.SpawnActor(actor.NewWalkableActor("Luke", actor.PlayerType, 2), level.NewPosition2D(1, 1))
	testGameState.MoveActor("Luke", level.NewPosition2D(1, 2))

	if testGameState.GetActor("Luke").Position != level.NewPosition2D(1, 2) {
		t.Fail()
	}
}

/* ----------------------- TEST DATA GENERATION FUNCTIONS ------------------------------- */

func generateTestGameState() *GameState {
	testLevel := generateTestLevel()
	return &GameState{
		LevelNum:    1,
		Level:       &testLevel,
		Players:     nil,
		Adversaries: nil,
	}
}

func generateTestLevel() level.Level {
	newLevel, err := level.NewEmptyLevel(32, 32)
	if err != nil {
		panic(err)
	}
	// first room from 0,0 to 5, 4
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(0,0),
	6,
	5,
	[]level.Position2D{level.NewPosition2D(1, 4), level.NewPosition2D(5, 3)})
	if err != nil {
		panic(err)
	}

	// second room from 9, 9 to 14, 16
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(9,9),
	6,
	8,
	[]level.Position2D{level.NewPosition2D(9, 13)})
	if err != nil {
		panic(err)
	}

	// Third room from 20, 21 to 28, 29
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(20,21),
	9,
	9,
	[]level.Position2D{level.NewPosition2D(20, 25)})
	if err != nil {
		panic(err)
	}

	// connecting hallways
	hallwayWaypoints := []level.Position2D{{7, 3}, {7,13}}
	err = newLevel.GenerateHallway(level.NewPosition2D(5, 3), level.NewPosition2D(9, 13), hallwayWaypoints)
	if err != nil {
		panic(err)
	}

	hallwayWaypoints = []level.Position2D{{1, 25}}
	err = newLevel.GenerateHallway(level.NewPosition2D(1, 4), level.NewPosition2D(20, 25), hallwayWaypoints)
	if err != nil {
		panic(err)
	}

	newLevel.PlaceExit(level.NewPosition2D(12, 14))

	newLevel.PlaceItem(level.NewPosition2D(25, 25), level.Item{Type: level.KeyID})

	return newLevel
}