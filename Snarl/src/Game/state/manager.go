package state

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"time"
)

const defaultPlayerViewDistance = 2

// runs the main game loop
func GameManager(firstLevel level.Level,
	playerClients []UserClient,
	registeredPlayers []actor.Actor,
	adversaries []actor.Actor,
	observers []GameObserver,
	numLevels int) {

	if len(playerClients) < 1 || len(playerClients) > 4 || len(registeredPlayers) != len(playerClients) { // we cannot start the game without the right number of players
		return
	}

	state := initGameState(firstLevel, registeredPlayers, adversaries)

	for _, client := range playerClients {
		client.SendPartialState(state.GeneratePartialState(state.GetActor(client.GetName()).Position, defaultPlayerViewDistance))

	}
	// initialize observers
	for _, observer := range observers {
		go observer.Begin()
		observer.GameStateChannel <- *state
	}

	// main game loop
	for !state.CheckVictory() {
		// handle player input
		for _, client := range playerClients {

			timeOut := false

			getUserResponseWithTimeout := func() Response {
				respChan := make(chan Response, 1)
				client.SendMessage("invalid move. Try again")
				go func() {
					respChan <- client.GetInput()
				}()
				select {
				case response := <-respChan:
					return response
				case <-time.After(60 * time.Second):
					timeOut = true
					return Response{}
				}
			}

			// check for input here
			response := getUserResponseWithTimeout()
			clientName := client.GetName()

			if timeOut {
				client.SendMessage("skipping move for inactivity")
				continue
			}

			// check that the new game state is valid (if we get past this loop, we know it's valid)
			for !IsValidMove(*state, clientName, response.Move) {
				client.SendMessage("invalid move. Try again")
				response = getUserResponseWithTimeout()
			}

			if timeOut {
				client.SendMessage("skipping move for inactivity")
				continue
			}

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
			if playerTile != nil && playerTile.Item != nil && playerTile.Item.Type == level.KeyID {
				state.Level.UnlockExits()
				state.Level.ClearItem(newPos)
			}

			// if the player's new pos is an unlocked door, remove the player from the gamestate
			if playerTile != nil && playerTile.Type == level.UnlockedExit {
				// TODO: Add this to a temporary array somewhere. Right now it isn't an issue because there's only 1 level
				state.RemoveActor(clientName)
			}

			client.SendMessage(fmt.Sprintf("%s moved to %d, %d\n", clientName, newPos.Row, newPos.Col))

			// update all clients
			for _, updateClient := range playerClients {
				clientPosition := state.GetActor(client.GetName()).Position
				updateClient.SendPartialState(state.GeneratePartialState(clientPosition, defaultPlayerViewDistance))
			}
			for _, observer := range observers {
				observer.GameStateChannel <- *state
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
