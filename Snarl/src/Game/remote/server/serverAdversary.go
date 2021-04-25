package server

import (
	"bufio"
	"encoding/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"log"
	"net"
)

// Adversary represents a server-side implementation of the AdversaryClient
type Adversary struct {
	currentLevel           level.Level
	name                   string
	adversaryType          int
	currPos                level.Position2D
	activeConnection       net.Conn
	activeConnectionReader *bufio.Reader
}

func NewServerAdversary(name string, adversaryType int, activeConnection net.Conn) *Adversary {
	return &Adversary{
		currentLevel:           level.Level{},
		name:                   name,
		adversaryType:          adversaryType,
		currPos:                level.NewPosition2D(-1, -1),
		activeConnection:       activeConnection,
		activeConnectionReader: bufio.NewReader(activeConnection),
	}
}

func (a *Adversary) CalculateMove(playerPosns []level.Position2D, adversaryPositions []level.Position2D) state.Response {

	var actorPositions []remote.ActorPosition

	// appendToActorPositions converts a list of positions to ActorPositions and appends to actorPositions
	appendToActorPositions := func(actorType string, posns []level.Position2D) {
		for _, posn := range posns {
			actorPositions = append(actorPositions, remote.ActorPosition{
				Type:     "player",
				Name:     "",
				Position: remote.PointFromPos2d(posn),
			})
		}
	}

	appendToActorPositions("player", playerPosns)
	appendToActorPositions("adversary", adversaryPositions)

	// Package the whole thing into a player update and wait for a response
	remote.NewAdversaryUpdateMessage(a.currentLevel, a.currPos, actorPositions)

	moveInput := remote.BlockingRead(a.activeConnectionReader)

	// input should be a Move
	var move remote.PlayerMove
	err := json.Unmarshal(*moveInput, &move)
	if err != nil {
		log.Println("invalid move sent by adversary, moving 0, 0")
		return state.Response{
			PlayerName: a.name,
			Move:       level.NewPosition2D(0, 0),
			Actions:    nil,
		}
	}

	// return the move
	return state.Response{
		PlayerName: a.name,
		Move:       move.To.ToPos2D(),
		Actions:    nil,
	}
}

func (a *Adversary) UpdatePosition(d level.Position2D) {
	a.currPos = d
}

func (a *Adversary) GetName() string {
	return a.name
}

func (a *Adversary) GetType() int {
	return a.adversaryType
}

func (a *Adversary) UpdateLevel(level level.Level) {
	a.currentLevel = level
}
