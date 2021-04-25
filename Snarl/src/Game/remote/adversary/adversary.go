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
	Client             state.AdversaryClient
	PlayerPositions    []level.Position2D
	AdversaryPositions []level.Position2D
	GameWindow         fyne.Window
	Conn               net.Conn
	ConnReader         *bufio.Reader
}

// HandleMove gets user input and then writes it to the server
func (adversary Adversary) HandleMove() {
	for {
		// send move to server
		move := adversary.Client.CalculateMove(adversary.PlayerPositions, adversary.AdversaryPositions).Move
		fmt.Printf("Moving to: %v", move)
		moveData, err := json.Marshal(remote.PlayerMove{
			Type: "move",
			To:   remote.PointFromPos2d(move),
		})
		if err != nil {
			panic(err)
		}
		adversary.Conn.Write(append(moveData, '\n'))

		// get result of move and act
		rawData := remote.BlockingRead(adversary.ConnReader)
		var result remote.Result
		json.Unmarshal(*rawData, &result)
		fmt.Printf("Result of move was: %s\n", result)
		switch result {
		case remote.InvalidResult:
			continue
		default:
			// update the position and continue
			adversary.Client.UpdatePosition(move)
			return
		}
	}
}
