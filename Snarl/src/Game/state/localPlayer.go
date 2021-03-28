package state

import (
	"bufio"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/internal/render"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"os"
)

// struct for a local player
type LocalClient struct {
	Name string
	GameWindow fyne.Window
}

func (player *LocalClient) RegisterClient() (actor.Actor, error) {
	player.GameWindow = fyne.CurrentApp().NewWindow("snarl client")
	return actor.NewWalkableActor(player.Name, actor.PlayerType, 2).MoveActor(level.NewPosition2D(-1, -1)), nil
}

func (player *LocalClient) SendPartialState(tiles [][]*level.Tile, actors []actor.Actor, pos level.Position2D) error {
	render.GuiState(tiles, actors, actors, player.GameWindow)
	return nil
}

func (player *LocalClient) SendMessage(message string, pos level.Position2D) error {
	println(message)
	return nil
}

func (player *LocalClient) GetInput() Response {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter move [row, col]: ")
	move, _ := reader.ReadBytes('\n')
	var moveData [2]int
	json.Unmarshal(move, &moveData)
	return Response{
		PlayerName: player.Name,
		Move: level.Position2D{
			Row: moveData[0],
			Col: moveData[1],
		},
		Actions: nil,
	}
}

func (player *LocalClient) GetName() string {
	return player.Name
}
