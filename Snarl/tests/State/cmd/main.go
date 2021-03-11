package main

import (
	"fmt"
	stateTest "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/State"
	"os"
)

func main() {
	state, name, point := stateTest.ParseStateTestJson(os.Stdin)
	fmt.Printf("%+v\n%s\n%+v\n", state, name, point)
}
