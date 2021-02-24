package state

import (
	"fyne.io/fyne/v2"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/render"
)

func initGameState(firstLevel level.Level, players []actor.Actor) GameState {
	gs := GameState{
		Level:         &firstLevel,
		//Adversaries: GenerateAdversaries(numPlayers),
	}
	for _, player := range players {
		gs.SpawnActor(player)
	}
	return gs
}


func GameLoop(firstLevel level.Level, players []actor.Actor, gameWindow fyne.Window) {
	state := initGameState(firstLevel, players)
	gameWindow.Resize(fyne.Size{Width: 800, Height: 800})
	for !state.CheckVictory() {
		render.GuiState(state.Level, state.Players, state.Adversaries, gameWindow)
		gameWindow.ShowAndRun()
	}
}