package level

import (
	"testing"
)

func TestGenerateRectangularRoom(t *testing.T) {
	// generate a level with a rectangular room
	genLevel, _ := NewEmptyLevel(3, 3)
	doors := []Position2D{NewPosition2D(1, 0), NewPosition2D(1, 2)}
	err := genLevel.GenerateRectangularRoom(NewPosition2D(0, 0), 3, 3, doors)
	if err != nil {
		t.Fatal("unable to generate room")
	}
	expectedLevelTiles := generateSmallTestLevelTiles()
	testAllTilesEqual(expectedLevelTiles, genLevel.Tiles, t)

	// try placing another room on top of the existing one
	err = genLevel.GenerateRectangularRoom(NewPosition2D(0, 0), 3, 3, doors)
	if err == nil {
		t.Fail()
	}

	// test an invalid level size
	genLevel, _ = NewEmptyLevel(3, 3)
	err = genLevel.GenerateRectangularRoom(NewPosition2D(0, 0), 2, 2, doors)
	if err == nil {
		t.Fail()
	}

	// test a level that expands the level
	genLevel, _ = NewEmptyLevel(1, 1)
	err = genLevel.GenerateRectangularRoom(NewPosition2D(0, 0), 3, 3, doors)
	if err != nil {
		t.Fail()
	}
	testAllTilesEqual(expectedLevelTiles, genLevel.Tiles, t)


}

func TestGenerateHallway(t *testing.T) {
	// set up the expected level
	genLevel, _ := NewEmptyLevel(8, 8)
	firstRoomDoor, secondRoomDoor := []Position2D{NewPosition2D(3,2)}, []Position2D{NewPosition2D(5,4)}
	err := genLevel.GenerateRectangularRoom(NewPosition2D(0,0), 4, 4, firstRoomDoor)
	if err != nil {
		t.Fatal(err)
	}
	err = genLevel.GenerateRectangularRoom(NewPosition2D(4, 4), 4, 4, secondRoomDoor)
	if err != nil {
		t.Fatal(err)
	}
	err = genLevel.GenerateHallway(firstRoomDoor[0], secondRoomDoor[0], []Position2D{NewPosition2D(5, 2)})
	if err != nil {
		t.Fatal(err)
	}

	// generate the level using the functions
	expectedLevelTiles := generateTestLevelWithHallwaysTiles()
	testAllTilesEqual(expectedLevelTiles, genLevel.Tiles, t)

}

func TestNewEmptyLevel(t *testing.T) {

	// test a valid level
	level, err := NewEmptyLevel(3, 3)
	if &level == nil || !level.Size.Equals(NewPosition2D(3, 3)) || err != nil {
		t.Fail()
	}

	// test a level with 0 size
	_, err = NewEmptyLevel(0, 0)
	if err == nil {
		t.Fail()
	}

	// test a level with negative size
	_, err = NewEmptyLevel(-1, -1)
	if err == nil {
		t.Fail()
	}

	// test a level with a large size
	level, err = NewEmptyLevel(2048, 2048)
	if !level.Size.Equals(NewPosition2D(2048, 2048)) || err != nil {
		t.Fail()
	}
}

func TestPlaceExit(t *testing.T) {

	level := setupSmallTestLevel(t)

	// test a valid exit
	err := level.PlaceExit(NewPosition2D(1, 1))
	if err != nil || level.GetTile(NewPosition2D(1, 1)).Type != LockedExit {
		t.Fail()
	}

	// test an invalid exit
	err = level.PlaceExit(NewPosition2D(0, 0))
	if err == nil {
		t.Fail()
	}

	// test a negative exit
	err = level.PlaceExit(NewPosition2D(-1, -1))
	if err == nil {
		t.Fail()
	}
}

func TestPlaceItem(t *testing.T) {
	level := setupSmallTestLevel(t)

	// valid item position
	err := level.PlaceItem(NewPosition2D(1, 1), Item{Type: KeyID})
	if err != nil || level.GetTile(NewPosition2D(1, 1)).Item.Type != KeyID {
		t.Fail()
	}

	// invalid item position
	err = level.PlaceItem(NewPosition2D(0, 0), Item{Type: KeyID})
	if err == nil {
		t.Fail()
	}
}

func TestClearItem(t *testing.T) {
	// setup
	level := setupSmallTestLevel(t)
	err := level.PlaceItem(NewPosition2D(1, 1), Item{Type: KeyID})
	if err != nil {
		t.Fatal(err)
	}

	// test removing the placed item
	level.ClearItem(NewPosition2D(1,1))
	if level.GetTile(NewPosition2D(1, 1)).Item != nil {
		t.Fail()
	}
}

func TestExpandLevel(t *testing.T) {
	level := setupSmallTestLevel(t)

	// test an non-expanding call
	level.expandLevel(NewPosition2D(1, 1))
	if !level.Size.Equals(NewPosition2D(3, 3)) {
		t.Fail()
	}

	// test an expanding call (checking that the 2d slice was actually reallocated/expanded)
	level.expandLevel(NewPosition2D(5, 5))
	if !level.Size.Equals(NewPosition2D(5, 5)) || len(level.Tiles) != 5 {
		t.Fail()
	}
	for _, col := range level.Tiles {
		if len(col) != 5 {
			t.Fail()
		}
	}
}

