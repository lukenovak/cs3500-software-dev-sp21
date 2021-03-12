package State

import (
	"encoding/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	levelJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
	"io"
)

const parseErrorMsg = "unable to parse input. Check that your input JSON is formatted correctly"

// state JSON input
type GameStateObject struct {
	Type string                       `json:"type"`
	Level levelJson.TestLevelObject   `json:"level"`
	Players []ActorPositionObject     `json:"players"`
	Adversaries []ActorPositionObject `json:"adversaries"`
	ExitLocked bool                   `json:"exit-locked"`
}

// Converts a game state to a state object for JSON output
func gameStateObjectFromGameState(gs state.GameState, testLevel levelJson.TestLevelObject) GameStateObject {
	generatePlayerObjects := func(gameActors []actor.Actor) []ActorPositionObject {
		var actorObjs []ActorPositionObject
		for _, a := range gameActors {
			actorObjs = append(actorObjs, ActorPosObjFromGameActor(a))
		}
		return actorObjs
	}

	// Check the level to determine if the exits are locked
	exitLockStatus := true
	if gs.Level.Exits[0] != nil && gs.Level.Exits[0].Type == level.UnlockedExit {
		exitLockStatus = false
	}

	return GameStateObject{
		Type: "state",
		Level: testLevel,
		Players: generatePlayerObjects(gs.Players),
		Adversaries: generatePlayerObjects(gs.Adversaries),
		ExitLocked: exitLockStatus,
	}
}

// Actor-Position JSON input
type ActorPositionObject struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Position levelJson.LevelTestPoint `json:"position"`
}

// Converts a game Actor into an ActorPositionObject for JSON
func ActorPosObjFromGameActor(gameActor actor.Actor) ActorPositionObject {
	var actorType string
	switch gameActor.Type {
	case actor.PlayerType:
		actorType = "player"
	case actor.ZombieType:
		actorType = "zombie"
	case actor.GhostType:
		actorType = "ghost"
	}

	return ActorPositionObject{
		Type:     actorType,
		Name:     gameActor.Name,
		Position: levelJson.NewTestPointFromPosition2D(gameActor.Position),
	}
}

// parses the incoming test json into its three constituent parts
func ParseStateTestJson(r io.Reader) (GameStateObject, string, levelJson.LevelTestPoint) {
	d := json.NewDecoder(r)
	var input []json.RawMessage
	err := d.Decode(&input)
	if err != nil || len(input) != 3 {
		println(len(input))
		panic(parseErrorMsg)
	}

	// parse the state
	var stateObject GameStateObject
	if err = json.Unmarshal(input[0], &stateObject); err != nil {
		panic(err)
	}

	// parse the name
	var name string
	if json.Unmarshal(input[1], &name) != nil {
		panic(parseErrorMsg)
	}

	// parse the point
	var point levelJson.LevelTestPoint
	if json.Unmarshal(input[2], &point) != nil {
		panic(parseErrorMsg)
	}

	return stateObject, name, point

}