package remote

import (
	"encoding/json"
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

type ServerWelcome struct {
	Type string `json:"type"`
	Info string `json:"info"`
}

func NewServerWelcomeMessage() json.RawMessage {
	welcome := &ServerWelcome{
		Type: "welcome",
		Info: fmt.Sprintf("Snarl Server version %s", ServerVersion),
	}
	message, _ := json.Marshal(welcome)
	return message
}

type StartLevel struct {
	Type    string `json:"type"`
	Level   int    `json:"level"`
	Players []string `json:"players"`
}

func NewStartLevel(levelNum int, playerList []string) *StartLevel {
	return &StartLevel{
		Type:    "start-level",
		Level:   levelNum,
		Players: playerList,
	}
}

// Point represents a row, column point. Maps nicely to level.Position2D
type Point [2]int

func (p Point) ToPos2D() level.Position2D {
	return level.NewPosition2D(p[0], p[1])
}

// PointFromPos2d creates a Point from a level.Position2D
func PointFromPos2d(d level.Position2D) Point {
	return [2]int{d.Row, d.Col}
}

// Object represents a game "item" or "object" located on a tile
type Object struct {
	Type     string `json:"type"`
	Position Point  `json:"position"`
}

type ActorPosition struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Position Point  `json:"position"`
}

// NewActorPositionFromActor creates an ActorPosition object from an actor.Actor
func NewActorPositionFromActor(a actor.Actor) *ActorPosition {
	return &ActorPosition{
		Type:     "actor-position",
		Name:     a.Name,
		Position: PointFromPos2d(a.Position),
	}
}

type PlayerUpdateMessage struct {
	Type     string          `json:"type"`
	Layout   [][]int         `json:"layout"`
	Position Point           `json:"position"`
	Objects  []Object        `json:"objects"`
	Actors   []ActorPosition `json:"actors"`
	Message  string          `json:"message"`
}

// NewPlayerUpdateMessage constructs a PlayerUpdateMessage from the necessary fields
func NewPlayerUpdateMessage(layout [][]int, position Point, objects []Object, actors []ActorPosition, message string) *PlayerUpdateMessage {
	return &PlayerUpdateMessage{
		Type:     "player-update",
		Layout:   layout,
		Position: position,
		Objects:  objects,
		Actors:   actors,
		Message:  message,
	}
}

type PlayerMove struct {
	Type string `json:"type"`
	To   Point  `json:"to"`
}

type Result string

const (
	OKResult      = "OK"
	KeyResult     = "Key"
	ExitResult    = "Exit"
	EjectResult   = "Eject"
	InvalidResult = "Invalid"
)

type EndLevel struct {
	Type   string `json:"type"`
	Key    string   `json:"key"`
	Exits  []string `json:"exits"`
	Ejects []string `json:"ejects"`
}

type PlayerScore struct {
	Type   string `json:"type"`
	Name   string   `json:"name"`
	Exits  int    `json:"exits"`
	Ejects int    `json:"ejects"`
	Keys   int    `json:"keys"`
}

type EndGame struct {
	Type   string        `json:"type"`
	Scores []PlayerScore `json:"scores"`
}
