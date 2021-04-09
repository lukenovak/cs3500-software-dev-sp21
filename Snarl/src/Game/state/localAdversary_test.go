package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"testing"
)

func TestZombieClient_CalculateMove(t *testing.T) {
	testLevel := generateTestLevel()
	testZombie := ZombieClient{
		Name:         "z1",
		LevelData:    testLevel,
		MoveDistance: 1,
		CurrentPosn:  level.NewPosition2D(1, 1),
	}

	zombiePosns := []level.Position2D{level.NewPosition2D(1, 1)}

	// test that the zombie moves to the player if the player is next to the zombie
	testMove := testZombie.CalculateMove([]level.Position2D{level.NewPosition2D(1, 2)}, zombiePosns)
	if !testMove.Move.Equals(level.NewPosition2D(1, 2)) {
		t.Fail()
	}

	// test that the zombie moves towards the player if the player is in the same room
	testMove = testZombie.CalculateMove([]level.Position2D{level.NewPosition2D(1, 4)}, zombiePosns)
	if !testMove.Move.Equals(level.NewPosition2D(1, 2)) {
		t.Fail()
	}
}

func TestGhostClient_CalculateMove(t *testing.T) {
	testGhost := GhostClient{
		Name:         "g1",
		LevelData:    generateTestLevel(),
		MoveDistance: 1,
		CurrentPosn:  level.Position2D{1,1},
	}

	ghostPosns := []level.Position2D{level.NewPosition2D(1, 1)}

	// test that the zombie moves to the player if the player is next to the zombie
	testMove := testGhost.CalculateMove([]level.Position2D{level.NewPosition2D(1, 2)}, ghostPosns)
	if !testMove.Move.Equals(level.NewPosition2D(1, 2)) {
		t.Fail()
	}

	// test that the zombie moves towards the player if the player is in the same room
	testMove = testGhost.CalculateMove([]level.Position2D{level.NewPosition2D(1, 4)}, ghostPosns)
	if !testMove.Move.Equals(level.NewPosition2D(1, 2)) {
		t.Fail()
	}

	// test that the ghost moves into a wall if there is a player in another room
	testMove = testGhost.CalculateMove([]level.Position2D{level.NewPosition2D(22, 22)}, ghostPosns)
	if testGhost.LevelData.GetTile(testMove.Move).Type != level.Wall {
		t.Fail()
	}
}