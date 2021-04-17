package client

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote"
	"github.com/eiannone/keyboard"
	"net"
)

type Player struct {
	Name string
	Posn level.Position2D
	GameWindow fyne.Window
}

// HandleMove gets user input and then writes it to the server
func (player Player) HandleMove(conn net.Conn) {
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
		moveData, err := json.Marshal(remote.PlayerMove{
			Type: "move",
			To:   remote.PointFromPos2d(player.Posn.AddPosition(move)),
		})
		if err != nil {
			panic(err)
		}
		conn.Write(moveData)

		// get result of move and act
		rawData := remote.BlockingRead(conn)
		var result remote.Result
		json.Unmarshal(*rawData, &result)
		fmt.Printf("Result of move was: %v", result)
		switch result {
		case remote.InvalidResult:
			continue
		default:
			return
		}
	}
}
