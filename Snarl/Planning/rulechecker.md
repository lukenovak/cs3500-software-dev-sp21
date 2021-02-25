# MEMORANDUM

#### DATE: February 24, 2021
#### TO: Growl, Inc.
#### FROM: William Victoria and Luke Novak, Software Developers, Ormegland Inc.
#### SUBJECT: Snarl Rule Checker

---

This memo presents our plan for the Rule Checker component that will be used in our implementation of Snarl. Our Rule Checker will take the form of a package with public functions that will be called by the Game Manager. The Rule Checker has three primary functions:

1. Determine if moves taken by actors are valid.
2. Determine if the level is over.
3. Determine if the game is over.

### Public Functions

```go
IsValidMove(oldGamestate GameState, newGamestate GameState) bool { ... }
```

This first function is called by the Game Manager at the end of each move with the game state before the move is made and the game state after the move is made. It returns a boolean, if it is true we continue with the new game state and if it is false we discard the new game state.

`IsValidMove` validates moves by ensuring all `Actors` are on valid tiles. 
To better support a wide variety of movement types, we will be augmenting our proposed `Actor` struct from the previous
design memo. The improved class will include two new fields: 
1. a lambda called `CanOccupyTile` that takes in a tile Type and returns
a `bool` describing whether or not the `Actor` can be on that type of tile. We decided to make this a lambda because
it allows for rich movement functionality, should future additions be able to change an `Actor`'s movement capabilities. 
   
2. an integer field called `MaxMoveDistance` that contains the maximum distance that the `Actor` can move in one turn.
This distance is measured as a manhattan distance.
   
The `IsValidMove` function will also need to verify that no two `Actors` are occupying the same tile at the end of the
move. It will do this using private sub-functions.

This function will also need to verify that the interactions carried out are correct. For now, this means confirming
that the status of the exits (locked versus unlocked) is correct, but could be extended in the future for other items
and potentially player inventory systems.

```go
IsLevelEnd(currentState GameState) bool { ... }
```
This function is called after a valid move is made. It checks to see if the level has ended. 
It returns `true` if the level has been ended and `false` if it has not.

```go
IsGameEnd(currentState GameState) bool { ... }
```
This function is called from the game manager if `IsLevelEnd` returns `true`. 
It returns `true` if the game has been ended (i.e. the players have lost or there are no more levels left)
and `false` if it has not.
