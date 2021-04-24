package adversary

import (
	"bufio"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"net"
)

type Adversary struct {
	Client          state.AdversaryClient
	PlayerPositions []level.Position2D
	GameWindow      fyne.Window
}

// HandleMove gets user input and then writes it to the server
func (adversary Adversary) HandleMove(conn net.Conn, connReader *bufio.Reader) {
	for {
		// send move to server
		moveData, err := json.Marshal(remote.PlayerMove{
			Type: "move",
			To:   remote.PointFromPos2d(adversary.Client.CalculateMove(adversary.PlayerPositions).Move),
		})
		if err != nil {
			panic(err)
		}
		conn.Write(append(moveData, '\n'))

		// get result of move and act
		rawData := remote.BlockingRead(connReader)
		var result remote.Result
		json.Unmarshal(*rawData, &result)
		fmt.Printf("Result of move was: %s\n", result)
		switch result {
		case remote.InvalidResult:
			continue
		default:
			return
		}
	}
}
