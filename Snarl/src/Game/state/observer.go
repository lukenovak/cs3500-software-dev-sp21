package state

type GameObserver struct {
	GameStateChannel chan GameState
	callback         func(gameState GameState)
}

// Used to create a new Game Observer with a buffer size of 1 state
func NewGameObserver(callback func(GameState)) GameObserver {
	return GameObserver {
		GameStateChannel: make(chan GameState, 1),
		callback: callback,
	} // buffer size of 1
}

// Begins listening for GameStates on the channel. When received, executes the given callback
func (observer GameObserver) Begin() {
	for state := range observer.GameStateChannel {
		observer.callback(state)
	}
}