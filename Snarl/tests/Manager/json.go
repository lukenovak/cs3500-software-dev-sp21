package Manager

import (
	"encoding/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	testJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
	"io"
)

type ActorMove struct {
	Type string `json:"type"`
	To *testJson.LevelTestPoint // needs to be a pointer to support nil points
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

func ParseManagerInput(reader io.Reader) ([]string, level.Level, testJson.TestLevelObject, int, []level.Position2D, [][]ActorMove) {
	d := json.NewDecoder(reader)
	var inputContents []json.RawMessage
	err := d.Decode(&inputContents)
	if err != nil || len(inputContents) != 5 {
		panic("invalid input")
	}

	parsePos := 0
	var nameList []string
	var gameLevel level.Level
	var testLevel testJson.TestLevelObject
	var nat int
	var posList []level.Position2D
	var actorMoveListList [][]ActorMove

	for err == nil && parsePos < 5 {
		err = json.Unmarshal(inputContents[parsePos], &nameList)
		parsePos += 1

		err = json.Unmarshal(inputContents[parsePos], &testLevel)
		gameLevel = testLevel.ToGameLevel()
		parsePos += 1

		err = json.Unmarshal(inputContents[parsePos], &nat)
		parsePos += 1

		var pointList []testJson.LevelTestPoint
		err = json.Unmarshal(inputContents[parsePos], &pointList)
		for _, point := range pointList {
			posList = append(posList, point.ToPosition2D())
		}

		err = json.Unmarshal(inputContents[parsePos], &actorMoveListList)
		parsePos += 1
	}

	return nameList, gameLevel, testLevel, nat, posList, actorMoveListList

}
