package level

import (
	"encoding/json"
	"fmt"
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
func (point levelPoint) toPosition2D() Position2D {
	return NewPosition2D(point[0], point[1])
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
func (input inputLevel) toGameLevel() Level {
	var newLevel, err = NewEmptyLevel(4, 4)
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
		var newWaypoints []Position2D
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
			err = newLevel.PlaceItem(object.Position.toPosition2D(), &Item{Type: KeyID})
			hasKey = true
		case "exit":
			err = newLevel.PlaceItem(object.Position.toPosition2D(), &Item{Type: UnlockedExit})
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

/* ---------- External Parsing API ---------- */

// ParseLevelFile parses a valid SNARL level file
func ParseLevelFile(filename string, startLevel int) ([]Level, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	d := json.NewDecoder(file)

	var parsedLevels []Level
	numLevels := 0
	d.Decode(&numLevels)
	var decodedLevel inputLevel
	for i := 0; i < numLevels; i++ {
		d.Decode(&decodedLevel)
		parsedLevels = append(parsedLevels, decodedLevel.toGameLevel())
	}

	// error check for the start level
	if startLevel > len(parsedLevels) {
		return nil, fmt.Errorf("start level greater than the number of levels")
	}

	return parsedLevels[startLevel - 1:], nil
}
