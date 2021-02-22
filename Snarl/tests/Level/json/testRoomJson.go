package json

import (
	"encoding/json"
	"io"
)

const invalidMsg = "invalid input JSON"

/* ---------- JSON structs ---------- */

type levelTestRoom struct {
	Type   string          	`json:"type"`
	Origin LevelTestPoint  	`json:"origin"`
	Bounds json.RawMessage	`json:"bounds"`
	Layout levelTestLayout 	`json:"layout"`
}

type levelTestLayout [][]int

/*func (layout levelTestLayout) equals(cmpLayout levelTestLayout) bool {
	if len(layout) != len(cmpLayout) {
		return false
	} else {
		for idx, row := range layout {
			if len(row) != len(cmpLayout)
			layout.
		}
	}
} */

type LevelTestPoint [2]int

/*func (point LevelTestPoint) equals(cmpPoint LevelTestPoint) bool {
	return point[0] == cmpPoint[0] && point[1] == cmpPoint [1]
} */

type LevelTestInput struct {
	Room  levelTestRoom
	Point LevelTestPoint
}

/* ---------- Parsing JSON ---------- */

func ParseLevelTestJson(r io.Reader) LevelTestInput {
	d := json.NewDecoder(r)
	var input []json.RawMessage

	err := d.Decode(&input)
	if err != nil {
		panic(invalidMsg)
	}

	var room levelTestRoom
	err = json.Unmarshal(input[0], &room)

	var point LevelTestPoint
	err = json.Unmarshal(input[1], &point)

	return LevelTestInput{
		Room: room,
		Point: point,
	}

}