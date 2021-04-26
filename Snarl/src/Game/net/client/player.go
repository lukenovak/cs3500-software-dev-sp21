package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	snarlNet "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/net"
	"github.com/eiannone/keyboard"
	"net"
)

// A Player is the client-side representation of a Player. It gets the raw input from the human player and writes
// it back to the server to update the game state
type Player struct {
	Name string
	Posn level.Position2D
	GameWindow fyne.Window
}

// HandleMove gets user input and then writes it to the server
func (player Player) HandleMove(conn net.Conn, connReader *bufio.Reader) {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := keyboard.Close()
		if err != nil {
			panic(err)
		}
	}()
	for {
		// get move from user
		move := level.NewPosition2D(0, 0)
		for {
			fmt.Printf("current move is %d, %d\n", move.Row, move.Col)
			_, key, _ := keyboard.GetKey()
			if key == keyboard.KeyEnter {
				break
			}
			if key == keyboard.KeyArrowRight {
				move = level.NewPosition2D(move.Row, move.Col+1)
			}
			if key == keyboard.KeyArrowLeft {
				move = level.NewPosition2D(move.Row, move.Col-1)
			}
			if key == keyboard.KeyArrowUp {
				move = level.NewPosition2D(move.Row-1, move.Col)
			}
			if key == keyboard.KeyArrowDown {
				move = level.NewPosition2D(move.Row+1, move.Col)
			}
		}

		// send move to server
		moveData, err := json.Marshal(snarlNet.PlayerMove{
			Type: "move",
			To:   snarlNet.PointFromPos2d(player.Posn.AddPosition(move)),
		})
		if err != nil {
			panic(err)
		}
		conn.Write(append(moveData, '\n'))

		// get result of move and act
		rawData := snarlNet.BlockingRead(connReader)
		var result snarlNet.Result
		json.Unmarshal(*rawData, &result)
		fmt.Printf("Result of move was: %s\n", result)
		switch result {
		case snarlNet.InvalidResult:
			continue
		default:
			return
		}
	}
}
