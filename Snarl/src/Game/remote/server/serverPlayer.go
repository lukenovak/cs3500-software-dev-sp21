package server

import (
	"encoding/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"net"
	"time"
)

// PlayerClient ServerPlayerClient represents a player as tracked from the server. Must be loaded with an active TCP Connection!!!
type PlayerClient struct {
	name             string
	activeConnection net.Conn
	timeout          time.Time
}

// NewPlayerClient creates a PlayerClient according to the given parameters
func NewPlayerClient(name string, activeConnection net.Conn, timeout time.Time) *PlayerClient {
	return &PlayerClient{name: name, activeConnection: activeConnection, timeout: timeout}
}

// RegisterClient reads from the active connection to get the PlayerClient name
func (s *PlayerClient) RegisterClient() (actor.Actor, error) {
	err := s.activeConnection.SetReadDeadline(s.timeout)
	if err != nil {
		return actor.Actor{}, err
	}

	// read from connection to get name
	var nameBytes []byte
	_, err = s.activeConnection.Read(nameBytes)
	if err != nil {
		return actor.Actor{}, err
	}

	// return the correct actor
	s.name = string(nameBytes)
	return actor.NewPlayerActor(s.name, actor.PlayerType, 2), nil
}

func (s *PlayerClient) SendPartialState(layout [][]*level.Tile, actors []actor.Actor, pos level.Position2D) error {

	// TODO: maybe move this out of SendPartialState?
	// Local function that searches through all the tiles in the layout and returns a list of the Objects in those tiles
	getObjectsInLayout := func(layout [][]*level.Tile) []remote.Object {
		var objects []remote.Object
		for _, row := range layout {
			for _, tile := range row {
				if tile.Item != nil {
					var tileType string
					switch tile.Item.Type {
					case level.KeyID:
						tileType = "key"
					case level.UnlockedExit, level.LockedExit:
						tileType = "exit"
					default:
						tileType = "unknown-item"
					}
					objects = append(objects, remote.Object{
						Type:     tileType,
						Position: remote.PointFromPos2d(pos),
					})
				}
			}
		}
		return objects
	}

	// convert game actors to ActorPositions
	var convertedActors []remote.ActorPosition
	for _, a := range actors {
		convertedActors = append(convertedActors, *remote.NewActorPositionFromActor(a))
	}

	// Generate and send the partial state
	partialState := remote.NewPlayerUpdateMessage(tileLayoutToIntLayout(layout), remote.PointFromPos2d(pos), getObjectsInLayout(layout), convertedActors, "update")
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
	return s.SendJsonMessage([]byte(message))
}

func (s *PlayerClient) GetInput() state.Response {
	// error response so that the game can continue if a client is sending bad data
	errorResponse := state.Response{
		PlayerName: s.name,
		Move:       level.NewPosition2D(0, 0),
		Actions:    nil,
	}

	var moveInput []byte
	_, err := s.activeConnection.Read(moveInput)
	if err != nil {
		return errorResponse
	} else {
		// marshall to correct struct then convert to state response
		var move remote.PlayerMove
		err = json.Unmarshal(moveInput, &move)
		if err != nil {
			return errorResponse
		}
		movePoint := *move.To
		return state.Response{
			PlayerName: s.name,
			Move:       movePoint.ToPos2D(),
			Actions:    nil,
		}
	}
}

func (s *PlayerClient) GetName() string {
	return s.name
}

// SendJsonMessage writes a raw json message over the active connection and returns the status of that write
func (s *PlayerClient) SendJsonMessage(message json.RawMessage) error {
	_, err := s.activeConnection.Write(message)
	return err
}

// tileLayoutToIntLayout converts a 2d slice of Tile to a 2d slice of int for network communication
func tileLayoutToIntLayout(tiles [][]*level.Tile) [][]int {
	output := make([][]int, 0)
	for _, tileRow := range tiles {
		outputRow := make([]int, 0)
		for _, tile := range tileRow {
			if tile == nil {
				outputRow = append(outputRow, 0)
			} else {
				outputRow = append(outputRow, tile.Type)
			}
		}
		output = append(output, outputRow)
	}
	return output
}
