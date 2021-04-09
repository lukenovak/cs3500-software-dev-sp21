package json

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParseLevelTestJson(t *testing.T) {

	rawInput := `
[ { "type" : "room",
    "origin" : [0, 0],
    "bounds" : [3, 3],
    "layout" : [ [0, 0, 0],
                 [0, 1, 0],
                 [0, 0, 0]
               ]
  },
  [1, 0]
]`
	expected := MakeExampleLevelTestInput()

	actual := ParseRoomTestJson(bytes.NewReader([]byte(rawInput)))

	if expected.Room.Type != actual.Room.Type ||
		!reflect.DeepEqual(actual.Room.Layout, expected.Room.Layout) ||
		!reflect.DeepEqual(actual.Room.Origin, expected.Room.Origin) ||
		!reflect.DeepEqual(actual.Point, expected.Point) {
		t.Fail()
	}
}

func MakeExampleLevelTestInput() LevelTestRoomInput {
	return LevelTestRoomInput{
		Room: levelTestRoom{
			Type:   "room",
			Origin: [2]int{0, 0},
			Bounds: levelTestBounds{},
			Layout: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		Point: [2]int{1, 0},
	}
}
