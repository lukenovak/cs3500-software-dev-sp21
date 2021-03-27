package main

import (
	"encoding/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Manager"
	"os"
)

func main() {
	outputMessage, err := json.Marshal(Manager.Test(Manager.ParseManagerInput(os.Stdin)))
	if err != nil {
		panic(err)
	}
	_, _ = os.Stdout.Write(outputMessage)
}