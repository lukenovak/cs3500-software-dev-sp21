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

`IsValidMove` validates moves by ensuring all `Actors` are on valid tiles. A valid tile for an Actor is determined by an `Actor`'s `CanOccupyTile` lambda. The function will also check the `maxMoveDistance` field of the `Actor` to verify they have not moved further than allowed. 

```go
IsLevelEnd(currentState GameState) bool { ... }
```
This function is called after a valid move is made. It returns `True` if the level has been ended and `False` if it has not.

```go
IsGameEnd(currentState GameState) bool { ... }
```
This function is called if the level has ended. It returns `True` if the game has been ended and `False` if it has not.
