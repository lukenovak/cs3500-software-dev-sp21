package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
)

func main() {
	a := app.New()
	w := a.NewWindow("snarl 0.0.1")
	w.Resize(fyne.Size{Height: 800, Width: 800})
	state.GameLoop(generateGameStateLevel(), generatePlayers(),  w)
}

func generateGameStateLevel() level.Level {
	newLevel, err := level.NewEmptyLevel(32, 32)
	if err != nil {
		panic(err)
	}
	// first room from 0,0 to 5, 4
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(0,0),
		6,
		5,
		[]level.Position2D{level.NewPosition2D(1, 4), level.NewPosition2D(5, 3)})
	if err != nil {
		panic(err)
	}

	// second room from 9, 9 to 14, 16
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(9,9),
		6,
		8,
		[]level.Position2D{level.NewPosition2D(9, 13)})
	if err != nil {
		panic(err)
	}

	// Third room from 20, 21 to 28, 29
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(20,21),
		9,
		9,
		[]level.Position2D{level.NewPosition2D(20, 25)})
	if err != nil {
		panic(err)
	}

	// connecting hallways
	hallwayWaypoints := []level.Position2D{{7, 3}, {7,13}}
	err = newLevel.GenerateHallway(level.NewPosition2D(5, 3), level.NewPosition2D(9, 13), hallwayWaypoints)
	if err != nil {
		panic(err)
	}

	hallwayWaypoints = []level.Position2D{{1, 25}}
	err = newLevel.GenerateHallway(level.NewPosition2D(1, 4), level.NewPosition2D(20, 25), hallwayWaypoints)
	if err != nil {
		panic(err)
	}

	return newLevel
}

func generatePlayers() []actor.Actor {
	return []actor.Actor{{Type: actor.PlayerType}}
}