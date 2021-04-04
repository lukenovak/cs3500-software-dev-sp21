package Manager

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	testJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/State"
)

type TestPlayer struct {
	Name           string
	MoveList       []ActorMove
	Position       level.Position2D
	VisibleLayout  [][]int
	VisibleActors  []State.ActorPositionObject
	VisibleObjects []testJson.TestLevelObject
	moveCount      int
	MaxMoves       int
	StopSignal     chan bool
	TraceFeed      chan interface{}
}

type MoveMessage []interface{}

type PlayerUpdate struct {
	Type     string                      `json:"type"`
	Layout   [][]int                     `json:"layout"`
	Position testJson.LevelTestPoint     `json:"position"`
	Objects  []testJson.TestLevelObject  `json:"objects"`
	Actors   []State.ActorPositionObject `json:"actors"`
}

func (t *TestPlayer) RegisterClient() (actor.Actor, error) {
	playerActor := actor.NewWalkableActor(t.Name, actor.PlayerType, 2)
	playerActor = playerActor.MoveActor(t.Position)
	t.moveCount = 0
	return playerActor, nil
}

func (t *TestPlayer) SendPartialState(tiles [][]*level.Tile, actors []actor.Actor, pos level.Position2D) error {
	t.VisibleLayout = buildLayout(tiles)
	t.VisibleActors = make([]State.ActorPositionObject, 0)
	t.Position = pos
	for _, visibleActor := range actors {
		t.VisibleActors = append(t.VisibleActors, State.ActorPosObjFromGameActor(visibleActor))
	}
	newUpdate := PlayerUpdate{
		Type:     "player-update",
		Layout:   t.VisibleLayout,
		Position: testJson.NewTestPointFromPosition2D(t.Position),
		Objects:  nil,
		Actors:   t.VisibleActors,
	}
	updateMessage := []interface{}{t.Name, newUpdate}
	t.TraceFeed <- updateMessage
	return nil
}

func (t *TestPlayer) SendMessage(message string, pos level.Position2D) error {
	if message != state.InvalidMessage {
		t.moveCount += 1
	}
	var moveMessage MoveMessage
	moveMessage = append(moveMessage, t.Name)
	moveMessage = append(moveMessage, testJson.NewTestPointFromPosition2D(pos))
	moveMessage = append(moveMessage, message)
	t.TraceFeed <- moveMessage
	if t.MoveList == nil || t.moveCount == t.MaxMoves {
		t.StopSignal <- true
	}
	return nil
}

func (t *TestPlayer) GetInput() state.Response {
	if len(t.MoveList) > 0 {
		input := t.MoveList[0]
		if len(t.MoveList) == 1 {
			t.MoveList = nil
		} else {
			t.MoveList = t.MoveList[1:]
		}
		// build the relative move
		if input.To != nil {
			relativeMove := [2]int{input.To[0] - t.Position.Row, input.To[1] - t.Position.Col}
			input.To = (*testJson.LevelTestPoint)(&relativeMove)
		} else {
			nilMove := [2]int{0, 0}
			input.To = (*testJson.LevelTestPoint)(&nilMove)
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
	return t.Name
}

func buildLayout(visibleTiles [][]*level.Tile) [][]int {

	getLayoutNumber := func(tile *level.Tile) int {
		if tile == nil {
			return 0
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
