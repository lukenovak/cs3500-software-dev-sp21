package adversary

import (
	"bufio"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	snarlNet "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/net"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"net"
)

// Adversary represents a server-side adversary that the server-run game manager can interact with. We do not need to
// know what type of Adversary this is since this representation calls out to the client to get a move
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
		moveData, err := json.Marshal(snarlNet.PlayerMove{
			Type: "move",
			To:   snarlNet.PointFromPos2d(move),
		})
		if err != nil {
			panic(err)
		}
		adversary.Conn.Write(append(moveData, '\n'))

		// get result of move and act
		rawData := snarlNet.BlockingRead(adversary.ConnReader)
		var result snarlNet.Result
		json.Unmarshal(*rawData, &result)
		fmt.Printf("Result of move was: %s\n", result)
		switch result {
		case snarlNet.InvalidResult:
			continue
		default:
			// update the position and continue
			adversary.Client.UpdatePosition(move)
			return
		}
	}
}
