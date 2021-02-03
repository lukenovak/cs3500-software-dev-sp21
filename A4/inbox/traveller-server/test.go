package main

import "fmt"

func main() {
	network := []*Town{}
	network, a := CreateTown("A", network)
	network, b := CreateTown("B", network)
	network, c := CreateTown("C", network)
	network, d := CreateTown("D", network)

	LinkTowns(a, b)
	UnlinkTowns(b, a)

	char1, _ := PlaceNewCharacter(a, "mom")
	char2, _ := PlaceNewCharacter(c, "dad")
	PlaceExistingCharacter(b, char1)
	RemoveCharacterFromTown(a)
	RemoveCharacterFromTown(b)

	LinkTowns(a, b)
	LinkTowns(b, c)
	LinkTowns(c, a)
	LinkTowns(c, d)

	PlaceExistingCharacter(a, char1)

	fmt.Println(CanTravel(char1, d, network))

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)

	fmt.Println(char1)
	fmt.Println(char2)
}
