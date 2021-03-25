package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

const defaultPlayerViewDistance = 2

// runs the main game loop
func GameManager(firstLevel level.Level,
	playerClients []UserClient,
	adversaries []actor.Actor,
	numLevels int) {
	if len(playerClients) < 1 || len(playerClients) > 4 { // we cannot start the game without the right number of players
		return
	}

	var players []actor.Actor
	for _, client := range playerClients {
		// adding new actor to collection
		newPlayer := actor.NewWalkableActor(client.GetName(), actor.PlayerType, 2)
		players = append(players, newPlayer)

	}

	state := initGameState(firstLevel, players, adversaries)

	// initialize players from UserClients

	// main game loop
	for !state.CheckVictory() {
		// handle player input
		for _, client := range playerClients {

			// check for input here
			response := client.GetInput()
			clientName := client.GetName()

			// check that the new game state is valid
			if IsValidMove(*state, clientName, response.Move) {
				// move the player
				state.MoveActorRelative(client.GetName(), level.NewPosition2D(response.Move.Row, response.Move.Col))

				// handle interactions
				newPos := state.GetActor(clientName).Position
				// if there's an adversary here, kill the player
				if ActorsOccupyPosition(adversaries, newPos) {
					state.RemoveActor(clientName)
				}

				playerTile := state.Level.GetTile(newPos)
				// if there's a key there, remove the key and unlock the doors {
				if playerTile != nil && playerTile.Item.Type == level.KeyID {
					state.Level.UnlockExits()
					state.Level.ClearItem(newPos)
				}

				// if the player's new pos is an unlocked door, remove the player from the gamestate
				if playerTile != nil && playerTile.Type == level.UnlockedExit {
					// TODO: Add this to a temporary array somewhere. Right now it isn't an issue because there's only 1 level
					state.RemoveActor(clientName)
				}

			}

			// update all clients
			for _, updateClient := range playerClients {
				clientPosition := state.GetActor(client.GetName()).Position
				updateClient.SendPartialState(state.GeneratePartialState(clientPosition, defaultPlayerViewDistance))
			}

			// check if this is the end of the level
			if IsLevelEnd(*state) {
				if IsGameEnd(*state, numLevels) {
					break
				} else {
					// TODO: FUTURE MILESTONE PROVIDES MULTI_LEVEL_SUPPORT
				}
			}
		}

		// TODO: Move the adversaries
		// for _, adversary := range adversaries { ... }
	}
}
