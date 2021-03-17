package json

import (
	"encoding/json"
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"io"
)

const invalidMsg = "invalid input JSON"

/* ---------- JSON structs ---------- */

type levelTestRoom struct {
	Type   string          `json:"type"`
	Origin LevelTestPoint  `json:"origin"`
	Bounds json.RawMessage `json:"bounds"`
	Layout levelTestLayout `json:"layout"`
}

type levelTestLayout [][]int

type levelTestHall struct {
	From      LevelTestPoint   `json:"from"`
	To        LevelTestPoint   `json:"to"`
	Waypoints []LevelTestPoint `json:"waypoints"`
}

type LevelTestPoint [2]int

type TestLevelObject struct {
	Rooms    []levelTestRoom   `json:"rooms"`
	Hallways []levelTestHall   `json:"hallways"`
	Objects  []levelTestObject `json:"objects"`
}

type levelTestObject struct {
	Type     string
	Position LevelTestPoint
}

type LevelTestRoomInput struct {
	Room  levelTestRoom
	Point LevelTestPoint
}

type LevelTestLevelInput struct {
	Level TestLevelObject
	Point LevelTestPoint
}

/* ---------- Parsing JSON ---------- */

func ParseRoomTestJson(r io.Reader) LevelTestRoomInput {
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
		Room:  room,
		Point: point,
	}

}

// Converts the raw json from a level tile data test into go structs
func ParseLevelTileDataTestJson(r io.Reader) LevelTestLevelInput {
	d := json.NewDecoder(r)
	var topLevelInputData []json.RawMessage

	err := d.Decode(&topLevelInputData)
	if err != nil {
		panic(invalidMsg)
	}

	var testLevel TestLevelObject
	err = json.Unmarshal(topLevelInputData[0], &testLevel)

	var point LevelTestPoint
	err = json.Unmarshal(topLevelInputData[1], &point)

	if err != nil {
		panic(invalidMsg)
	}

	return LevelTestLevelInput{
		Level: testLevel,
		Point: point,
	}
}

/* ---------- Utility ---------- */

// Converts a LevelTestPoint to a Position2D
func (point *LevelTestPoint) ToPosition2D() level.Position2D {
	return level.NewPosition2D(point[0], point[1])
}

func NewTestPointFromPosition2D(d level.Position2D) LevelTestPoint {
	return [2]int{d.Row, d.Col}
}

func (testLevel TestLevelObject) ToGameLevel() level.Level {
	var newLevel, err = level.NewEmptyLevel(4, 4)
	if err != nil {
		panic(err)
	}
	// generate rooms
	for _, room := range testLevel.Rooms {
		newOrigin := room.Origin.ToPosition2D()
		err = newLevel.GenerateRectangularRoomWithLayout(newOrigin, len(room.Layout), len(room.Layout[0]), room.Layout)
		if err != nil {
			panic(err)
		}
	}

	// generate hallways
	for _, hallway := range testLevel.Hallways {
		newFrom := hallway.From.ToPosition2D()
		newTo := hallway.To.ToPosition2D()
		var newWaypoints []level.Position2D
		for _, point := range hallway.Waypoints {
			newWaypoints = append(newWaypoints, point.ToPosition2D())
		}
		err = newLevel.GenerateHallway(newFrom, newTo, newWaypoints)
		if err != nil {
			panic(err)
		}
	}

	hasKey := false
	// place objects
	for _, object := range testLevel.Objects {
		switch object.Type {
		case "key":
			err = newLevel.PlaceItem(object.Position.ToPosition2D(), level.NewKey())
			hasKey = true
		case "exit":
			err = newLevel.PlaceExit(object.Position.ToPosition2D())
		default:
			err = fmt.Errorf("unknown item type")
		}
		if err != nil {
			panic(err)
		}
	}

	// if there's no key we assume the exit is unlocked
	if !hasKey {
		newLevel.UnlockExits()
	}

	return newLevel
}

func (testLevel *TestLevelObject) UnlockExits() {
	for idx, item := range testLevel.Objects {
		if item.Type == "key" {
			testLevel.Objects = append(testLevel.Objects[:idx], testLevel.Objects[idx+1:]...)
		}
	}
}
