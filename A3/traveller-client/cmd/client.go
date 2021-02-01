package main

import (
	"encoding/json"
	"os"

	"github.ccs.neu.edu/CS4500-S21/Ormegland/A3/traveller-client/parse"
)

func main() {
	commandList := parse.ParseCommands(os.Stdin)

}

func ExecuteCommands(commandList []parse.Command) {
	for _, command := range commandList {
		switch command.Command {
		case "roads":
			var roads parse.RoadArray
			json.Unmarshal(command.Params, &roads)
			// write network create func
		case "place":
			// write place character func
		case "passage-safe?":
			// write solve query func
		}
	}
}
