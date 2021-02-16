package render

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level"
	"testing"
)

func TestRenderRoom(t *testing.T) {
	testLevel := level.NewEmptyLevel(3, 3)
	err := testLevel.GenerateRectangularRoom(level.NewPosition2D(0, 0), 3, 3, nil)
	if err != nil {
		t.Fatal("unable to generate level")
	}
	render := RenderLevel(testLevel)
	exampleRender :=
			"▓▓▓\n" +
			"▓░▓\n" +
			"▓▓▓\n"

	if render != exampleRender {
		t.Fail()
	} else {
		print(render)
	}
}

func TestRenderLevel(t *testing.T) {
	genLevel := level.NewEmptyLevel(8, 8)
	firstRoomDoor, secondRoomDoor := []level.Position2D{level.NewPosition2D(3,2)}, []level.Position2D{level.NewPosition2D(5,4)}
	genLevel.GenerateRectangularRoom(level.NewPosition2D(0,0), 4, 4, firstRoomDoor)
	genLevel.GenerateRectangularRoom(level.NewPosition2D(4, 4), 4, 4, secondRoomDoor)
	err := genLevel.GenerateHallway(firstRoomDoor[0], secondRoomDoor[0], []level.Position2D{level.NewPosition2D(5, 2)})
	if err != nil {
		t.Fatal(err)
	}

	render := RenderLevel(genLevel)
	expectedRender :=
			"▓▓▓▓    \n" +
			"▓░░▓▓▓▓ \n" +
			"▓░░D░░▓ \n" +
			"▓▓▓▓▓░▓ \n" +
			"    ▓D▓▓\n" +
			"    ▓░░▓\n" +
			"    ▓░░▓\n" +
			"    ▓▓▓▓\n"

	print(render)
	if render != expectedRender {
		t.Fail()
	}
}