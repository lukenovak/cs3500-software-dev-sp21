package tester

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level"
	testJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
)

func TestLevel(testInput testJson.LevelTestInput) string {
	newLevel, err := level.NewEmptyLevel(testInput.Point[0], testInput.Point[1])
	if err != nil {
		return "FAILURE"
	}

	// figure out where the doors may be located
	for x := 0; x < testInput.Point[0]; x++ {
		for y := 0; y < testInput.Point[1]; y++ {
			newLevel.Tiles[x][y] = level.GenerateTile(testInput.Room.Layout[y][x])
		}
	}

	traversablePoints := newLevel.GetWalkableTilePositions(level.NewPosition2D(testInput.Point[0], testInput.Point[1]), 1)

	if len(traversablePoints) > 0 {
		return generateSuccessMessage(testInput.Point, traversablePoints)
	} else {
		return generateFailureMessage(testInput.Point)
	}
}

func generateSuccessMessage(point testJson.LevelTestPoint, traversablePts []level.Position2D) string {
	msg := fmt.Sprintf("Success! Traversable points from [%d, %d] are: ", point[0], point[1])
	for _, pt := range traversablePts {
		msg += fmt.Sprintf(" [%d, %d]", pt.X, pt.Y)
	}
	return msg
}

func generateFailureMessage(point testJson.LevelTestPoint) string {
	return "u suck lmao"
}