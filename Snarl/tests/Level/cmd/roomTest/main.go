package main

import (
	"encoding/json"
	"fmt"
	testJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/tester"
	"os"
)

func main() {
	outputMsg := tester.TestRoomTraversables(testJson.ParseRoomTestJson(os.Stdin))
	outputJsonBytes, err := json.Marshal(outputMsg)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(outputJsonBytes))
}
