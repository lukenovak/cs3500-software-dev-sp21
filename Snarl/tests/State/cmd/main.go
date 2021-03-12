package main

import (
	"encoding/json"
	"fmt"
	stateTest "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/State"
	"os"
)

func main() {
	outputArray := stateTest.TestUpdateState(stateTest.ParseStateTestJson(os.Stdin))
	outputJsonBytes, err := json.Marshal(outputArray)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(outputJsonBytes))
}
