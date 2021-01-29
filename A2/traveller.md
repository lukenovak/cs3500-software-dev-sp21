# Traveller Specification

This document is a specification for the Traveller module. This program will be written in [Go](https://golang.org/) using version `1.15.7`.

## General

This module should contain the required data types to represent a network of towns. Each town should have a name and be able to contain a named character. This module should contain the required functions to create a town network. This module should contain a function to determine whether a given character can reach a given town in a town network.

## Structs

These `structs` are used to represent the data for the town network.	

```Go
type Network []Town;

type Town struct {
	Name String
	Character Character
	AdjacentTowns []Town
}

type Character struct {
	Name String
}
```

## Functions

```Go
// Create a town in the network with the given name
func CreateTown(String name, Network n) : Town {...}

// Link two towns by mutating their adjacency lists
func LinkTowns(Town t1, Town t2): void {...}

// Place a Character with the given name in a town
func PlaceCharacter(Town t, String name): (Character, Error) {...}

// This function determines if the given Character can reach the given Town in the given Network
func CanTravel(Character c, Town t, Network n) bool {...}
```
