# Traveller Specification

This document is a specification for the Traveller package. This program will be written in [Go](https://golang.org/) using version `1.15.7`.

## General
The `traveller` package supports traveling functionality in a game, and as such all of its functions
and structures should be contained in a single Go package called `traveller`

This package should contain the required data types to represent a network of towns. 
Each town should have a name and be able to contain a named character. 
This package should contain the required functions to create a town Network. 
This package should contain a function to determine whether a given Character can travel from one Town to another
within the Network without passing through another Town with a Character in it.

## Structs

These `structs` are used to represent the data for the town network.	

A network represents a collection of towns that are connected to each other, and is
simple an array of Town objects
```Go
type Network []Town;
```
A Town struct represents a named town. Towns can have Characters, and they
can either be initialized with Characters or be made `nil` if there is no
character at the time it is initialized
```Go
type Town struct {
	Name String
	Character Character
	AdjacentTowns []Town
}
```
A Character represents a named character. Characters need only to keep track
of their names, but are represented by a struct for future extensibility.
```Go
type Character struct {
	Name String
}
```

## Functions
These are the necessary publicly-available functions that the `traveler` package should offer.
The functionality of each of these functions is given by the comments above them. 
Implementation will be left to the team.

```Go
// Creates a new Town struct in the given Network with the given name
func CreateTown(String name, Network n) : Town {...}

// Link two towns by adding each other to their AdjacentTowns lists
func LinkTowns(Town t1, Town t2): void {...}

// Unlinks two connected towns by removing each other from their AdjacentTowns lists
func UnlinkTowns(Town t1, Town t2): void {...}

// Creates a new Character with the given name in a town if the town does have a character
func PlaceNewCharacter(Town t, String name): (Character, Error) {...}

// Places an existing Character into a town if the town does not have a character
func PlaceExistingCharacter(Town t, Character c) {...}

// Removes a Character from a given town, returning that character
func RemoveCharacterFromTown(Town t): Character {...}

// This function determines if the given Character can travel to the given Town in the given Network
// Without passing through any other Town with a Character
func CanTravel(Character c, Town t, Network n) bool {...}
```
