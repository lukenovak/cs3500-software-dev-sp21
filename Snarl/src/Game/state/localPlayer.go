package state

import (
	"bufio"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/internal/render"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.com/eiannone/keyboard"
	"os"
)

// struct for a local player
// Local Players get input from the console, and render to a fyne window
type LocalClient struct {
	Name string
	GameWindow fyne.Window
}

// produces a player actor corresponding with the name of this client
func (player *LocalClient) RegisterClient() (actor.Actor, error) {
	player.GameWindow = fyne.CurrentApp().NewWindow("snarl client")
	return actor.NewPlayerActor(player.Name, actor.PlayerType, 2).MoveActor(level.NewPosition2D(-1, -1)), nil
}

// Gets a partial state and renders it to the fyne game
func (player *LocalClient) SendPartialState(tiles [][]*level.Tile, actors []actor.Actor, pos level.Position2D) error {
	render.GuiState(tiles, actors, actors, player.GameWindow)
	return nil
}

// Prints the sent message to the console
func (player *LocalClient) SendMessage(message string, pos level.Position2D) error {
	println(message)
	return nil
}

// asks for a move in JSON format
func (player *LocalClient) GetInput() Response {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter move [row, col] for player %s: ", player.Name)
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


// struct for a local player controlled by the keyboard
// Key clients work exactly the same as normal clients, but the input is dicated by the keyboard
type LocalKeyClient struct {
	Name string
	GameWindow fyne.Window
}

func (player *LocalKeyClient) RegisterClient() (actor.Actor, error) {
	player.GameWindow = fyne.CurrentApp().NewWindow("snarl client")
	if err := keyboard.Open(); err != nil {
		return actor.Actor{}, err
	}
	return actor.NewPlayerActor(player.Name, actor.PlayerType, 2).MoveActor(level.NewPosition2D(-1, -1)), nil
}

// SendPartialState works the same as the normal client
func (player *LocalKeyClient) SendPartialState(tiles [][]*level.Tile, actors []actor.Actor, pos level.Position2D) error {
	render.GuiState(tiles, actors, actors, player.GameWindow)
	return nil
}

func (player *LocalKeyClient) SendMessage(message string, pos level.Position2D) error {
	println(message)
	return nil
}

// GetInput Reads from the keyboard until the user hits enter, which locks in their move
func (player *LocalKeyClient) GetInput() Response {
	player.GameWindow.RequestFocus()
	move := level.NewPosition2D(0, 0)
	for {
		fmt.Printf("current move for player %s is %d, %d\n", player.Name, move.Row, move.Col)
		_, key, _ := keyboard.GetKey()
		if key == keyboard.KeyEnter {
			break
		}
		if key == keyboard.KeyArrowRight {
			move = level.NewPosition2D(move.Row, move.Col + 1)
		}
		if key == keyboard.KeyArrowLeft {
			move = level.NewPosition2D(move.Row, move.Col - 1)
		}
		if key == keyboard.KeyArrowUp {
			move = level.NewPosition2D(move.Row - 1, move.Col)
		}
		if key == keyboard.KeyArrowDown {
			move = level.NewPosition2D(move.Row + 1, move.Col)
		}
	}
	return Response{
		PlayerName: player.Name,
		Move: move,
		Actions: nil,
	}
}

func (player *LocalKeyClient) GetName() string {
	return player.Name
}