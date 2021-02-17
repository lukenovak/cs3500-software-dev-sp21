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
pointers, a list of the exit tiles and a size. Tiles store their type and an item id. 
The level layout is generated using several functions that generate and validate rooms and
hallways.

```go
type Tile struct {
	Type int
	Item int
}

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

### Actor Data

Our Game Manager stores actor data as an ordered list of `Actors`. Each actor has a position, 
an input source, either STDIN or TCP, and a type, either player or adversary. Our actor data 
representation can also be easily modified to include attributes like attack points or health 
points or even and inventory system.

During the gameplay loop our Game Manager will need to generate a partial view for each actor 
and transmit that view to them and then receive their action for that turn. 

```go
type Actor struct {
  Type int
  Position Position2D
  Input Reader
  Output Writer
}

func (level Level) GenerateActorView(position Position2D, viewDistance int) ([][]*Tile, error) { ... }

func (actor Actor) MoveActor(action int) error { ... }
```

### Game State

We store all the data the Game Manager requires in the `GameState` struct. It contains the 
level data and the actor data. All level data including room layouts, door and exit locations, 
and item locations is all stored in the `Level` struct. All actor data including actor 
position, move order, input and output sources, is all stored in the ordered lists of `Actor`s.

We check the victory condition for a given level at the end of each game loop by checking to see if any player have exited the level.

```go
type GameState struct {
	Level Level
  Players []Actor
  Adversaries []Actor
}

func (gs GameState) GenerateLevel(size Position2D) error { ... }

func (gs GameState) SpawnActor(position Position2D, actor Actor) error { ... }

func (gs GameState) CheckVictory() bool { ... }
```

---

The specified data representations allow for the required functionality to run a game of Snarl.
