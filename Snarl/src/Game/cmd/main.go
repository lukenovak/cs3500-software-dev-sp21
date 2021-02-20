package main

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/render"
)

func main() {
	println("Welcome to snarl ascii level renderer!")
	println("Level 1: Simple 3 x 3 room:")
	firstLevel := generateFirstLevel()
	print(render.ASCIILevel(firstLevel))

	println("Level 2: Larger room:")
	secondLevel := generateSecondLevel()
	print(render.ASCIILevel(secondLevel))

	println("Level 3: Two rooms with hallways")
	thirdLevel := generateThirdLevel()
	print(render.ASCIILevel(thirdLevel))
}

func generateFirstLevel() level.Level {
	testLevel, err := level.NewEmptyLevel(3, 3)
	if &testLevel == nil {
		panic("unable to generate empty level")
	}
	err = testLevel.GenerateRectangularRoom(level.NewPosition2D(0, 0), 3, 3, nil)
	if err != nil {
		panic("unable to generate level")
	}
	return testLevel
}

func generateSecondLevel() level.Level {
	// test a large level
	testLevel, err := level.NewEmptyLevel(9, 6)
	if &testLevel == nil {
		panic("unable to generate empty level")
	}
	err = testLevel.GenerateRectangularRoom(level.NewPosition2D(0, 0), 9, 6, nil)
	if err != nil {
		panic("unable to generate level")
	}
	return testLevel
}

func generateThirdLevel() level.Level {
	genLevel, err := level.NewEmptyLevel(8, 8)
	if &genLevel == nil {
		panic(err)
	}
	firstRoomDoor, secondRoomDoor := []level.Position2D{level.NewPosition2D(3,2)}, []level.Position2D{level.NewPosition2D(5,4)}
	err = genLevel.GenerateRectangularRoom(level.NewPosition2D(0,0), 4, 4, firstRoomDoor)
	if err != nil {
		panic(err)
	}
	err = genLevel.GenerateRectangularRoom(level.NewPosition2D(4, 4), 4, 4, secondRoomDoor)
	if err != nil {
		panic(err)
	}
	err = genLevel.GenerateHallway(firstRoomDoor[0], secondRoomDoor[0], []level.Position2D{level.NewPosition2D(5, 2)})
	if err != nil {
		panic(err)
	}
	return genLevel
}