package main

import (
	testJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/tester"
	"os"
)

func main()  {
	output := tester.TestLevelTileData(testJson.ParseLevelTileDataTestJson(os.Stdin))
	println(output)
}