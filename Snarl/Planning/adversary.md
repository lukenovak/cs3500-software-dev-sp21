# MEMORANDUM

#### DATE: March 11, 2021
#### TO: Growl, Inc.
#### FROM: William Victoria and Luke Novak, Software Developers, Ormegland Inc.
#### SUBJECT: Snarl Observer

---

This memo presents our plan for the Adversary component.

Adversaries are very similar to players and observers with a few key
differences.
1. Adversaries have full access to a level at the beginning of the level
2. Adversaries get full access to player locations before they need to move

Note, adversary clients (which will move the adversaries) should be distinguished
from adversary Actors, which are an internal game state representation of Adversaries.
This "Adversary Client" is basically a controller that determines where one game state
Adversary should move on any given turn.

### Planned Design

```Go
interface AdversaryClient {
    // Returns a Response with the best relative move
    func (a AdversaryClient) CalculateMove(playerPosns []Position2d) Response
}

type ExampleAdvClient struct {
    Name       string
    LevelData  level.Level
    MoveDistance int
    CurrentPosn   Position2D
}
```

Our adversary is a wrapper interface containing the necessary information to
calculate the best possible move given the adversary's current
position. The last piece of data added that is added when the function is
called are the player positions. This ensures both that the player positions
are up to date, and that the Adversary Client does not have access to the data
when they are not about to have their turn.

The decision to make the Adversary Client an interface gives us additional flexibility
in creating clients to determine Adversary moves. We can code different behavior for
different adversaries or difficulties. We could have an implementing `struct` that
calls out over a TCP connection to get the best move. We could also put test behavior
here to be able to test that the Adversary Client hooks in well with our game loop.

The adversaries will move after the players in our game loop function. The call to the
`CalculateMove` function will look something like this:

```Go
for _, adversary := range adversaryClient {
    moveResponse := adversary.CalculateMove(playerPositions)
    gameState.moveActorRelative(adversary.Name, Response.Move)
}
```

The main game loop will use the calculated relative move returned from the
`CalculateMove` function to move the player relatively. Since our `GameState`
already has this ability, and adversaries are guaranteed to not break the rules,
we do not need to rule check these moves. Should we be receiving user or external
input for adversary moves, adding a rule check would be fairly trivial and would
follow the pattern used for rule checking the players.
