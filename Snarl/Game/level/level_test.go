package level

import (
	"testing"
)

func TestGenerateRectangularRoom(t *testing.T) {
	genLevel := NewEmptyLevel(3, 3)
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
	genLevel := NewEmptyLevel(8, 8)
	firstRoomDoor, secondRoomDoor := []Position2D{NewPosition2D(3,2)}, []Position2D{NewPosition2D(5,4)}
	genLevel.GenerateRectangularRoom(NewPosition2D(0,0), 4, 4, firstRoomDoor)
	genLevel.GenerateRectangularRoom(NewPosition2D(4, 4), 4, 4, secondRoomDoor)
	err := genLevel.GenerateHallway(firstRoomDoor[0], secondRoomDoor[0], []Position2D{NewPosition2D(5, 2)})
	if err != nil {
		t.Fatal(err)
	}
	areSameRoom := true
	testLevelTiles := generateTestLevelWithHallwaysTiles()
	for i := range testLevelTiles {
		for j := range testLevelTiles[i] {
			if testLevelTiles[i][j] == nil {
				areSameRoom = genLevel.Tiles[i][j] == nil
			} else if genLevel.Tiles[i][j] == nil {
				t.Fail()
			} else {
				areSameRoom = areSameRoom && testLevelTiles[i][j].Equals(*genLevel.Tiles[i][j])
			}
		}
	}
	if !areSameRoom {
		t.Fail()
	}
}

// generates a 3x3 example room
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