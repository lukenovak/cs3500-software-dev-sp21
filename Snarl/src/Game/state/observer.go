package state

type GameObserver struct {
	GameStateChannel chan GameState
	callback         func(gameState GameState)
}
