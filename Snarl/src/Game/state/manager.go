package state

import (
	"fyne.io/fyne/v2"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/internal/render"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

// runs the main game loop
func GameManager(firstLevel level.Level, players []actor.Actor, adversaries []actor.Actor, gameWindow fyne.Window) {
	if len(players) < 1 || len(players) > 4 { // we cannot start the game without the right number of players
		return
	}
	state := initGameState(firstLevel, players, adversaries)
	gameWindow.Resize(fyne.Size{Width: 800, Height: 800})

	// main game loop
	for !state.CheckVictory() {
		// check for input here

		// create an intermediate game state from the resulting input

		// check that the new game state is valid

		// TODO: Rule Checker
		//if IsValidMove(*state, *newGameState)

		// if it is, change the game state (we do not want to do this for now)

		// render the new game state
		render.GuiState(state.Level, state.Players, state.Adversaries, gameWindow)
		gameWindow.ShowAndRun()
	}
}

