# Traveller Specification

This document is a specification for the Traveller module. This program will be written in [Go](https://golang.org/) using version `1.15.7`.

## General

This module should contain the required data types to represent a network of towns. Each town should have a name and be able to contain a named character. This module should contain the required functions to create a town network. This module should contain a function to determine whether a given character can reach a given town in a town network.

## Structs

```Go
type Network []Town;

type Town struct {
	name String
	character Character
	adjacentTowns []Town
}

type Character struct {
	name String
}
```

## Functions

```Go
// functions to create a town network
func AddTown(String s, Network n) : Town {...}
func AddLink(Town t1, Town t2): void {...}
func PlaceCharacter(Town t, String name): (Character, Error) {...}

// function to determine if the given Character can reach the given Town in the given Network
func CanTravel(Character c, Town t, Network n) {...}
```
