package tester

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	testJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
)

func TestLevel(testInput testJson.LevelTestInput) string {
	newLevel, err := level.NewEmptyLevel(testInput.Room.Origin[1]+len(testInput.Room.Layout[0]), testInput.Room.Origin[0]+len(testInput.Room.Layout))
	if err != nil {
		return "FAILURE"
	}

	// figure out where the doors may be located
	for r := testInput.Room.Origin[0]; r < len(testInput.Room.Layout); r++ {
		for c := testInput.Room.Origin[1]; c < len(testInput.Room.Layout[r]); c++ {
			newLevel.Tiles[c][r] = level.GenerateTile(testInput.Room.Layout[r][c])
		}
	}

	traversablePoints := newLevel.GetWalkableTilePositions(level.NewPosition2D(testInput.Point[1], testInput.Point[0]), 1)

	if len(traversablePoints) > 0 {
		return generateSuccessMessage(testInput.Point, traversablePoints)
	} else {
		return generateFailureMessage(testInput.Point)
	}
}

func generateSuccessMessage(point testJson.LevelTestPoint, traversablePts []level.Position2D) string {
	msg := fmt.Sprintf("Success! Traversable points from [%d, %d] are: ", point[0], point[1])
	for _, pt := range traversablePts {
		msg += fmt.Sprintf(" [%d, %d]", pt.Y, pt.X)
	}
	return msg
}

func generateFailureMessage(point testJson.LevelTestPoint) string {
	return "u suck lmao"
}
