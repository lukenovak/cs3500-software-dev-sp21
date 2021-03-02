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

type levelTestHall struct {
	From LevelTestPoint			`json:"from"`
	To LevelTestPoint			`json:"to"`
	Waypoints []LevelTestPoint	`json:"waypoints"`
}

type LevelTestPoint [2]int

type levelTestLevel struct {
	Rooms []levelTestRoom		`json:"rooms"`
	Hallways []levelTestHall	`json:"hallways"`
	Objects []levelTestObject	`json:"objects"`
}

type levelTestObject struct {
	Type string
	Position LevelTestPoint
}

type LevelTestRoomInput struct {
	Room  levelTestRoom
	Point LevelTestPoint
}

/* ---------- Parsing JSON ---------- */

func ParseLevelTestJson(r io.Reader) LevelTestRoomInput {
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

	return LevelTestRoomInput{
		Room: room,
		Point: point,
	}

}