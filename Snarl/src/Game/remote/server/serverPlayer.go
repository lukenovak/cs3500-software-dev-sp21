package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"log"
	"net"
	"time"
)

// PlayerClient ServerPlayerClient represents a player as tracked from the server. Must be loaded with an active TCP Connection!!!
type PlayerClient struct {
	name                   string
	activeConnection       net.Conn
	activeConnectionReader *bufio.Reader
	timeout                time.Duration
	currPosition           level.Position2D
}

// NewPlayerClient creates a PlayerClient according to the given parameters
func NewPlayerClient(name string, activeConnection net.Conn, timeout time.Duration) *PlayerClient {
	return &PlayerClient{
		name: name,
		activeConnection: activeConnection,
		activeConnectionReader: bufio.NewReader(activeConnection),
		timeout: timeout,
	}
}

// RegisterClient reads from the active connection to get the PlayerClient name
func (s *PlayerClient) RegisterClient() (actor.Actor, error) {
	newActor := actor.NewPlayerActor(s.name, actor.PlayerType, 2)
	s.currPosition = newActor.Position
	return newActor, nil
}

func (s *PlayerClient) SendPartialState(layout [][]*level.Tile, actors []actor.Actor, pos level.Position2D) error {

	// update position
	s.currPosition = pos

	// convert the actors to absolute coordinates to meet protocol standards
	var absoluteActors []actor.Actor
	for _, a := range actors {
		absoluteActors = append(absoluteActors, a.MoveActor(level.NewPosition2D(s.currPosition.Row - 2 + a.Position.Row, s.currPosition.Col - 2 + a.Position.Col)))
	}

	// convert game actors to ActorPositions
	var convertedActors []remote.ActorPosition
	for _, a := range absoluteActors {
		convertedActors = append(convertedActors, *remote.NewActorPositionFromActor(a))
	}

	// Generate and send the partial state
	partialState := remote.NewPlayerUpdateMessage(remote.TilesToArray(layout), remote.PointFromPos2d(pos), remote.GetObjectsFromLayout(layout), convertedActors, fmt.Sprintf("%s moved", s.name))
	message, err := json.Marshal(partialState)
	if err != nil {
		// If we cannot marshal a partial state into a communicable json, that is a fatal error and we crash the server
		// TODO: Softer error handling
		s.SendMessage("Fatal Server Error", level.NewPosition2D(-1, -1))
		s.activeConnection.Close()
		panic(err)
	}
	return s.SendJsonMessage(message)
}

func (s *PlayerClient) SendMessage(message string, pos level.Position2D) error {
	// start level is a special case, so we handle it here
	if message == "start-level" {
		messageJson := remote.StartLevel{
			Type:    "start-level",
			Level:   0,
			Players: nil,
		}
		msgBytes, _ := json.Marshal(messageJson)
		return s.SendJsonMessage(msgBytes)
	}

	message = fmt.Sprintf("\"%s\"", message)
	return s.SendJsonMessage([]byte(message))
}

func (s *PlayerClient) GetInput() state.Response {
	// error response so that the game can continue if a client is sending bad data
	errorResponse := state.Response{
		PlayerName: s.name,
		Move:       level.NewPosition2D(0, 0),
		Actions:    nil,
	}

	// prompt the player for a move
	log.Println("sending move command")
	s.activeConnection.Write([]byte("\"move\"\n"))

	moveInput := remote.BlockingRead(s.activeConnectionReader)

	// marshall to correct struct then convert to state response
	var move remote.PlayerMove
	err := json.Unmarshal(*moveInput, &move)
	if err != nil {
		return errorResponse
	}
	movePoint := move.To
	relativeMove := level.NewPosition2D(movePoint[0]-s.currPosition.Row, movePoint[1]-s.currPosition.Col)
	return state.Response{
		PlayerName: s.name,
		Move:       relativeMove,
		Actions:    nil,
	}
}

func (s *PlayerClient) GetName() string {
	return s.name
}

// SendJsonMessage writes a raw json message over the active connection and returns the status of that write
func (s *PlayerClient) SendJsonMessage(message json.RawMessage) error {
	log.Printf("sent message %s to %s", string(message), s.name)
	// endl terminator
	message = append(message, '\n')
	_, err := s.activeConnection.Write(message)
	return err
}

func (s *PlayerClient) AsUserClient() state.UserClient {
	return s
}

func (s *PlayerClient) CloseConnection() error {
	return s.activeConnection.Close()
}
