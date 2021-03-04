package tester

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/item"
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
	newLevel, err := level.NewEmptyLevel(4, 4)
	if err != nil {
		panic("unable to generate new empty level")
	}

	// generate rooms
	for _, room := range testLevel.Level.Rooms {
		newOrigin := room.Origin.To2DPosition()
		err = newLevel.GenerateRectangularRoomWithLayout(newOrigin, len(room.Layout[0]), len(room.Layout), room.Layout)
		if err != nil {
			panic(err)
		}
	}

	// generate hallways
	for _, hallway := range testLevel.Level.Hallways {
		newFrom := hallway.From.To2DPosition()
		newTo := hallway.To.To2DPosition()
		var newWaypoints []level.Position2D
		for _, point := range hallway.Waypoints {
			newWaypoints = append(newWaypoints, point.To2DPosition())
		}
		err = newLevel.GenerateHallway(newFrom, newTo, newWaypoints)
		if err != nil {
			panic(err)
		}
	}

	// place objects
	for _, object := range testLevel.Level.Objects {
		switch object.Type {
		case "key":
			err = newLevel.PlaceItem(object.Position.To2DPosition(), item.NewKey())
		case "exit":
			err = newLevel.PlaceExit(object.Position.To2DPosition())
		default:
			err = fmt.Errorf("unknown item type")
		}
		if err != nil {
			panic(err)
		}
	}

	// get the data from the requested tile
	return getTileData(newLevel, testLevel.Point.To2DPosition())
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
		case item.KeyID:
			tileData.Object = "key"
		default:
			tileData.Object = "unknown object"
		}
	}

	// type
	roomData := tileLevel.RoomDataGraph[tile.RoomId]
	tileData.Type = roomData.Type()

	// Reachable
	tileData.Reachable = getReachableRooms(roomData)

	return *tileData
}

// gets the reachable rooms from a room
func getReachableRooms(roomData level.RoomGraphNode) []json.LevelTestPoint {
	var reachables []json.LevelTestPoint
	reachables = append(reachables, [2]int{roomData.GetStartPoint().Y, roomData.GetStartPoint().X})

	switch roomData.Type() {
	// in the case of rooms, we traverse the hallways to their ends
	case "room":
		for _, hallway := range roomData.GetConnections() {
			for _, room := range hallway.GetConnections() {
				if room.GetId() != roomData.GetId() {
					reachables = append(reachables, [2]int{room.GetStartPoint().Y, room.GetStartPoint().X})
				}
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