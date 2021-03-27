package Manager

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/State"
)

type TestPlayer struct {
	Name string
	MoveList []ActorMove
	InitialPosition level.Position2D
	VisibleLayout [][]int
	VisibleActors []State.ActorPositionObject
}

func (t *TestPlayer) RegisterClient() (actor.Actor, error) {
	playerActor := actor.NewWalkableActor(t.Name, actor.PlayerType, 2)
	playerActor = playerActor.MoveActor(t.InitialPosition)
	return playerActor, nil
}

func (t *TestPlayer) SendPartialState(tiles [][]*level.Tile, actors []actor.Actor) error {
	t.VisibleLayout = buildLayout(tiles)
	t.VisibleActors = make([]State.ActorPositionObject, 0)
	for _, visibleActor := range actors {
		t.VisibleActors = append(t.VisibleActors, State.ActorPosObjFromGameActor(visibleActor))
	}
	return nil
}

func (t *TestPlayer) SendMessage(message string) error {
	return nil // we ignore all messages that are sent to produce the proper output
}

func (t *TestPlayer) GetInput() state.Response {
	if len(t.MoveList) > 0 {
		input := t.MoveList[0]
		if len(t.MoveList) == 1 {
			t.MoveList = nil
		} else {
			t.MoveList = t.MoveList[1:]
		}
		return input.toResponse(t.Name)
	}
	return state.Response{
		PlayerName: t.Name,
		Move:       level.Position2D{Row: 0, Col: 0},
		Actions:    nil,
	}
}

func (t *TestPlayer) GetName() string {
	return t.GetName()
}

func buildLayout(visibleTiles [][]*level.Tile) [][]int {

	getLayoutNumber := func(tile *level.Tile) int {
		if tile == nil {
			return 0
		} else if tile.Type == level.UnlockedExit || tile.Type == level.LockedExit {
			return 1
		} else {
			return tile.Type
		}
	}

	var layout [][]int
	for rowNum, r := range visibleTiles {
		newRow := make([]int, 0)
		layout = append(layout, newRow)
		for _, t := range r {
			layout[rowNum] = append(layout[rowNum], getLayoutNumber(t))
		}
	}
	return layout
}
