package main

import (
	testJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/tester"
	"os"
)

func main() {
	 outputMsg := tester.TestLevel(testJson.ParseLevelTestJson(os.Stdin))
	 println(outputMsg)
}