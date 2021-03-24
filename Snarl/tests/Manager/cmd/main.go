package main

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/tests/Manager"
	"os"
)

func main() {
	names, _, levelObj, numMoves, _, _ := Manager.ParseManagerInput(os.Stdin)

	println("players are:")
	for _, name := range names {
		println(name)
	}
	println("\nrooms are:")
	for _, room := range levelObj.Rooms {
		println(fmt.Sprintf("%d, %d", room.Origin[0], room.Origin[1]))
	}
	println(levelObj.Rooms[0].Type)
	println("\nMax moves:")
	println(numMoves)
}