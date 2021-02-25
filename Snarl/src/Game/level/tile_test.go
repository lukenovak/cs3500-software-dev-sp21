package level

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/item"
	"testing"
)

func TestTile_Equals(t *testing.T) {
	testTile := newWalkable()
	identicalTest := newWalkable()
	differentType := newWall()
	differentItem := newWalkable()
	differentItem.Item = &item.Item{Type: 1}

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
	if !expectedTile.Equals(*GenerateTile(Walkable)) {
		t.Fail()
	}
}

// returns a new Walkable tile without calling the tile generation function
func newWalkable() Tile {
	return Tile{Type: Walkable, Item: &item.Item{0}}
}

// returns a new Wall tile without calling the tile generation function
func newWall() Tile {
	return Tile{Type: Wall, Item: &item.Item{0}}
}