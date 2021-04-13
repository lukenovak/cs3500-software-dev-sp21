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

const (
	defaultTimeout = 60 * time.Second
)

// ServerPlayerClient represents a player as tracked from the server. Must be loaded with an active TCP Connection!!!
type ServerPlayerClient struct {
	name         string
	nextResponse state.Response
	activeConnection net.Conn
}

// sends the welcome message to the player
func (s ServerPlayerClient) RegisterClient() (actor.Actor, error) {
	rawMsg, err := json.Marshal(remote.NewServerWelcome())
	if err != nil {
		return actor.Actor{}, err
	}
	s.SendJsonMessage(rawMsg)
	return actor.NewPlayerActor(s.name, actor.PlayerType, 2), nil
}

func (s ServerPlayerClient) SendPartialState(layout [][]*level.Tile, actors []actor.Actor, pos level.Position2D) error {

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

func (s ServerPlayerClient) SendMessage(message string, pos level.Position2D) error {
	return s.SendJsonMessage([]byte(message))
}

func (s ServerPlayerClient) GetInput() state.Response {
	panic("implement me")
}

func (s ServerPlayerClient) GetName() string {
	return s.name
}

func (s ServerPlayerClient) SendJsonMessage(message json.RawMessage) error {
	panic("implement me")
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
