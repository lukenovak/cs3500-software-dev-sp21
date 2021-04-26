package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"testing"
)

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

func TestGameState_UnlockExits(t *testing.T) {
	// test a level with locked exits
	testGameState := generateTestGameState()
	testGameState.UnlockExits()
	if testGameState.Level.GetTile(level.NewPosition2D(12, 14)).Item.Type != level.UnlockedExit {
		t.Fail()
	}
}

func TestGameState_MoveActor(t *testing.T) {
	testGameState := generateTestGameState()
	testGameState.SpawnActor(actor.NewPlayerActor("Luke", actor.PlayerType, 2), level.NewPosition2D(1, 1))
	testGameState.MoveActorRelative("Luke", level.NewPosition2D(0, 2))

	if testGameState.GetActor("Luke").Position != level.NewPosition2D(1, 3) {
		t.Fail()
	}
}

/* ----------------------- TEST DATA GENERATION FUNCTIONS ------------------------------- */

// Generates an example game state with the generated test level
func generateTestGameState() *GameState {
	testLevel := generateTestLevel()
	return &GameState{
		LevelNum:      1,
		Level:         &testLevel,
		PlayerClients: nil,
		Players:       nil,
		Adversaries:   nil,
	}
}

// generates an example level to be used for testing
func generateTestLevel() level.Level {
	newLevel, err := level.NewEmptyLevel(32, 32)
	if err != nil {
		panic(err)
	}
	// first room from 0,0 to 5, 4
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(0, 0),
		6,
		5,
		[]level.Position2D{level.NewPosition2D(1, 4), level.NewPosition2D(5, 3)})
	if err != nil {
		panic(err)
	}

	// second room from 9, 9 to 14, 16
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(9, 9),
		6,
		8,
		[]level.Position2D{level.NewPosition2D(9, 13)})
	if err != nil {
		panic(err)
	}

	// Third room from 20, 21 to 28, 29
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(20, 21),
		9,
		9,
		[]level.Position2D{level.NewPosition2D(20, 25)})
	if err != nil {
		panic(err)
	}

	// connecting hallways
	hallwayWaypoints := []level.Position2D{{7, 3}, {7, 13}}
	err = newLevel.GenerateHallway(level.NewPosition2D(5, 3), level.NewPosition2D(9, 13), hallwayWaypoints)
	if err != nil {
		panic(err)
	}

	hallwayWaypoints = []level.Position2D{{1, 25}}
	err = newLevel.GenerateHallway(level.NewPosition2D(1, 4), level.NewPosition2D(20, 25), hallwayWaypoints)
	if err != nil {
		panic(err)
	}

	exit := level.Item{Type: level.LockedExit}
	newLevel.PlaceItem(level.NewPosition2D(12, 14), &exit)

	key := level.Item{Type: level.KeyID}
	newLevel.PlaceItem(level.NewPosition2D(25, 25), &key)

	return newLevel
}