// ------------------------------- UTILITY FUNCTIONS ------------------------------- //

// checks that all tiles are the same
func testAllTilesEqual(expected [][]*Tile, actual [][]*Tile, t *testing.T) {
	for x := range expected {
		for y := range expected[x] {
			testTile, generatedTile := expected[x][y], actual[x][y]
			if generatedTile != nil { // needs to be nested to avoid nil dereference
				if testTile == nil || !(testTile.Equals(*generatedTile)) {
					t.Fail()
				}
			} else if testTile != nil {
				t.Fail()
			}
		}
	}
}

// ------------------------------- SETUP FUNCTIONS ------------------------------- //

func setupSmallTestLevel(t *testing.T) *Level { // set up a test level
	level, _ := NewEmptyLevel(3, 3)
	err := level.GenerateRectangularRoom(NewPosition2D(0, 0), 3, 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	return &level
}

// generates a 3x3 example level grid with a 3x3 room
func generateSmallTestLevelTiles() [][]*Tile {
	roomTiles := allocateLevelTiles(3, 3)
	roomTiles[0][0] = GenerateTile(Wall, 0)
	roomTiles[0][1] = GenerateTile(Wall, 0)
	roomTiles[0][2] = GenerateTile(Wall, 0)
	roomTiles[1][0] = GenerateTile(Door, 0)
	roomTiles[1][1] = GenerateTile(Walkable, 0)
	roomTiles[1][2] = GenerateTile(Door, 0)
	roomTiles[2][0] = GenerateTile(Wall, 0)
	roomTiles[2][1] = GenerateTile(Wall, 0)
	roomTiles[2][2] = GenerateTile(Wall, 0)
	return roomTiles
}

// generates an 8x8 level with too rooms and a hallway
func generateTestLevelWithHallwaysTiles() [][]*Tile {
	levelTiles := allocateLevelTiles(8, 8)

	// generate the first room
	levelTiles[0][0] = GenerateTile(Wall, 0)
	levelTiles[0][1] = GenerateTile(Wall, 0)
	levelTiles[0][2] = GenerateTile(Wall, 0)
	levelTiles[0][3] = GenerateTile(Wall, 0)
	levelTiles[1][0] = GenerateTile(Wall, 0)
	levelTiles[1][1] = GenerateTile(Walkable, 0)
	levelTiles[1][2] = GenerateTile(Walkable, 0)
	levelTiles[1][3] = GenerateTile(Wall, 0)
	levelTiles[2][0] = GenerateTile(Wall, 0)
	levelTiles[2][1] = GenerateTile(Walkable, 0)
	levelTiles[2][2] = GenerateTile(Walkable, 0)
	levelTiles[2][3] = GenerateTile(Wall, 0)
	levelTiles[3][0] = GenerateTile(Wall, 0)
	levelTiles[3][1] = GenerateTile(Wall, 0)
	levelTiles[3][2] = GenerateTile(Door, 0)
	levelTiles[3][3] = GenerateTile(Wall, 0)

	// generate the second room
	levelTiles[4][4] = GenerateTile(Wall,1)
	levelTiles[4][5] = GenerateTile(Wall,1)
	levelTiles[4][6] = GenerateTile(Wall,1)
	levelTiles[4][7] = GenerateTile(Wall,1)
	levelTiles[5][4] = GenerateTile(Door,1)
	levelTiles[5][5] = GenerateTile(Walkable,1)
	levelTiles[5][6] = GenerateTile(Walkable,1)
	levelTiles[5][7] = GenerateTile(Wall,1)
	levelTiles[6][4] = GenerateTile(Wall,1)
	levelTiles[6][5] = GenerateTile(Walkable,1)
	levelTiles[6][6] = GenerateTile(Walkable,1)
	levelTiles[6][7] = GenerateTile(Wall,1)
	levelTiles[7][4] = GenerateTile(Wall,1)
	levelTiles[7][5] = GenerateTile(Wall,1)
	levelTiles[7][6] = GenerateTile(Wall,1)
	levelTiles[7][7] = GenerateTile(Wall,1)

	// generate the hallway
	levelTiles[4][3] = GenerateTile(Wall, 2)
	levelTiles[4][2] = GenerateTile(Walkable, 2)
	levelTiles[5][2] = GenerateTile(Walkable, 2)
	levelTiles[5][3] = GenerateTile(Walkable, 2)
	levelTiles[6][3] = GenerateTile(Wall, 2)
	levelTiles[6][2] = GenerateTile(Wall, 2)
	levelTiles[6][1] = GenerateTile(Wall, 2)
	levelTiles[5][1] = GenerateTile(Wall, 2)
	levelTiles[4][1] = GenerateTile(Wall, 2)

	return levelTiles
}