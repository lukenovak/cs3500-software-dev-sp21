package render

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level"
	"testing"
)

func TestRenderLevelSingleRoom(t *testing.T) {

	// test a normal level
	testLevel, err := level.NewEmptyLevel(3, 3)
	if &testLevel == nil {
		t.Fatal("unable to generate empty level")
	}
	err = testLevel.GenerateRectangularRoom(level.NewPosition2D(0, 0), 3, 3, nil)
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

	// test a large level
	testLevel, err = level.NewEmptyLevel(9, 6)
	if &testLevel == nil {
		t.Fatal("unable to generate empty level")
	}
	err = testLevel.GenerateRectangularRoom(level.NewPosition2D(0, 0), 9, 6, nil)
	if err != nil {
		t.Fatal("unable to generate level")
	}
	render = RenderLevel(testLevel)
	secondTestRender :=
			"▓▓▓▓▓▓▓▓▓\n" +
			"▓░░░░░░░▓\n" +
			"▓░░░░░░░▓\n" +
			"▓░░░░░░░▓\n" +
			"▓░░░░░░░▓\n" +
			"▓▓▓▓▓▓▓▓▓\n"

	if render != secondTestRender {
		t.Fail()
	} else {
		print(render)
	}
}

func TestRenderFullLevel(t *testing.T) {
	genLevel, err := level.NewEmptyLevel(8, 8)
	if &genLevel == nil {
		t.Fatal("unable to generate empty level")
	}
	firstRoomDoor, secondRoomDoor := []level.Position2D{level.NewPosition2D(3,2)}, []level.Position2D{level.NewPosition2D(5,4)}
	err = genLevel.GenerateRectangularRoom(level.NewPosition2D(0,0), 4, 4, firstRoomDoor)
	if err != nil {
		t.Fatal(err)
	}
	err = genLevel.GenerateRectangularRoom(level.NewPosition2D(4, 4), 4, 4, secondRoomDoor)
	if err != nil {
		t.Fatal(err)
	}
	err = genLevel.GenerateHallway(firstRoomDoor[0], secondRoomDoor[0], []level.Position2D{level.NewPosition2D(5, 2)})
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