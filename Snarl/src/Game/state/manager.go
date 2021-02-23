package state

import (
	"fyne.io/fyne/v2"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/render"
)

const defaultLevelSize = 32

func initGameState(numPlayers int) GameState {
	firstLevel, err := level.GenerateNewLevel(level.NewPosition2D(defaultLevelSize, defaultLevelSize))
	playerOne := actor.Actor{
		Type:     actor.PlayerType,
		Id:       1,
		Position: level.Position2D{2, 2},
		Input:    nil,
		Output:   nil,
	}
	if err != nil {
		panic(err)
	}
	gs := GameState{
		Level:         firstLevel,
		//Players:     NewPlayerList(numPlayers, ),
		//Adversaries: GenerateAdversaries(numPlayers),
	}
	gs.SpawnActor(playerOne)
	//players := actor.NewPlayerList(numPlayers, firstLevel)
	//adversaries := actor.GenerateAdversaries(numPlayers, firstLevel)
	return gs
}


func GameLoop(numPlayers int, gameWindow fyne.Window) {
	state := initGameState(numPlayers)
	gameWindow.Resize(fyne.Size{800, 800})
	for !state.CheckVictory() {
		render.GuiState(state.Level, state.Players, gameWindow)
		gameWindow.ShowAndRun()
	}
}