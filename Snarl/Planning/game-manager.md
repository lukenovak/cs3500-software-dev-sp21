# MEMORANDUM

#### DATE: March 3, 2021
#### TO: Growl, Inc.
#### FROM: William Victoria and Luke Novak, Software Developers, Ormegland Inc.
#### SUBJECT: Snarl Game Manager

---

The memo presents our plan for the Snarl Game Manager. It includes a description of the expected functionality of our Game Manager and some functions that we expect to implement as a part of our Game Manager.

What does the Game Manager need to do:
1. Initialize Player components
2. Produce partial game states for all clients
3. Send those partial game states to their corresponding clients
4. Wait for and then handle responses from the clients
5. Rule check responses, if invalid, GOTO 4
6. Update the game state, including position and interactions
7. Check if the players have won or lost. If so, end the game
8. GOTO 2

Our Game Manager is not an interface or class onto itself. It is simply a public function to start and handle the main game loop. That game loop will delegate to other components to make changes to game state, or send and receive data from clients. The definition for the client response can be found in `player.md`.

```Go
// The main game management function. Returns an error if something goes wrong
func ManageGame(level *level.Level, users []UserClient) error { 
    
    // End the game if we have a bad number of users
    if len(users) < 0 || len(users) > maxUsers {
        return nil
    }

    // Create Player Actors for all human Players
    var players []Actor
    for _, user := range users { ... }

    // Generate Adversaries
    adversaries := generateAdversaries(level)

    currentState := initGameState(level, players, adversaries)

    // start game, and run until it ends
    for !isGameEnd(currentState) {
        
        // TODO: Multithread this so users can put moves in while others are moving???
        for _, user := range users {
            // Generate partial state
            // Send partial state to user
            err =: user.SendPartialState(generatePlayerVisibleState(level, player))
            
            // Get response
            response := user.GetInput()
            
            // build updated game state
            newState := buildNewStateFromResponse(currentState, response)

            // rule check the state
            if IsValidMove(currentState, newState) {
                // update for valid moves
                currentState = newState
            }
        }
    }
}
```

These are some private functions that will perform actions during setup and during the game loop for the Game Manager.

```Go
/* -------------- PRIVATE FUNCTIONS ----------------- */

// Generates n adversaries for the level
func generateAdversaries(gameLevel *level.Level) []Actor { ... }

// Builds the new game state from the response
func buildNewStateFromResponse(currentState GameState, response ClientResponse) { ... }
```
