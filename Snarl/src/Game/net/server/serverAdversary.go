package server

import (
	"bufio"
	"encoding/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/net"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"log"
	"net"
)

// Adversary represents a server-side implementation of the AdversaryClient. It interfaces with the client-side adversary
type Adversary struct {
	currentLevel           level.Level
	name                   string
	adversaryType          int
	currPos                level.Position2D
	activeConnection       net.Conn
	activeConnectionReader *bufio.Reader
}

// NewServerAdversary is a constructor that generates a new Adversary from a limited set of information
// The activeConnectionReader is extrapolated from the activeConnection, and the currPos is set to -1, -1 to match
// the Game Manager. The level will be set later, so the new Adversary has an empty level.
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

	var actorPositions []net.ActorPosition

	// appendToActorPositions converts a list of positions to ActorPositions and appends to actorPositions
	appendToActorPositions := func(actorType string, posns []level.Position2D) {
		for _, posn := range posns {
			actorPositions = append(actorPositions, net.ActorPosition{
				Type:     "player",
				Name:     "",
				Position: net.PointFromPos2d(posn),
			})
		}
	}

	appendToActorPositions("player", playerPosns)
	appendToActorPositions("adversary", adversaryPositions)

	// Package the whole thing into a player update and wait for a response
	adversaryMessage, _ := json.Marshal(net.NewAdversaryUpdateMessage(a.currentLevel, a.currPos, actorPositions))
	err := a.SendJsonMessage(adversaryMessage)
	if err != nil {
		log.Println("unable to communicate with adversary. moving 0, 0")
		return state.Response{
			PlayerName: a.name,
			Move:       level.NewPosition2D(0, 0),
			Actions:    nil,
		}
	}

	log.Println("sending move command to an adversary")
	// tell the net adversary to return a move
	a.activeConnection.Write([]byte("\"move\"\n"))

	moveInput := net.BlockingRead(a.activeConnectionReader)

	// input should be a Move
	var move net.PlayerMove
	err = json.Unmarshal(*moveInput, &move)
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

// SendJsonMessage writes a raw json message over the active connection and returns the status of that write
func (a *Adversary) SendJsonMessage(message json.RawMessage) error {
	log.Printf("sent message %s to adversary %s", string(message), a.name)
	// endl terminator
	message = append(message, '\n')
	_, err := a.activeConnection.Write(message)
	return err
}
