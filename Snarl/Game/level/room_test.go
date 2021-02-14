package level

import "testing"

func TestGenerateRectangularRoom(t *testing.T) {
	exampleRoom := generateExampleRoom()
	doors := []Position2D{NewPosition2D(1, 0), NewPosition2D(1, 2)}
	generatedRoom, err := GenerateRectangularRoom(NewPosition2D(0, 0), 3, 3, doors)
	if err != nil {
		t.Fatal("unable to generate room")
	}
	areSameRoom := true
	for i := range generatedRoom.Tiles {
		for j := range generatedRoom.Tiles[i] {
			areSameRoom = areSameRoom && exampleRoom.Tiles[i][j].Equals(*generatedRoom.Tiles[i][j])
		}
	}
	if !areSameRoom {
		t.Fail()
	}
}

// generates a 3x3 example room
func generateExampleRoom() Room {
	roomTiles := allocateRoomTiles(3, 3)
	roomTiles[0][0] = GenerateTile(Wall, 0, 0)
	roomTiles[0][1] = GenerateTile(Wall, 0, 1)
	roomTiles[0][2] = GenerateTile(Wall, 0, 2)
	roomTiles[1][0] = GenerateTile(Door, 1, 0)
	roomTiles[1][1] = GenerateTile(Walkable, 1, 1)
	roomTiles[1][2] = GenerateTile(Door, 1, 2)
	roomTiles[2][0] = GenerateTile(Wall, 2, 0)
	roomTiles[2][1] = GenerateTile(Wall, 2, 1)
	roomTiles[2][2] = GenerateTile(Wall, 2, 2)
	return Room{NewPosition2D(0,0),roomTiles}
}
