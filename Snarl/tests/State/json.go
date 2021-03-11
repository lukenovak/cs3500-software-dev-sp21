package State

import (
	"encoding/json"
	levelJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
	"io"
)

const parseErrorMsg = "unable to parse input. Check that your input JSON is formatted correctly"

type GameStateObject struct {
	Type string `json:"type"`
	Level levelJson.TestLevelObject `json:"level"`
	Players []actorPositionObject `json:"players"`
	Adversaries []actorPositionObject `json:"adversaries"`
	ExitLocked bool `json:"exit-locked"`
}

type actorPositionObject struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Position levelJson.LevelTestPoint `json:"position"`
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
	var state GameStateObject
	if err = json.Unmarshal(input[0], &state); err != nil {
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

	return state, name, point

}