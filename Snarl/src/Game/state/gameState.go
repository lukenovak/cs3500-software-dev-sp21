package state

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

type GameState struct {
	Level *level.Level
	Players []*actor.Actor
	Adversaries []*actor.Actor
}

// Generates a new level
func (gs GameState) GenerateLevel(size level.Position2D) error {
	newLevel, err := level.NewEmptyLevel(size.X, size.Y)
	if err != nil {
		return err
	}
	gs.Level = &newLevel
	return nil
}

func (gs GameState) SpawnActor(position level.Position2D, actor actor.Actor) error {
	return fmt.Errorf("Not yet implemented")
}

func (gs GameState) CheckVictory() bool {
	for _, player := range gs.Players {
		if tile := gs.Level.GetTile(player.Position); tile != nil && tile.Type == level.UnlockedExit {
			return true
		}
	}
	return false
}
