package Manager

import (
	"encoding/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	testJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
	"io"
)

type ActorMove struct {
	Type string                   `json:"type"`
	// VERY IMPORTANT: THIS IS AN ABSOLUTE POINT, NOT A RELATIVE POINT
	To   *testJson.LevelTestPoint `json:"to"` // needs to be a pointer to support nil points
}

func (move ActorMove) toResponse(playerName string) state.Response {
	if move.To == nil {
		var nilMove testJson.LevelTestPoint
		nilMove = [2]int{0, 0}
		move.To = &nilMove
	}
	return state.Response{
		PlayerName: playerName,
		Move:       move.To.ToPosition2D(),
		Actions:    nil,
	}
}

func ParseManagerInput(reader io.Reader) ([]string, level.Level, int, []level.Position2D, [][]ActorMove) {
	d := json.NewDecoder(reader)
	var inputContents []json.RawMessage
	err := d.Decode(&inputContents)
	if err != nil || len(inputContents) != 5 {
		panic("invalid input")
	}

	var nameList []string
	var gameLevel level.Level
	var testLevel testJson.TestLevelObject
	var nat int
	var posList []level.Position2D
	var actorMoveListList [][]ActorMove
	var actorMoveListListRaw []json.RawMessage

	err = json.Unmarshal(inputContents[0], &nameList)

	err = json.Unmarshal(inputContents[1], &testLevel)
	gameLevel = testLevel.ToGameLevel()

	err = json.Unmarshal(inputContents[2], &nat)

	var pointList []testJson.LevelTestPoint
	err = json.Unmarshal(inputContents[3], &pointList)
	for _, point := range pointList {
		posList = append(posList, point.ToPosition2D())
	}

	err = json.Unmarshal(inputContents[4], &actorMoveListListRaw)

	for _, rawMoveList := range actorMoveListListRaw {
		var moveList []ActorMove
		err = json.Unmarshal(rawMoveList, &moveList)
		actorMoveListList = append(actorMoveListList, moveList)
	}

	return nameList, gameLevel, nat, posList, actorMoveListList

}
