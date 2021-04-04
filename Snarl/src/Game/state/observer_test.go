package state

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"testing"
)

// tests that the activation of the observer listening works
// also tests that callback behavior works properly
func TestGameObserver_Begin(t *testing.T) {
	testGameState := generateTestGameState()
	testTileType := -1
	testObserver := NewGameObserver(func(state GameState) {
		testTileType = state.Level.GetTile(level.NewPosition2D(1, 1)).Type // should be 1
	})
	go func() {
		testObserver.GameStateChannel <- *testGameState
		close(testObserver.GameStateChannel)
	}()
	testObserver.Begin()

	if testTileType != 1 {
		t.Fail()
	}
}