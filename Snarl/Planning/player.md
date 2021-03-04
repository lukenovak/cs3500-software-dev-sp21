# MEMORANDUM

#### DATE: March 3, 2021
#### TO: Growl, Inc.
#### FROM: William Victoria and Luke Novak, Software Developers, Ormegland Inc.
#### SUBJECT: Snarl Player Client

---

The memo presents our plan for the Snarl Player interface. It includes a description of the expected functionality of our Player and some functions that we expect to implement as a part of our Player.

What does the Player (client) component need to do:
0. Start the game (this is separate from the Player client)
1. Receive a partial game state from the server
2. Return move input to the server
3. GOTO 1

To differentiate them from the Player type of `Actor`, this Player component will be called a `UserClient`. Since we could have local users or remote users, `UserClient` will be an interface.

```Go
interface UserClient {
    // Sends a new state to the player
    SendPartialState([][]*Tile visibleTiles, []Actor visibleActors) error

    // Sends a message to the player (used for invalid moves);
    SendMessage(msg string) error

    // Waits for a player input then returns the player's action after input
    GetInput() []ClientResponse { ... }
}
```

`ClientResponse` is a struct that contains a client's response with all of the actions they wish to take. We chose to leave this broad so that we could potentially support more actions beyond moving in the future.

```Go
type ClientResponse struct {
    PlayerId    int
    Move        Position2D
    Actions     map[string]{}interface
}
```
