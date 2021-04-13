package remote

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

type ServerWelcome struct {
	Type string `json:"type"`
	Info string `json:"info"`
}

func NewServerWelcome() *ServerWelcome {
	return &ServerWelcome{
		Type: "welcome",
		Info: fmt.Sprintf("Snarl Server version %s\n", ServerVersion),
	}
}

type Name string

type StartLevel struct {
	Type    string `json:"type"`
	Level   int    `json:"level"`
	Players []Name `json:"players"`
}

// Point represents a row, column point. Maps nicely to level.Position2D
type Point [2]int

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
		Type: "actor-position",
		Name: a.Name,
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
	To   *Point `json:"to"`
}

type Result string

type EndLevel struct {
	Type   string `json:"type"`
	Key    Name   `json:"key"`
	Exits  []Name `json:"exits"`
	Ejects []Name `json:"ejects"`
}

type PlayerScore struct {
	Type   string `json:"type"`
	Name   Name   `json:"name"`
	Exits  int    `json:"exits"`
	Ejects int    `json:"ejects"`
	Keys   int    `json:"keys"`
}

type EndGame struct {
	Type   string        `json:"type"`
	Scores []PlayerScore `json:"scores"`
}
