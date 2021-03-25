package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/internal/render"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"os"
)

// struct for a local player
type LocalClient struct {
	Name string
}

func (player LocalClient) RegisterClient() error {
	return nil
}

// TODO: Include this functionality in a future Milestone
func (player LocalClient) SendPartialState(tiles [][]*level.Tile, actors []actor.Actor) error {
	print(render.ASCIILevel(level.Level{Tiles: tiles}))
	return nil
}

// TODO: Include this functionality in a future Milestone
func (player LocalClient) SendMessage(message string) error {
	println(message)
	return nil
}

// TODO: Include this functionality in a future Milestone
func (player LocalClient) GetInput() Response {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter move [row, col]: ")
	move, _ := reader.ReadString('\n')
	var moveData []interface{}
	json.Unmarshal([]byte(move), moveData)
	return Response{
		PlayerId:   0,
		PlayerName: player.Name,
		Move: level.Position2D{
			Row: moveData[0].(int),
			Col: moveData[1].(int),
		},
		Actions: nil,
	}
}

func (player LocalClient) GetName() string {
	return player.Name
}
