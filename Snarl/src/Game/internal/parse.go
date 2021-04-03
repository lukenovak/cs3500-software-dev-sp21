package internal

import (
	"encoding/json"
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"os"
)

/* ---------- JSON structs ---------- */

// room representation to match expected json input
type levelRoom struct {
	Type   string          `json:"type"`
	Origin levelPoint      `json:"origin"`
	Bounds levelTestBounds `json:"bounds"`
	Layout roomLayout      `json:"layout"`
}

// room layout
type roomLayout [][]int

// Hallway json representation
type levelHall struct {
	From      levelPoint   `json:"from"`
	To        levelPoint   `json:"to"`
	Waypoints []levelPoint `json:"waypoints"`
}

// point json representation
type levelPoint [2]int

// converts a levelPoint to a Position2D
func (point levelPoint) toPosition2D() level.Position2D {
	return level.NewPosition2D(point[0], point[1])
}

// Bounds json representation
type levelTestBounds struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}

// Mirrors json representation of a Level
type inputLevel struct {
	Rooms    []levelRoom       `json:"rooms"`
	Hallways []levelHall       `json:"hallways"`
	Objects  []levelTestObject `json:"objects"`
}

// converts an inputLevel to a game level
func (input inputLevel) toGameLevel() level.Level {
	var newLevel, err = level.NewEmptyLevel(4, 4)
	if err != nil {
		panic(err)
	}
	// generate rooms
	for _, room := range input.Rooms {
		newOrigin := room.Origin.toPosition2D()
		err = newLevel.GenerateRectangularRoomWithLayout(newOrigin, len(room.Layout), len(room.Layout[0]), room.Layout)
		if err != nil {
			panic(err)
		}
	}

	// generate hallways
	for _, hallway := range input.Hallways {
		newFrom := hallway.From.toPosition2D()
		newTo := hallway.To.toPosition2D()
		var newWaypoints []level.Position2D
		for _, point := range hallway.Waypoints {
			newWaypoints = append(newWaypoints, point.toPosition2D())
		}
		err = newLevel.GenerateHallway(newFrom, newTo, newWaypoints)
		if err != nil {
			panic(err)
		}
	}

	hasKey := false
	// place objects
	for _, object := range input.Objects {
		switch object.Type {
		case "key":
			err = newLevel.PlaceItem(object.Position.toPosition2D(), level.NewKey())
			hasKey = true
		case "exit":
			err = newLevel.PlaceExit(object.Position.toPosition2D())
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

// json representation of an object/Item
type levelTestObject struct {
	Type     string
	Position levelPoint
}

/* ---------- Parsing JSON ---------- */

// Parses a single level
func ParseLevelFile(filename string) ([]level.Level, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	d := json.NewDecoder(file)
	var decodedLevel inputLevel
	d.Decode(&decodedLevel)
	return []level.Level{decodedLevel.toGameLevel()}, nil
}
