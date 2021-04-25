package Manager

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	levelJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/State"
)

func Test(names []string,
	gameLevel level.Level,
	maxMoves int,
	initialPosn []level.Position2D,
	moves [][]ActorMove) []interface{} {

	var gameStateTrace []state.GameState

	testObserverCallback := func(gs state.GameState) {
		gameStateTrace = append(gameStateTrace, gs)
	}

	traceFeedChannel := make(chan interface{})
	var clientStopSignals []chan bool
	var testUserClients []state.UserClient
	for idx, name := range names {
		newStopSignal := make(chan bool, 1)
		clientStopSignals = append(clientStopSignals, newStopSignal)
		testUserClients = append(testUserClients, &TestPlayer{
			Name:       name,
			MoveList:   moves[idx],
			Position:   initialPosn[idx],
			MaxMoves:   maxMoves,
			TraceFeed:  traceFeedChannel,
			StopSignal: newStopSignal,
		})
	}

	// initialize adversaries
	var testAdversaries []actor.Actor
	if len(initialPosn) > len(names) {
		for i := len(names); i < len(initialPosn); i++ {
			adversaryName := fmt.Sprintf("ghost%d", i - len(names))
			testAdversaries = append(testAdversaries, actor.NewWalkableActor(adversaryName, actor.GhostType, 2).MoveActor(initialPosn[i]))
		}
	}

	var testPlayers []actor.Actor
	// register clients
	for _, client := range testUserClients {
		player, _ := client.RegisterClient()
		testPlayers = append(testPlayers, player)
	}

	testObservers := []state.GameObserver{state.NewGameObserver(testObserverCallback)}

	var managerTrace []interface{}

	go state.ManageGame(gameLevel, testUserClients, testPlayers, testAdversaries, testObservers, 1)

	// wait for a signal to stop from the last player
	for {
		managerTrace = append(managerTrace, <-traceFeedChannel)
		println("got here")
		select {
		case <-clientStopSignals[len(clientStopSignals)-1]:
			// do nothing
		default:
			continue
		}
		break
	}

	testLevel := levelJson.LevelToTestLevel(*gameStateTrace[len(gameStateTrace) - 1].Level)
	stateOutput := State.GameStateObjectFromGameState(gameStateTrace[len(gameStateTrace) - 1], testLevel)

	return []interface{}{stateOutput, managerTrace}
}
