package level

import "testing"

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