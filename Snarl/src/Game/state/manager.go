package state

import (
	"encoding/json"
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/net"
	"sort"
	"time"
)

const (
	SuccessMessage = "Ok"
	InvalidMessage = "Invalid"
	ExitMessage    = "Exit"
	KeyMessage     = "Key"
	EjectMessage   = "Eject"
	TimeoutMessage = "Timeout"
	timeoutLimit   = 60 * time.Second
)

// PlayerViewDistance is the number of tiles distance that can be viewed to all four sides of the player
const PlayerViewDistance = 2

// ManageGame Runs the main game loop. It gathers input from players, calls to the rule checker to ensure that the player
// moves are valid, then moves the players if they are. It also handles all of the interactions between players, adversaries,
// and the game level. At the end of the game (if all players are ejected or the manager runs out of levels) it returns a
// ranked list of players with their scores for the main function to print
func ManageGame(gameLevels []level.Level, // The level struct for the first level
	playerClients []UserClient,           // The list of clients that control the players
	registeredPlayers []actor.Actor,      // the list of actor objects generated by the client's registration function
	observers []GameObserver,             // a list of observers that get access to all full game states
	outsideAdversaries []AdversaryClient, // An optional list of AdversaryClient if we want an outside source to control adversaries
	numLevels int,                        /* the number of levels in the game */
) []net.PlayerScore {

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
	ejectCounts := make(map[string]int)

	// temporary arrays for ejected and exited players
	var ejectedPlayers []actor.Actor
	var exitedPlayers []actor.Actor

	// first level is the first level in the array
	levelNumber := 1

	// generate adversary actors
	adversaries := generateAdversaries(levelNumber, outsideAdversaries)

	state := initGameState(gameLevels[levelNumber-1], registeredPlayers, adversaries)

	for _, client := range playerClients {
		// send first game state
		client.SendPartialState(state.GeneratePartialState(state.GetActor(client.GetName()).Position, PlayerViewDistance))
		// init scoreboard
		exitCounts[client.GetName()] = 0
		keyCounts[client.GetName()] = 0
		ejectCounts[client.GetName()] = 0
	}

	// initialize observers
	for _, observer := range observers {
		go observer.Begin()
		observer.GameStateChannel <- *state
	}

	// init adversaries
	adversaryClients := generateAdversaryClients(state.Adversaries, *state.Level, outsideAdversaries)

	// local function that updates players and observers with new game states
	updatePlayersAndObservers := func(userClients []UserClient, observers []GameObserver) {
		for _, client := range userClients {
			clientPlayer := state.GetActor(client.GetName())
			if clientPlayer == nil {
				continue
			}
			client.SendPartialState(state.GeneratePartialState(clientPlayer.Position, PlayerViewDistance))
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

			clientName := client.GetName()
			playerActor := state.GetActor(clientName)
			if playerActor == nil {
				continue
			}

			// ask the user for their move
			response, attemptedMovePos, timeOut := getUserResponseWithTimeout(client, playerActor.Position)

			if timeOut {
				client.SendMessage(TimeoutMessage, attemptedMovePos)
				continue
			}

			// check that the new game state is valid (if we get past this loop, we know it's valid)
			for !IsValidMove(*state, clientName, response.Move) {
				client.SendMessage(InvalidMessage, attemptedMovePos)
				response, attemptedMovePos, timeOut = getUserResponseWithTimeout(client, playerActor.Position)
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
			message := func(msg string) string {
				return fmt.Sprintf("Player %s %s", clientName, msg)
			}
			// if there's an adversary here, kill the player
			if ActorsOccupyPosition(adversaries, newPos) {
				exitedPlayers = append(ejectedPlayers, *playerActor)
				state.RemoveActor(clientName)
				client.SendMessage(message(EjectMessage), newPos)
				// update the scoreboard
				ejectCounts[playerActor.Name] += 1
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

					// collect exits
					var exits []string
					for _, player := range exitedPlayers {
						exits = append(exits, player.Name)
					}

					// collect ejects
					var ejects []string
					for _, player := range ejectedPlayers {
						ejects = append(ejects, player.Name)
					}

					// get the new level number
					levelNumber += 1
					state.LevelNum = levelNumber
					state.Level = &gameLevels[levelNumber-1]
					combinedList := append(exitedPlayers, ejectedPlayers...)

					// generate the actors for the next level and place them
					var nextLevelPlayers []actor.Actor
					for _, player := range combinedList {
						nextLevelPlayers = append(nextLevelPlayers, player.MoveActor(level.NewPosition2D(-1, -1)))
					}
					placeActors(state, nextLevelPlayers, getTopLeftUnoccupiedWalkable, level.NewPosition2D(0, 0))
					for _, playerClient := range playerClients {
						endLevel, _ := json.Marshal(net.EndLevel{
							Type:   "end-level",
							Key:    playerClient.GetName(),
							Exits:  exits,
							Ejects: ejects,
						})
						playerClient.SendMessage(string(endLevel), newPos)
						startLevel, _ := json.Marshal(net.NewStartLevel(levelNumber, nil))
						playerClient.SendMessage(string(startLevel), newPos)
						playerClient.SendMessage("start-level", newPos)
						playerClient.SendPartialState(state.GeneratePartialState(state.GetActor(playerClient.GetName()).Position, PlayerViewDistance))
					}

					// generate the adversaries for the next level and place them. Generate clients (if necessary after that)
					adversaries = generateAdversaries(levelNumber, outsideAdversaries)
					state.Adversaries = nil
					placeActors(state, adversaries, getBottomRightUnoccupiedWalkable, state.Level.Size)
					adversaryClients = generateAdversaryClients(state.Adversaries, *state.Level, outsideAdversaries)
					adversaries = state.Adversaries // updated positions in the new list
				}
			}
			playerPosns = append(playerPosns, newPos)
		}

		// get adversary positions
		var adversaryPosns []level.Position2D
		for _, adversary := range adversaries {
			adversaryPosns = append(adversaryPosns, adversary.Position)
		}

		// Move the adversaries
		for _, adversary := range adversaryClients {
			moveResponse := adversary.CalculateMove(playerPosns, adversaryPosns)
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

	var playerScores []net.PlayerScore
	for _, player := range playerRanks {
		playerScores = append(playerScores, net.PlayerScore{
			Type:   "player-score",
			Name:   player,
			Exits:  exitCounts[player],
			Ejects: ejectCounts[player],
			Keys:   keyCounts[player],
		})
	}

	return playerScores
}

// getUserResponseWithTimeout is a wrapper to timeout users that do not respond in a timely manner
// allows the game to continue if a player is not moving.
func getUserResponseWithTimeout(client UserClient, currentPlayerPosition level.Position2D) (Response, level.Position2D, bool){
	respChan := make(chan Response, 1)
	go func() {
		respChan <- client.GetInput()
	}()
	select {
	case response := <-respChan:
		return response, response.Move.AddPosition(currentPlayerPosition), false
	case <-time.After(timeoutLimit):
		return Response{}, level.NewPosition2D(0, 0), true

	}
}

/* ----------------------- Adversary & Adversary Client Generation -------------------------------- */

// generateAdversaries returns a list of Actor positioned at -1, -1 that corresponds to the Actors on each level.
func generateAdversaries(levelNum int, outsideAdversaries []AdversaryClient) []actor.Actor {
	var adversaries []actor.Actor
	// figure out how many of each monster to generate
	zombieNum := levelNum/2+1
	ghostNum := (levelNum-1)/2

	// generate outside adversaries first
	for _, client := range outsideAdversaries {
		// if there's nothing left to generate we return right away
		if zombieNum + ghostNum == 0 {
			return adversaries
		}
		switch client.GetType() {
		case actor.ZombieType:
			if zombieNum > 0 {
				adversaries = append(adversaries, actor.NewAdversaryActor(actor.ZombieType, client.GetName(), 1))
				zombieNum -=1
			}
		case actor.GhostType:
			if ghostNum > 0 {
				adversaries = append(adversaries, actor.NewAdversaryActor(actor.GhostType, client.GetName(), 1))
				ghostNum -=1
			}
		}
	}

	// generate remaining zombies
	for i := 0; i < zombieNum; i++ {
		adversaries = append(adversaries, actor.NewAdversaryActor(actor.ZombieType, fmt.Sprintf("z%d", i), 1))
	}
	// generate remaining ghosts
	for i := 0; i < ghostNum; i++ {
		adversaries = append(adversaries, actor.NewAdversaryActor(actor.GhostType, fmt.Sprintf("g%d", i), 1))
	}
	return adversaries
}

// generateAdversaryClients returns a consolidated list of all the AdversaryClient that will be in a level.
func generateAdversaryClients(adversaries []actor.Actor, levelData level.Level, outsideClients []AdversaryClient) []AdversaryClient {
	var adversaryClients []AdversaryClient
	for _, adversary := range adversaries {
		for _, client := range outsideClients {
			// if the name matches an existing client name, append it to the array and move on
			if adversary.Name == client.GetName() {
				client.UpdateLevel(levelData)
				client.UpdatePosition(adversary.Position)
				adversaryClients = append(adversaryClients, client)
				continue
			}
		}
		switch adversary.Type {
		case actor.ZombieType:
			adversaryClients = append(adversaryClients, &ZombieClient{
				Name:         adversary.Name,
				LevelData:    levelData,
				MoveDistance: adversary.MaxMoveDistance,
				CurrentPosn:  adversary.Position,
			})
		case actor.GhostType:
			adversaryClients = append(adversaryClients, &GhostClient{
				Name:         adversary.Name,
				LevelData:    levelData,
				MoveDistance: adversary.MaxMoveDistance,
				CurrentPosn:  adversary.Position,
			})
		}
	}
	return adversaryClients
}
