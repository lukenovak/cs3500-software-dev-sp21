package state

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"sort"
	"time"
)

const (
	SuccessMessage = "Success"
	InvalidMessage = "Invalid"
	ExitMessage    = "Exit"
	KeyMessage     = "Key"
	EjectMessage   = "Eject"
	TimeoutMessage = "Timeout"
	timeoutLimit   = 60 * time.Second
)

const defaultPlayerViewDistance = 2

// GameManager Runs the main game loop. Returns a ranked list of players
func GameManager(gameLevels []level.Level, // The level struct for the first level
	playerClients []UserClient, // The list of clients that control the players
	registeredPlayers []actor.Actor, // the list of actor objects generated by the client's regestration function
	adversaries []actor.Actor, // the list of adversary actor objects
	observers []GameObserver, // a list of observers that get access to all full game states
	numLevels int, /* the number of levels in the game */
) []string {

	// if there are no levels, we panic
	if len(gameLevels) == 0 {
		panic("no levels")
	}

	// if we have the wrong number of players we panic
	if len(playerClients) < 1 || len(playerClients) > 4 || len(registeredPlayers) != len(playerClients) {
		panic("invalid number of players")
	}

	// allocate maps for the leaderboard data
	exitCounts := make(map[string]int)
	keyCounts := make(map[string]int)

	// temporary arrays for ejected and exited players
	var ejectedPlayers []actor.Actor
	var exitedPlayers []actor.Actor

	// first level is the first level in the array
	levelNumber := 1

	state := initGameState(gameLevels[levelNumber-1], registeredPlayers, adversaries)

	for _, client := range playerClients {
		// send first game state
		client.SendPartialState(state.GeneratePartialState(state.GetActor(client.GetName()).Position, defaultPlayerViewDistance))
		// init scoreboard
		exitCounts[client.GetName()] = 0
		keyCounts[client.GetName()] = 0
	}

	// initialize observers
	for _, observer := range observers {
		go observer.Begin()
		observer.GameStateChannel <- *state
	}

	// init adversaries
	var adversaryClients []AdversaryClient
	for _, adversary := range state.Adversaries {
		if adversary.Type == actor.ZombieType {
			adversaryClients = append(adversaryClients, &ZombieClient{
				Name:         adversary.Name,
				LevelData:    *state.Level,
				MoveDistance: adversary.MaxMoveDistance,
				CurrentPosn:  adversary.Position,
			})
		}
	}

	// local function that updates players and observers with new game states
	updatePlayersAndObservers := func(userClients []UserClient, observers []GameObserver) {
		for _, client := range userClients {
			clientPlayer := state.GetActor(client.GetName())
			if clientPlayer == nil {
				continue
			}
			client.SendPartialState(state.GeneratePartialState(clientPlayer.Position, defaultPlayerViewDistance))
		}
		for _, observer := range observers {
			observer.GameStateChannel <- *state
		}
	}

	// give windows time to initialize
	println("initializing...")
	time.Sleep(2 * time.Second)

	// main game loop
	for {
		var playerPosns []level.Position2D
		// handle player input
		for _, client := range playerClients {

			timeOut := false
			clientName := client.GetName()
			playerActor := state.GetActor(clientName)
			if playerActor == nil { // TODO: more robust error handling here. Could this produce an endless loop?
				continue
			}

			// local function wrapper to timeout users that do not respond in a timely manner
			// allows the game to continue if a player is not moving.
			getUserResponseWithTimeout := func() (Response, level.Position2D) {
				respChan := make(chan Response, 1)
				go func() {
					respChan <- client.GetInput()
				}()
				select {
				case response := <-respChan:
					return response, response.Move.AddPosition(playerActor.Position)
				case <-time.After(timeoutLimit):
					timeOut = true
					return Response{}, level.NewPosition2D(0, 0)
				}
			}

			response, attemptedMovePos := getUserResponseWithTimeout()

			if timeOut {
				client.SendMessage(TimeoutMessage, attemptedMovePos)
				timeOut = false
				continue
			}

			// check that the new game state is valid (if we get past this loop, we know it's valid)
			for !IsValidMove(*state, clientName, response.Move) {
				client.SendMessage(InvalidMessage, attemptedMovePos)
				response, attemptedMovePos = getUserResponseWithTimeout()
			}

			if timeOut {
				client.SendMessage(TimeoutMessage, attemptedMovePos)
				timeOut = false
				continue
			}

			// move the player
			state.MoveActorRelative(client.GetName(), level.NewPosition2D(response.Move.Row, response.Move.Col))

			// handle interactions
			newPos := state.GetActor(clientName).Position
			playerTile := state.Level.GetTile(newPos)
			message := func(msg string) string {
				return fmt.Sprintf("Player %s %s", clientName, msg)
			}
			// if there's an adversary here, kill the player
			if ActorsOccupyPosition(adversaries, newPos) {
				exitedPlayers = append(ejectedPlayers, *playerActor)
				state.RemoveActor(clientName)
				client.SendMessage(message(EjectMessage), newPos)
			} else if playerTile != nil && playerTile.Item != nil && playerTile.Item.Type == level.KeyID {
				// grab the key if we land on it
				state.Level.UnlockExits()
				state.Level.ClearItem(newPos)
				client.SendMessage(message(KeyMessage), newPos)
				// update the scoreboard
				keyCounts[playerActor.Name] += 1
			} else if playerTile != nil && playerTile.Item != nil && playerTile.Item.Type == level.UnlockedExit {
				exitedPlayers = append(exitedPlayers, playerActor.MoveActor(level.NewPosition2D(-1, -1)))
				state.RemoveActor(clientName)
				client.SendMessage(message(ExitMessage), newPos)
				// update the scoreboard
				exitCounts[playerActor.Name] += 1
			} else {
				// normal movement, send a success
				client.SendMessage(message(SuccessMessage), newPos)
			}

			// update all clients
			updatePlayersAndObservers(playerClients, observers)

			// check if this is the end of the level
			if IsLevelEnd(*state) {
				// check if this is the end of the game. If it is, break the loop
				if IsGameEnd(*state, numLevels) {
					break
				} else {
					// If the level is over but the game is not, start the next level
					levelNumber += 1
					state.LevelNum = levelNumber
					state.Level = &gameLevels[levelNumber-1]
					combinedList := append(exitedPlayers, ejectedPlayers...)
					var nextLevelPlayers []actor.Actor
					for _, player := range combinedList {
						nextLevelPlayers = append(nextLevelPlayers, player.MoveActor(level.NewPosition2D(-1, -1)))
					}
					placeActors(state, nextLevelPlayers, getTopLeftUnoccupiedWalkable, level.NewPosition2D(0, 0))
					for _, playerClient := range playerClients {
						playerClient.SendPartialState(state.GeneratePartialState(state.GetActor(playerClient.GetName()).Position, defaultPlayerViewDistance))
					}
					// TODO: Adversaries
				}
			}
			playerPosns = append(playerPosns, newPos)
		}

		// Move the adversaries
		for _, adversary := range adversaryClients {
			moveResponse := adversary.CalculateMove(playerPosns, []level.Position2D{})
			state.MoveActorAbsolute(adversary.GetName(), moveResponse.Move)
			// check for ejections
			if ActorsOccupyPosition(state.Players, moveResponse.Move) {
				killedPlayer := GetActorAtPosition(state.Players, moveResponse.Move)
				state.RemoveActor(killedPlayer.Name)
				ejectedPlayers = append(ejectedPlayers, *killedPlayer)
			}
			adversary.UpdatePosition(moveResponse.Move)
			updatePlayersAndObservers(playerClients, observers)
		}

		// check to see if the last player has been killed
		if len(state.Players) == 0 {
			break
		}

		// reset player positions for pathfinding
		playerPosns = make([]level.Position2D, len(state.Players))

	}

	var playerRanks []string
	for _, player := range playerClients {
		playerRanks = append(playerRanks, player.GetName())
	}

	sort.Slice(playerRanks, func(i, j int) bool {
		if exitCounts[playerRanks[i]] > exitCounts[playerRanks[j]] {
			return true
		} else if exitCounts[playerRanks[i]] == exitCounts[playerRanks[j]] {
			return keyCounts[playerRanks[i]] > keyCounts[playerRanks[j]]
		} else {
			return false
		}
	})

	playerRanks = append(playerRanks, fmt.Sprint("Player Leaderboard:\n"))
	for _, player := range playerRanks {
		playerRanks = append(playerRanks, fmt.Sprintf("%v, %v, %v\n", player, exitCounts[player], keyCounts[player]))
	}

	return playerRanks
}
