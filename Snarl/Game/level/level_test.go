package level

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/item"
	"testing"
)

func TestGenerateRectangularRoom(t *testing.T) {
	genLevel, _ := NewEmptyLevel(3, 3)
	doors := []Position2D{NewPosition2D(1, 0), NewPosition2D(1, 2)}
	err := genLevel.GenerateRectangularRoom(NewPosition2D(0, 0), 3, 3, doors)
	if err != nil {
		t.Fatal("unable to generate room")
	}
	testLevelTiles := generateSmallTestLevelTiles()
	areSameRoom := true
	for i := range testLevelTiles {
		for j := range testLevelTiles[i] {
			areSameRoom = areSameRoom && testLevelTiles[i][j].Equals(*genLevel.Tiles[i][j])
		}
	}
	if !areSameRoom {
		t.Fail()
	}
}

func TestGenerateHallway(t *testing.T) {
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
	testLevelTiles := generateTestLevelWithHallwaysTiles()
	for i := range testLevelTiles {
		for j := range testLevelTiles[i] {
			testTile, generatedTile := testLevelTiles[i][j], genLevel.Tiles[i][j]
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
	if err != nil || level.getTile(NewPosition2D(1, 1)).Type != LockedExit {
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
	level.PlaceItem(NewPosition2D(1, 1), item.KeyID)
}

// ------------------------------- SETUP FUNCTIONS ------------------------------- //

func setupSmallTestLevel(t *testing.T) *Level {	// set up a test level
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
	roomTiles[0][0] = GenerateTile(Wall, 0, 0)
	roomTiles[0][1] = GenerateTile(Wall, 0, 1)
	roomTiles[0][2] = GenerateTile(Wall, 0, 2)
	roomTiles[1][0] = GenerateTile(Door, 1, 0)
	roomTiles[1][1] = GenerateTile(Walkable, 1, 1)
	roomTiles[1][2] = GenerateTile(Door, 1, 2)
	roomTiles[2][0] = GenerateTile(Wall, 2, 0)
	roomTiles[2][1] = GenerateTile(Wall, 2, 1)
	roomTiles[2][2] = GenerateTile(Wall, 2, 2)
	return roomTiles
}

// generates an 8x8 level with too rooms and a hallway
func generateTestLevelWithHallwaysTiles() [][]*Tile {
	levelTiles := allocateLevelTiles(8, 8)

	// generate the first room
	levelTiles[0][0] = GenerateTile(Wall, 0, 0)
	levelTiles[0][1] = GenerateTile(Wall, 0, 1)
	levelTiles[0][2] = GenerateTile(Wall, 0, 2)
	levelTiles[0][3] = GenerateTile(Wall, 0, 3)
	levelTiles[1][0] = GenerateTile(Wall, 1, 0)
	levelTiles[1][1] = GenerateTile(Walkable, 1, 1)
	levelTiles[1][2] = GenerateTile(Walkable, 1, 2)
	levelTiles[1][3] = GenerateTile(Wall, 1, 3)
	levelTiles[2][0] = GenerateTile(Wall, 2, 0)
	levelTiles[2][1] = GenerateTile(Walkable, 2, 1)
	levelTiles[2][2] = GenerateTile(Walkable, 2, 2)
	levelTiles[2][3] = GenerateTile(Wall, 2, 3)
	levelTiles[3][0] = GenerateTile(Wall, 3, 0)
	levelTiles[3][1] = GenerateTile(Wall, 3, 1)
	levelTiles[3][2] = GenerateTile(Door, 3, 2)
	levelTiles[3][3] = GenerateTile(Wall, 3, 3)

	// generate the second room
	levelTiles[4][4] = GenerateTile(Wall, 4, 4)
	levelTiles[4][5] = GenerateTile(Wall, 4, 5)
	levelTiles[4][6] = GenerateTile(Wall, 4, 6)
	levelTiles[4][7] = GenerateTile(Wall, 4, 7)
	levelTiles[5][4] = GenerateTile(Door, 5, 4)
	levelTiles[5][5] = GenerateTile(Walkable, 5, 5)
	levelTiles[5][6] = GenerateTile(Walkable, 5, 6)
	levelTiles[5][7] = GenerateTile(Wall, 5, 7)
	levelTiles[6][4] = GenerateTile(Wall, 6, 4)
	levelTiles[6][5] = GenerateTile(Walkable, 6, 5)
	levelTiles[6][6] = GenerateTile(Walkable, 6, 6)
	levelTiles[6][7] = GenerateTile(Wall, 6, 7)
	levelTiles[7][4] = GenerateTile(Wall, 7, 4)
	levelTiles[7][5] = GenerateTile(Wall, 7, 5)
	levelTiles[7][6] = GenerateTile(Wall, 7, 6)
	levelTiles[7][7] = GenerateTile(Wall, 7, 7)

	// generate the hallway
	levelTiles[4][3] = GenerateTile(Wall, 4, 3)
	levelTiles[4][2] = GenerateTile(Walkable, 4, 2)
	levelTiles[5][2] = GenerateTile(Walkable, 5, 2)
	levelTiles[5][3] = GenerateTile(Walkable, 5, 3)
	levelTiles[6][3] = GenerateTile(Wall, 6, 3)
	levelTiles[6][2] = GenerateTile(Wall, 6, 2)
	levelTiles[6][1] = GenerateTile(Wall, 6, 1)
	levelTiles[5][1] = GenerateTile(Wall, 5, 1)
	levelTiles[4][1] = GenerateTile(Wall, 4, 1)

	return levelTiles
}