package state

// main rule checking function
func IsValidMove(oldState GameState, newState GameState) bool {
	validMove := true
	for _, p := range newState.Players {
		validMove = validMove && p.CanOccupyTile(newState.Level.GetTile(p.Position))
	}

	return validMove
}