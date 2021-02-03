package main

import (
	"container/list"
	"errors"
)

//Character is represented by a name
type Character struct {
	Name string
}

//Town manages characters within and adjacent towns
type Town struct {
	Name          string
	Character     *Character
	AdjacentTowns []*Town
}

//Network represents a list of towns
type Network []*Town

// CreateTown creates a new Town struct in the given Network with the given name
func CreateTown(name string, n Network) (Network, *Town) {
	town := Town{Name: name}
	return append(n, &town), &town
}

// LinkTowns link two towns by adding each other to their AdjacentTowns lists
func LinkTowns(t1 *Town, t2 *Town) {
	t1.AdjacentTowns = append(t1.AdjacentTowns, t2)
	t2.AdjacentTowns = append(t2.AdjacentTowns, t1)
}

// UnlinkTowns unlinks two connected towns by removing each other from their AdjacentTowns lists
func UnlinkTowns(t1 *Town, t2 *Town) {
	t1.AdjacentTowns = remove(t1.AdjacentTowns, t2)
	t2.AdjacentTowns = remove(t2.AdjacentTowns, t1)
}

// PlaceNewCharacter creates a new Character with the given name in a town if the town does have a character
func PlaceNewCharacter(t *Town, name string) (*Character, error) {
	if isTownEmpty(*t) {
		newCharacter := Character{Name: name}
		t.Character = &newCharacter
		return &newCharacter, nil
	}
	return nil, errors.New("town already has character in it")
}

// PlaceExistingCharacter places an existing Character into a town if the town does not have a character
func PlaceExistingCharacter(t *Town, c *Character) {
	if isTownEmpty(*t) {
		t.Character = c
	}
}

// RemoveCharacterFromTown removes a Character from a given town, returning that character
func RemoveCharacterFromTown(t *Town) *Character {
	removedCharacter := t.Character
	t.Character = nil
	return removedCharacter
}

// CanTravel this function determines if the given Character can travel to the given Town in the given Network
// Without passing through any other Town with a Character
func CanTravel(c *Character, t *Town, n Network) bool {
	// Find starting town
	var start *Town
	for i := range n {
		town := n[i]
		if town.Character == c {
			start = town
			break
		}
	}

	// Check if start is the destination
	if t == start {
		return true
	}

	// Perform BFS on the network
	queue := list.New()
	queue.PushBack(start)
	visited := make(map[*Town]bool)
	for queue.Len() > 0 {
		front := queue.Front()
		curr := front.Value.(*Town)
		visited[curr] = true
		if isTownEmpty(*curr) || curr == start {
			if curr == t {
				return true
			}
			for i := range curr.AdjacentTowns {
				neighbor := curr.AdjacentTowns[i]
				if contains := visited[neighbor]; !contains {
					queue.PushBack(neighbor)
				}
			}
		}
		queue.Remove(front)
	}

	return false
}

// Remove a town from a townlist
func remove(townList []*Town, element *Town) []*Town {
	result := []*Town{}
	for i := range townList {
		if townList[i] != element {
			result = append(result, townList[i])
		}
	}
	return result
}

// Checks if a town is empty (ie no character is in this town)
func isTownEmpty(t Town) bool {
	// if town character is defaulted, it is empty
	return t.Character == nil
}

// Checks if towns are equal
func isTownEqual(t1 Town, t2 Town) bool {
	return t1.Name == t2.Name &&
		t1.Character == t2.Character
}
