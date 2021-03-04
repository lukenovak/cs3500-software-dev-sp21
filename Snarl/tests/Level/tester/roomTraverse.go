package tester

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	testJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
)

func TestRoomTraversables(testInput testJson.LevelTestRoomInput) []interface{} {
	newLevel, err := level.NewEmptyLevel(testInput.Room.Origin[1]+len(testInput.Room.Layout[0]), testInput.Room.Origin[0]+len(testInput.Room.Layout))
	if err != nil {
		panic("unable to generate new empty level")
	}

	// figure out where the doors may be located
	for r := 0; r < len(testInput.Room.Layout); r++ {
		for c := 0; c < len(testInput.Room.Layout[r]); c++ {
			newLevel.Tiles[c+testInput.Room.Origin[1]][r+testInput.Room.Origin[0]] = level.GenerateTile(testInput.Room.Layout[r][c], 0)
		}
	}

	traversablePoints := newLevel.GetWalkableTilePositions(level.NewPosition2D(testInput.Point[1], testInput.Point[0]), 1)

	if len(traversablePoints) > 0 {
		return generateTraversableSuccessMessage(testInput.Point, testInput.Room.Origin, traversablePoints)
	} else {
		return generateTraversableFailureMessage(testInput.Point, testInput.Room.Origin)
	}
}

func generateTraversableSuccessMessage(point testJson.LevelTestPoint, origin testJson.LevelTestPoint, traversablePts []level.Position2D) []interface{} {
	var traversablePtsAsSlice []testJson.LevelTestPoint
	for _, pt := range traversablePts {
		traversablePtsAsSlice = append(traversablePtsAsSlice, testJson.LevelTestPoint{0: pt.Y, 1: pt.X})
	}

	var msg []interface{}
	msg = append(msg, "Success: Traversable points from ", point, " in room at ", origin, " are ", traversablePtsAsSlice)
	return msg
}

func generateTraversableFailureMessage(point testJson.LevelTestPoint, origin testJson.LevelTestPoint) []interface{} {
	var msg []interface{}
	msg = append(msg, "Failure: Point ", point, " is not in room at ", origin)
	return msg
}
