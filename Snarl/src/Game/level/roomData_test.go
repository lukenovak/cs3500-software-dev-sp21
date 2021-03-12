package level

import "testing"

func TestRoomData_ConnectNode (t *testing.T) {
	graphStart := generateExampleGraph()
	newNode := HallData{
		Id:          3,
		Start:       Position2D{0,0},
		End:         Position2D{0,0},
		Waypoints:   nil,
		Connections: nil,
	}

	// check that both connections were made
	graphStart.ConnectNode(&newNode)
	containsNewConnection := false
	for _, connection := range graphStart.GetConnections() {
		if connection.GetId() == newNode.Id {
			for _, newNodeConnection := range newNode.Connections {
				if newNodeConnection == graphStart {
					containsNewConnection = true
				}
			}
		}
	}

	if !containsNewConnection {
		t.Fail()
	}
}


func generateExampleGraph() RoomGraphNode {
	// rooms carry dummy data
	startRoom := RoomData{
		Id:          0,
		TopLeft:     Position2D{0, 0},
		Size:        Position2D{0, 0},
		Connections: nil,
	}

	endRoom := RoomData{
		Id:          1,
		TopLeft:     Position2D{0, 0},
		Size:        Position2D{0, 0},
		Connections: nil,
	}

	connectingHall := HallData{
		Id:          2,
		Start:       Position2D{0, 0},
		End:         Position2D{0, 0},
		Waypoints:   nil,
		Connections: nil,
	}

	connectingHall.ConnectNode(&startRoom)
	connectingHall.ConnectNode(&endRoom)

	return &startRoom

}
