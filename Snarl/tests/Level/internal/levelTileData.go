package internal

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
)

type TileData struct {
	Traversable bool 					`json:"traversable"`
	Object		interface{}				`json:"object"` // we cannot have type safety with Object as it may be nil
	Type		string					`json:"type"`
	Reachable	[]json.LevelTestPoint	`json:"reachable"`
}

// Given a point, gives information about that particular tile
func TestLevelTileData(testLevel json.LevelTestLevelInput) TileData {
	newLevel := testLevel.Level.ToGameLevel()
	// get the data from the requested tile
	return getTileData(newLevel, testLevel.Point.ToPosition2D())
}

// gets the tile data in a struct corresponding to
func getTileData(tileLevel level.Level, pos level.Position2D) TileData {
	tile := tileLevel.GetTile(pos)

	if tile == nil {
		// all nil tiles will have the same data
		return TileData{
			Traversable: false,
			Object:      nil,
			Type:        "void",
			Reachable:   nil,
		}
	}

	tileData := &TileData{}

	// traversability
	switch tile.Type {
	case level.Wall:
		if tileLevel.RoomDataGraph[tile.RoomId].Type() == "hallway" {
			tileData.Type = "void"
		}
		tileData.Traversable = false
	case level.LockedExit, level.UnlockedExit:
		tileData.Object = "exit"
		tileData.Traversable = true
	case level.Walkable, level.Door:
		tileData.Traversable = true
	}

	// objects (exits are handled with tile types above)
	if tile.Item != nil {
		switch tile.Item.Type {
		case level.KeyID:
			tileData.Object = "key"
		default:
			tileData.Object = "unknown object"
		}
	}

	// type
	var roomData level.RoomGraphNode
	if tileData.Type != "void" { // handling the strange edge case on hallways
		roomData = tileLevel.RoomDataGraph[tile.RoomId]
		tileData.Type = roomData.Type()
	}


	// Reachable
	if roomData != nil {
		tileData.Reachable = getReachableRooms(roomData)
	} else {
		tileData.Reachable = make([]json.LevelTestPoint, 0)
	}

	return *tileData
}

// gets the reachable rooms from a room
func getReachableRooms(roomData level.RoomGraphNode) []json.LevelTestPoint {
	var reachables []json.LevelTestPoint

	switch roomData.Type() {
	// in the case of rooms, we traverse the hallways to their ends
	case "room":
		for _, hallway := range roomData.GetConnections() {
			for _, room := range hallway.GetConnections() {
				if room.GetId() != roomData.GetId() {
					reachables = append(reachables, [2]int{room.GetStartPoint().Y, room.GetStartPoint().X})
				}
			}

			// edge case for room self-link
			startRoom := hallway.GetConnections()[0]
			if hallway.GetConnections()[0].GetId() == hallway.GetConnections()[1].GetId() {
				reachables = append(reachables, [2]int{startRoom.GetStartPoint().Y, startRoom.GetStartPoint().X})
			}
		}
	// in hallways, we simply use the two connections
	case "hallway":
		for _, room := range roomData.GetConnections() {
			reachables = append(reachables, [2]int{room.GetStartPoint().Y, room.GetStartPoint().X})
		}

	}

	return reachables
}