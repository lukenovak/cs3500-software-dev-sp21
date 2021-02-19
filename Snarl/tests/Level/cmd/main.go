package main

import (
	"fmt"
	testJson "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Level/json"
	"os"
)

func main() {
	fmt.Printf("%+v\n", testJson.ParseLevelTestJson(os.Stdin))
}