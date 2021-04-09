package actor

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"testing"
)

const testActorName = "Ferd"

func TestNewPlayerActor(t *testing.T) {
	// test for a Player
	testActor := Actor{Name: testActorName, Type: 0, MaxMoveDistance: 2, CanOccupyTile: canOccupyWalkable}
	generatedActor := NewPlayerActor(testActorName, PlayerType, 2)

	if testActor.Name != generatedActor.Name || testActor.MaxMoveDistance != generatedActor.MaxMoveDistance || testActor.Type != generatedActor.Type {
		t.Fail()
	}

	// test the default position (-1, -1)
	if !generatedActor.Position.Equals(level.NewPosition2D(-1, -1)) {
		t.Fail()
	}
}

func TestNewAdversaryActor(t *testing.T) {
	// test for a Zombie
	testActor := Actor{Name: "z1", MaxMoveDistance: 1, Type: ZombieType}
	generatedActor := NewAdversaryActor( ZombieType, "z1", 1)
	if testActor.Name != generatedActor.Name || testActor.MaxMoveDistance != generatedActor.MaxMoveDistance || testActor.Type != generatedActor.Type {
		t.Fail()
	}


	// test for a ghost
	testActor = Actor{Name: "g1", MaxMoveDistance: 1, Type: GhostType}
	generatedActor = NewAdversaryActor(GhostType, "g1", 1)
	if testActor.Name != generatedActor.Name || testActor.MaxMoveDistance != generatedActor.MaxMoveDistance || testActor.Type != generatedActor.Type {
		t.Fail()
	}

	// test the default position (-1, -1)
	if !generatedActor.Position.Equals(level.NewPosition2D(-1, -1)) {
		t.Fail()
	}
}

func TestActor_MoveActor(t *testing.T) {
	testActor := NewPlayerActor(testActorName, PlayerType, 2)
	newActor := testActor.MoveActor(level.NewPosition2D(2, 2))

	// MoveActor creates a new actor, so we test that the pointers aren't the same.
	if &newActor == &testActor || !newActor.Position.Equals(level.NewPosition2D(2, 2)) {
		t.Fail()
	}
}
