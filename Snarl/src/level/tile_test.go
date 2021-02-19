package level

import "testing"

func TestTile_Equals(t *testing.T) {
	testTile := newWalkable(0, 0)
	identicalTest := newWalkable(0, 0)
	differentPosition := newWalkable(0, 1)
	differentType := newWall(0, 0)

	if !testTile.Equals(identicalTest) {
		t.Fail()
	}
	if testTile.Equals(differentPosition) {
		t.Fail()
	}
	if testTile.Equals(differentType) {
		t.Fail()
	}
}

func TestTile_IsPosition(t *testing.T) {
	testTile := newWalkable(0, 0)
	identicalTest := newWalkable(0, 0)
	differentPosition := newWalkable(0, 1)
	differentType := newWall(0, 0)

	if !testTile.IsPosition(identicalTest.Coordinates) {
		t.Fail()
	}

	if testTile.IsPosition(differentPosition.Coordinates) {
		t.Fail()
	}

	if !testTile.IsPosition(differentType.Coordinates) {
		t.Fail()
	}
}

func TestGenerateTile(t *testing.T) {
	expectedTile := newWalkable(0,0)
	if !expectedTile.Equals(*GenerateTile(Walkable, 0, 0)) {
		t.Fail()
	}
}

// returns a new Walkable tile without calling the tile generation function
func newWalkable(x int, y int) Tile {
	return Tile{Coordinates: NewPosition2D(x, y), Type: Walkable, Item: 0}
}

// returns a new Wall tile without calling the tile generation function
func newWall(x int, y int) Tile {
	return Tile{Coordinates: NewPosition2D(x, y), Type: Wall, Item: 0}
}