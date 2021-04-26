package net

import (
	"encoding/json"
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
)

const ServerVersion = "0.0.1"

// ServerWelcome represents a welcome message in JSON
type ServerWelcome struct {
	Type string `json:"type"`
	Info string `json:"info"`
}

// NewServerWelcomeMessage is a constructor that creates a new welcome message with the current server version
func NewServerWelcomeMessage() json.RawMessage {
	welcome := &ServerWelcome{
		Type: "welcome",
		Info: fmt.Sprintf("Snarl Server version %s", ServerVersion),
	}
	message, _ := json.Marshal(welcome)
	return message
}

// StartLevel represents a start-level message in JSON
type StartLevel struct {
	Type    string   `json:"type"`
	Level   int      `json:"level"`
	Players []string `json:"players"`
}

// NewStartLevel generates an StartLevel message from a given level number and player list
func NewStartLevel(levelNum int, playerList []string) *StartLevel {
	return &StartLevel{
		Type:    "start-level",
		Level:   levelNum,
		Players: playerList,
	}
}

// Point represents a row, column point. Maps nicely to level.Position2D
type Point [2]int

// ToPos2D returns the Position2D equivalent to the Point it is called on
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

// ActorPosition represents an "actor-position" json message
type ActorPosition struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Position Point  `json:"position"`
}

// Layout represents a grid of tiles
type Layout [][]int

// NewActorPositionFromActor creates an ActorPosition object from an actor.Actor
func NewActorPositionFromActor(a actor.Actor) *ActorPosition {
	return &ActorPosition{
		Type:     a.GetTypeAsString(),
		Name:     a.Name,
		Position: PointFromPos2d(a.Position),
	}
}

// PlayerUpdateMessage represents a player-update json message
type PlayerUpdateMessage struct {
	Type     string          `json:"type"`
	Layout   Layout          `json:"layout"`
	Position Point           `json:"position"`
	Objects  []Object        `json:"objects"`
	Actors   []ActorPosition `json:"actors"`
	Message  string          `json:"message"`
}

// NewPlayerUpdateMessage constructs a PlayerUpdateMessage from the necessary fields
func NewPlayerUpdateMessage(layout Layout, position Point, objects []Object, actors []ActorPosition, message string) *PlayerUpdateMessage {
	return &PlayerUpdateMessage{
		Type:     "player-update",
		Layout:   layout,
		Position: position,
		Objects:  objects,
		Actors:   actors,
		Message:  message,
	}
}

// Room represents the data for a game room, as part of the Level json message
type Room struct {
	Type   string `json:"type"`
	Origin Point  `json:"origin"`
	Bounds Bounds `json:"bounds"`
	Layout Layout `json:"layout"`
}

// Hallway contains the data for a game hallway, as part of the Level json message
type Hallway struct {
	From      Point   `json:"from"`
	To        Point   `json:"to"`
	Waypoints []Point `json:"waypoints"`
}

// Bounds represents the bounds of a json Level
type Bounds struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}

// Level is a json representation of a level, used to send adversaries full levels.
type Level struct {
	Rooms    []Room    `json:"rooms"`
	Hallways []Hallway `json:"hallways"`
}

// AdversaryUpdateMessage represents an "adversary-update" JSON messaged, which updates adversaries on the current
// game state
type AdversaryUpdateMessage struct {
	Type     string          `json:"type"`
	Level    Level           `json:"level"`
	Position Point           `json:"position"`
	Objects  []Object        `json:"objects"`
	Actors   []ActorPosition `json:"actors"`
	Message  string          `json:"message"`
}

// NewAdversaryUpdateMessage constructs an AdversaryUpdateMessage from a game Level, a Position2D, and some ActorPositions
func NewAdversaryUpdateMessage(gameLevel level.Level, pos level.Position2D, actors []ActorPosition) *AdversaryUpdateMessage {
	return &AdversaryUpdateMessage{
		Type:     "adversary-update",
		Level:    ConvertLevelToTestLevel(gameLevel),
		Position: PointFromPos2d(pos),
		Objects:  GetObjectsFromLevel(gameLevel),
		Actors:   actors,
		Message:  "update",
	}
}

// PlayerMove represents a move sent from the client to the server
type PlayerMove struct {
	Type string `json:"type"`
	To   Point  `json:"to"`
}

// A Result is the resulting outcome of a player's move
type Result string

// These constants represent the results that the server should be able to send the client
const (
	OKResult      = "OK"
	KeyResult     = "Key"
	ExitResult    = "Exit"
	EjectResult   = "Eject"
	InvalidResult = "Invalid"
)

// EndLevel represents an "end-level" json, which should end a level on the client side.
type EndLevel struct {
	Type   string   `json:"type"`
	Key    string   `json:"key"`
	Exits  []string `json:"exits"`
	Ejects []string `json:"ejects"`
}

// A PlayerScore represents a single entry on the scoreboard, as represented in JSON
type PlayerScore struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Exits  int    `json:"exits"`
	Ejects int    `json:"ejects"`
	Keys   int    `json:"keys"`
}

// ToString converts a PlayerScore to a string entry for the leaderboard
func (score PlayerScore) ToString() string {
	return fmt.Sprintf("%v, %v, %v, %v\n", score.Name, score.Exits, score.Keys, score.Ejects)
}

// EndGame represents an end-game json message, that should signal the end of the game to clients
type EndGame struct {
	Type   string        `json:"type"`
	Scores []PlayerScore `json:"scores"`
}

// TypedJson is used to unmarshal an unknown json to determine its type
type TypedJson struct {
	Type string `json:"type"`
}
