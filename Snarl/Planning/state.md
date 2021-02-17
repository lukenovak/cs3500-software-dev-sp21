# MEMORANDUM

#### DATE: February 14, 2021
#### TO: Growl, Inc.
#### FROM: William Victoria and Luke Novak, Software Developers, Ormegland Inc.
#### SUBJECT: Snarl Data Representations

---

This memo presents our data representations that will be used in our implementation 
of Snarl. We used the provided Snarl documentation including the overview, plan, 
and protocol documents to determine our data representations. 

### Level Data

Our Game Manager stores the level data as a `Level` struct. This struct has a 2d array of tile 
pointers, a list of the exit tiles and a size. Tiles store their type and possibly an item id. 
This array is generated using several functions that generate and validate rooms and hallways.

```go
type Level struct {
	Tiles [][]*Tile
	Exits []*Tile
	Size Position2D
}

func (level *Level) GenerateRectangularRoom(topLeft Position2D, width int, length int, doors []Position2D) error { ... }

func (level Level) checkRoomValidity(topLeft Position2D, width int, length int) error { ... }

func (level Level) GenerateHallway(start Position2D, end Position2D, waypoints []Position2D) error { ... }

func (level Level) validateHallway(start Position2D, end Position2D, waypoints []Position2D) error { ... }
```

### Player Data

Our Game Manager stores player data as an ordered list of players. Each player has a position, 
an input source, either STDIN or TCP, and a type, either human or adversary. Our player data 
representation can also be easily modified to include attributes like attack points or health 
points or even and inventory system.

During the gameplay loop our Game Manager will need to generate a partial view for each player 
and transmit that view to them and then receive their action for that turn. 

```go
func (level Level) GeneratePlayerView(playerPos Position2D, viewDistance int) ([][]*Tile, error) { ... }

func MovePlayer(playerPos Position2D, action int) (Position2D, error) { ... }
```

---

The specified data representations allow for the required functionality to run a game of Snarl.
