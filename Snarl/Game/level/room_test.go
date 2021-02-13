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
	for i := range generatedRoom {
		for j := range generatedRoom[i] {
			areSameRoom = areSameRoom && exampleRoom[i][j].Equals(*generatedRoom[i][j])
		}
	}
	if !areSameRoom {
		t.Fail()
	}
}


func generateExampleRoom() Room {
	room := allocateRoom(3, 3)
	room[0][0] = GenerateTile(wall, 0, 0)
	room[0][1] = GenerateTile(wall, 0, 1)
	room[0][2] = GenerateTile(wall, 0, 2)
	room[1][0] = GenerateTile(door, 1, 0)
	room[1][1] = GenerateTile(walkable, 1, 1)
	room[1][2] = GenerateTile(door, 1, 2)
	room[2][0] = GenerateTile(wall, 2, 0)
	room[2][1] = GenerateTile(wall, 2, 1)
	room[2][2] = GenerateTile(wall, 2, 2)
	return room
}
