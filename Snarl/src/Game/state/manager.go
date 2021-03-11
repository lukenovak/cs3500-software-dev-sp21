package state

import (
	"fyne.io/fyne/v2"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/client"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/internal/render"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

const viewDistance = 2

// runs the main game loop
func GameManager(firstLevel level.Level,
				 playerClients []client.UserClient,
				 adversaries []actor.Actor,
				 gameWindow fyne.Window,
				 numLevels int) {
	if len(playerClients) < 1 || len(playerClients) > 4 { // we cannot start the game without the right number of players
		return
	}

	var players []actor.Actor
	for _, client := range playerClients {
		newPlayer := actor.NewWalkableActor(client.GetName(), actor.PlayerType, 2)
		players = append(players, newPlayer)
	}

	state := initGameState(firstLevel, players, adversaries)
	gameWindow.Resize(fyne.Size{Width: 800, Height: 800})

	// initialize players from UserClients

	// main game loop
	for !state.CheckVictory() {
		// handle player input
		for _, client := range playerClients {

			// check for input here
			response := client.GetInput()

			// check that the new game state is valid
			if IsValidMove(*state, client.GetName(), response.Move) {
				state.MoveActorRelative(client.GetName(), level.NewPosition2D(response.Move.X, response.Move.Y))
			}

			// update all clients
			for _, updateClient := range playerClients {
				clientPosition := state.GetActor(client.GetName()).Position
				updateClient.SendPartialState(state.GeneratePartialState(clientPosition, viewDistance))
			}

			// render the new game state
			// TODO: Remove this in future milestones. This information can be handled at the playerClient level
			render.GuiState(state.Level, state.Players, state.Adversaries, gameWindow)
			gameWindow.ShowAndRun()

		}
	}
}

