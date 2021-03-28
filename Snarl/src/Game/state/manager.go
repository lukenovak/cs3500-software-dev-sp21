package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"time"
)

const (
	SuccessMessage = "Success"
	InvalidMessage = "Invalid"
	ExitMessage = "Exit"
	KeyMessage = "Key"
	EjectMessage = "Eject"
	TimeoutMessage = "Timeout"
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
			clientName := client.GetName()

			getUserResponseWithTimeout := func() (Response, level.Position2D) {
				respChan := make(chan Response, 1)
				go func() {
					respChan <- client.GetInput()
				}()
				select {
				case response := <-respChan:
					return response, response.Move.AddPosition(state.GetActor(clientName).Position)
				case <-time.After(60 * time.Second):
					timeOut = true
					return Response{}, level.NewPosition2D(0, 0)
				}
			}

			// check for input here
			response, attemptedMovePos := getUserResponseWithTimeout()

			if timeOut {
				client.SendMessage(TimeoutMessage, attemptedMovePos)
				continue
			}

			// check that the new game state is valid (if we get past this loop, we know it's valid)
			for !IsValidMove(*state, clientName, response.Move) {
				client.SendMessage(InvalidMessage, attemptedMovePos)
				response, attemptedMovePos = getUserResponseWithTimeout()
			}

			if timeOut {
				client.SendMessage(TimeoutMessage, attemptedMovePos)
				continue
			}

			// move the player
			state.MoveActorRelative(client.GetName(), level.NewPosition2D(response.Move.Row, response.Move.Col))

			// handle interactions
			newPos := state.GetActor(clientName).Position
			playerTile := state.Level.GetTile(newPos)
			// if there's an adversary here, kill the player
			if ActorsOccupyPosition(adversaries, newPos) {
				state.RemoveActor(clientName)
				client.SendMessage(EjectMessage, newPos)
			} else if playerTile != nil && playerTile.Item != nil && playerTile.Item.Type == level.KeyID {
				// grab the key if we land on it
				state.Level.UnlockExits()
				state.Level.ClearItem(newPos)
				client.SendMessage(KeyMessage, newPos)
			} else if playerTile != nil && playerTile.Type == level.UnlockedExit {
				// TODO: Add this to a temporary array somewhere. Right now it isn't an issue because there's only 1 level
				state.RemoveActor(clientName)
				client.SendMessage(ExitMessage, newPos)
			} else {
				// normal movement, send a success
				client.SendMessage(SuccessMessage, newPos)
			}

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
