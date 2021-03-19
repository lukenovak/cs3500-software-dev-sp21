package actor

import (
	"testing"
)

const testActorName = "Ferd"

func TestNewWalkableActor(t *testing.T) {
	// test for a Player
	testActor := Actor{Name: testActorName, Type: 0, MaxMoveDistance: 2, CanOccupyTile: canOccupyWalkable}
	generatedActor := NewWalkableActor(testActorName, PlayerType, 2)

	if testActor.Name != generatedActor.Name || testActor.MaxMoveDistance != generatedActor.MaxMoveDistance || testActor.Type != generatedActor.Type {
		t.Fail()
	}

	// test for a non-Player
	testActor = Actor{Name: "zombie1", MaxMoveDistance: 2, Type: ZombieType, CanOccupyTile: canOccupyWalkable}
	generatedActor = NewWalkableActor("zombie1", ZombieType, 2)
	if testActor.Name != generatedActor.Name || testActor.MaxMoveDistance != generatedActor.MaxMoveDistance || testActor.Type != generatedActor.Type {
		t.Fail()
	}

}
