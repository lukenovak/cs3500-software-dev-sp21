package tester

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
)

// Given a point, gives information about that particular tile
func TestLevelTileData(testLevel json.LevelTestLevelInput) string {
	newLevel, err := level.NewEmptyLevel(0, 0)
	if err != nil {
		panic("unable to generate new empty level")
	}

	// generate rooms
	for _, room := range testLevel.Level.Rooms {
		newOrigin := room.Origin.To2DPosition()
		newLevel.GenerateRectangularRoomWithLayout(newOrigin, room.Bounds[0], room.Bounds[1], room.Layout)
	}

	// generate hallways
	for _, hallway := range testLevel.Level.Hallways {
		newFrom := hallway.From.To2DPosition()
		newTo := hallway.To.To2DPosition()
		var newWaypoints []level.Position2D
		for _, point := range hallway.Waypoints {
			newWaypoints = append(newWaypoints, point.To2DPosition())
		}
		newLevel.GenerateHallway(newFrom, newTo, newWaypoints)
	}

	return newLevel.GetTileData(testLevel.Point.To2DPosition())
}
