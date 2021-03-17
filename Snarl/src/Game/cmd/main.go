package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/client"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("snarl 0.0.1")
	w.Resize(fyne.Size{Height: 800, Width: 800})
	w.SetOnClosed(func() {os.Exit(0)})
	state.GameManager(generateGameStateLevel(), generatePlayers(), generateAdversaries(), w, 1)
}

func generateGameStateLevel() level.Level {
	newLevel, err := level.NewEmptyLevel(32, 32)
	if err != nil {
		panic(err)
	}
	// first room from 0,0 to 5, 4
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(0,0),
		5,
		6,
		[]level.Position2D{level.NewPosition2D(4, 1), level.NewPosition2D(3, 5)})
	if err != nil {
		panic(err)
	}

	// second room from 9, 9 to 14, 16
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(9,9),
		8,
		6,
		[]level.Position2D{level.NewPosition2D(13, 9)})
	if err != nil {
		panic(err)
	}

	// Third room from 20, 21 to 28, 29
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(21,20),
		9,
		9,
		[]level.Position2D{level.NewPosition2D(25, 20)})
	if err != nil {
		panic(err)
	}

	// connecting hallways
	hallwayWaypoints := []level.Position2D{{3, 7}, {13,7}}
	err = newLevel.GenerateHallway(level.NewPosition2D(3, 5), level.NewPosition2D(13, 9), hallwayWaypoints)
	if err != nil {
		panic(err)
	}

	hallwayWaypoints = []level.Position2D{{25, 1}}
	err = newLevel.GenerateHallway(level.NewPosition2D(4, 1), level.NewPosition2D(25, 20), hallwayWaypoints)
	if err != nil {
		panic(err)
	}

	newLevel.PlaceExit(level.NewPosition2D(14, 12))

	newLevel.PlaceItem(level.NewPosition2D(25, 25), level.Item{Type: level.KeyID})

	return newLevel
}

func generatePlayers() []client.UserClient {
	return []client.UserClient{client.LocalPlayer{Name: "Luke"}}
}

func generateAdversaries() []actor.Actor {
	return []actor.Actor{{Type: 1}, {Type: 2}, {Type: 1}, {Type: 2}}
}