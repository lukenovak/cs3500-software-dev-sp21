package level

import (
	"testing"
)

func TestTile_Equals(t *testing.T) {
	testTile := newWalkable()
	identicalTest := newWalkable()
	differentType := newWall()
	differentItem := newWalkable()
	differentItem.Item = &Item{Type: 1}

	if !testTile.Equals(identicalTest) {
		t.Fail()
	}
	if testTile.Equals(differentType) {
		t.Fail()
	}
	if testTile.Equals(differentItem) {
		t.Fail()
	}
}

func TestGenerateTile(t *testing.T) {
	expectedTile := newWalkable()
	if !expectedTile.Equals(*GenerateTile(Walkable, 0)) {
		t.Fail()
	}
}

// returns a new Walkable tile without calling the tile generation function
func newWalkable() Tile {
	return Tile{Type: Walkable, RoomId: 0, Item: nil}
}

// returns a new Wall tile without calling the tile generation function
func newWall() Tile {
	return Tile{Type: Wall, RoomId: 0, Item: nil}
}
