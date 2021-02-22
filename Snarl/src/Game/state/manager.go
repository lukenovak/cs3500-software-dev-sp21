package state

import (
	"fyne.io/fyne/v2"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/render"
)

const defaultLevelSize = 32

func initGameState(numPlayers int) GameState {
	firstLevel, err := level.NewEmptyLevel(defaultLevelSize, defaultLevelSize)
	if err != nil {
		panic(err)
	}
	//players := actor.NewPlayerList(numPlayers, firstLevel)
	//adversaries := actor.GenerateAdversaries(numPlayers, firstLevel)
	return GameState{
		Level:         &firstLevel,
		//Players:     NewPlayerList(numPlayers, ),
		//Adversaries: GenerateAdversaries(numPlayers),
	}
}

func renderState(state GameState, window fyne.Window) {
	window.SetContent(render.GUILevel(*state.Level))
}

func GameLoop(numPlayers int, gameWindow fyne.Window) {
	state := initGameState(numPlayers)
	for !state.CheckVictory() {
		renderState(state, gameWindow)
		gameWindow.ShowAndRun()
	}
}