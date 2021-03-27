package Manager

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
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

	var testUserClients []state.UserClient
	for idx, name := range names {
		testUserClients = append(testUserClients, &TestPlayer{
			Name:            name,
			MoveList:        moves[idx],
			InitialPosition: initialPosn[idx],
			VisibleLayout:   nil,
			VisibleActors:   nil,
		})
	}

	var testPlayers []actor.Actor
	// register clients
	for _, client := range testUserClients {
		player, _ := client.RegisterClient()
		testPlayers = append(testPlayers, player)
	}

	testObservers := []state.GameObserver{state.NewGameObserver(testObserverCallback)}

	go state.GameManager(gameLevel, testUserClients, testPlayers, nil, testObservers, 1)
	for numMoves := 0; numMoves < maxMoves; numMoves++ {
		// TODO: DO something
	}

	return nil
}
