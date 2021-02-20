package level

import (
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/item"
)

const (
	vertical = 0
	horizontal = 1
	up = 2
	down = 3
	right = 4
	left = 5
)

/* -------------------------------- Level Representation & struct methods -------------------------------- */

// Represents the level and all of its tiles.
type Level struct {
	Tiles [][]*Tile
	Exits []*Tile
	Size  Position2D
}

// Generates a level with a nil-initialized 2-d tile array of the given size
func NewEmptyLevel(width int, length int) (Level, error) {
	if width < 1 || length < 1 {
		return Level{Tiles: nil, Size: NewPosition2D(0, 0)}, fmt.Errorf("invalid level size")
	}
	return Level{
		Tiles: allocateLevelTiles(width, length),
		Size:  NewPosition2D(width, length),
	}, nil
}

// expand the Level's 2d slice to match the new required position
func (level *Level) expandLevel(newSize Position2D) {
	for i := level.Size.Y; i < newSize.Y; i++ {
		for j := 0; j < level.Size.X; j++ {
			level.Tiles[j] = append(level.Tiles[j], nil)
		}
	}
	for i := level.Size.X; i < newSize.X; i++ {
		level.Tiles = append(level.Tiles, make ([]*Tile, max(newSize.Y, level.Size.Y)))
	}
	level.Size = getMaxPosition(level.Size, newSize)
}

// is this position within the bounds of the level?
func (level Level) isInboundsPosition(pos Position2D) bool {
	return 0 < pos.X && pos.X < level.Size.X && 0 < pos.Y && pos.Y < level.Size.Y
}

// Alternate method of retrieving tiles using a Position2D. RETURNS NIL IF POS IS OUT OF BOUNDS!
func (level Level) getTile(pos Position2D) *Tile {
	if level.isInboundsPosition(pos) {
		return level.Tiles[pos.X][pos.Y]
	}
	return nil
}

func (level Level) GetWalkableTiles(pos Position2D, numSteps int) []*Tile {
	var walkableTiles []*Tile
	if numSteps > 0 {
		adjacentWalkablePositions := level.getAdjacentWalkablePositions(pos)
		for _, adjPosn := range adjacentWalkablePositions {
			nextStep := level.GetWalkableTiles(adjPosn, numSteps - 1)
			for _, tile := range nextStep {
				walkableTiles = append(walkableTiles, tile)
			}
		}
		return walkableTiles
	} else {
		return []*Tile{level.getTile(pos)}
	}
}

func (level Level) GetWalkableTilePositions(pos Position2D, numSteps int) []Position2D {
	var walkablePosns []Position2D
	if numSteps > 0 {
		adjacentWalkablePositions := level.getAdjacentWalkablePositions(pos)
		for _, adjPosn := range adjacentWalkablePositions {
			nextStep := level.GetWalkableTilePositions(adjPosn, numSteps - 1)
			for _, posn := range nextStep {
				walkablePosns = append(walkablePosns, posn)
			}
		}
		return walkablePosns
	} else {
		return []Position2D{pos}
	}
}

func (level Level) getAdjacentWalkablePositions(pos Position2D) []Position2D {
	var walkablePositions []Position2D
	if leftTile := level.getTile(NewPosition2D(pos.X - 1, pos.Y)); leftTile != nil && leftTile.Type == Walkable {
		walkablePositions = append(walkablePositions, NewPosition2D(pos.X + 1, pos.Y))
	}
	if rightTile := level.getTile(NewPosition2D(pos.X + 1, pos.Y)); rightTile != nil && rightTile.Type == Walkable {
		walkablePositions = append(walkablePositions, NewPosition2D(pos.X + 1, pos.Y))
	}
	if upTile := level.getTile(NewPosition2D(pos.X, pos.Y + 1)); upTile != nil && upTile.Type == Walkable {
		walkablePositions = append(walkablePositions, NewPosition2D(pos.X + 1, pos.Y))
	}
	if downTile := level.getTile(NewPosition2D(pos.X, pos.Y - 1)); downTile != nil && downTile.Type == Walkable {
		walkablePositions = append(walkablePositions, NewPosition2D(pos.X + 1, pos.Y))
	}
	return walkablePositions
}

/* -------------------------------- Room + Hallway Generation -------------------------------- */

// adds a Room's tiles to a Level, and expands the Level if necessary
func (level *Level) GenerateRectangularRoom(topLeft Position2D, width int, length int, doors []Position2D) error {
	bottomRight := getRoomBottomRight(topLeft, width, length)
	level.expandLevel(getMaxPosition(level.Size, bottomRight))
	if width < 3 || length < 3 {
		return fmt.Errorf("invalid room dimensions")
	}
	err := level.checkRoomValidity(topLeft, width, length)
	if err != nil {
		return err
	}
	for i := topLeft.X; i < topLeft.X + width; i++ {
		for j := topLeft.Y; j < topLeft.Y + length; j++ {
			level.Tiles[i][j], err = generateRoomTile(topLeft, width, length, NewPosition2D(i, j), doors)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Checks to see that this room is valid (it does not overlap with another room)
func (level Level) checkRoomValidity(topLeft Position2D, width int, length int) error {
	for i := topLeft.X; i < topLeft.X + width; i++ {
		for j := topLeft.Y; j < topLeft.Y + length; j++ {
			if level.Tiles[i][j] != nil {
				return fmt.Errorf("invalid room placement. check that your room does not overlap with another room")
			}
		}
	}
	return nil
}

// generates a "hallway", which is a start and end point, with an ordered list of waypoints
func (level Level) GenerateHallway(start Position2D, end Position2D, waypoints []Position2D) error {
	err := level.validateHallway(start, end, waypoints)
	if err != nil {
		return err
	}
	currPos := start

	level.expandLevel(getListMaxPosition(waypoints))

	// go through the waypoints and generate all the necessary tiles
	for _, waypoint := range waypoints {
		level.generateBetweenWaypoints(&currPos, waypoint, true)
	}

	// the last "sub-hallway" to the end is special because it doesn't get capped
	if !currPos.Equals(end) {
		level.generateBetweenWaypoints(&currPos, end, false)
	}
	level.getTile(end).Type = Door
	return nil
}

func (level Level) generateBetweenWaypoints(startPos *Position2D, endPos Position2D, shouldCapEnd bool) {
	if endPos.X == startPos.X && endPos.Y > startPos.Y { // moving down
		for !startPos.Equals(endPos) {
			startPos.Y += 1
			level.generateHallwayStep(*startPos, vertical)
		}
		if shouldCapEnd {
			level.capHallwayEnd(*startPos, down)
		}
	} else if endPos.X == startPos.X && endPos.Y < startPos.Y { // moving up
		for !startPos.Equals(endPos) {
			startPos.Y -= 1
			level.generateHallwayStep(*startPos, vertical)
		}
		if shouldCapEnd {
			level.capHallwayEnd(*startPos, up)
		}
	} else if endPos.X > startPos.X { // moving right
		for !startPos.Equals(endPos) {
			startPos.X += 1
			level.generateHallwayStep(*startPos, horizontal)
		}
		if shouldCapEnd {
			level.capHallwayEnd(*startPos, right)
		}
	} else if endPos.Y < startPos.Y { // moving left
		for !startPos.Equals(endPos) {
			startPos.X -= 1
			level.generateHallwayStep(*startPos, horizontal)
		}
		if shouldCapEnd {
			level.capHallwayEnd(*startPos, left)
		}
	}
}

// Checks to see that the requested hallway is valid
func (level Level) validateHallway(start Position2D, end Position2D, waypoints []Position2D) error {
	err := validateWaypoints(start, end, waypoints)
	if err != nil {
		return err
	}

	if level.getTile(start).Type != Door || level.getTile(end).Type != Door {
		return fmt.Errorf("invalid hallway")
	}
	currPos := start

	// go through the waypoints and validate all the necessary tiles
	for _, waypoint := range append(waypoints) {
		if waypoint.X == currPos.X && waypoint.Y > currPos.Y { // moving up
			for !currPos.Equals(waypoint) {
				if err = level.validateHallwayStep(currPos, vertical); err != nil {
					return err
				}
				currPos.Y += 1
			}
			if err = level.validateHallwayStep(NewPosition2D(currPos.X, currPos.Y+1), vertical); err != nil {
				return err
			}
		} else if waypoint.X == currPos.X && waypoint.Y < currPos.Y { // moving down
			for !currPos.Equals(waypoint) {
				if err = level.validateHallwayStep(currPos, vertical); err != nil {
					return err
				}
				currPos.Y -= 1
			}
			if err = level.validateHallwayStep(NewPosition2D(currPos.X, currPos.Y-1), vertical); err != nil {
				return err
			}
		} else if waypoint.X > currPos.X { // moving right
			for !currPos.Equals(waypoint) {
				if err = level.validateHallwayStep(currPos, horizontal); err != nil {
					return err
				}
				currPos.X += 1
			}
			if err = level.validateHallwayStep(NewPosition2D(currPos.X+1, currPos.Y), horizontal); err != nil {
				return err
			}
		} else if waypoint.Y < currPos.Y { // moving left
			for !currPos.Equals(waypoint) {
				if err = level.validateHallwayStep(currPos, horizontal); err != nil {
					return err
				}
				currPos.X -= 1
			}
			if err = level.validateHallwayStep(NewPosition2D(currPos.X-1, currPos.Y), horizontal); err != nil {
				return err
			}
		}
	}

	return nil
}

// Checks to see that this "step" in the hallway is valid
func (level Level) validateHallwayStep(rowCenter Position2D, direction int) error {
	centerTile := level.getTile(rowCenter)
	switch direction {
	case vertical:
		leftTile := level.getTile(NewPosition2D(rowCenter.X - 1, rowCenter.Y))
		rightTile := level.getTile(NewPosition2D(rowCenter.X + 1, rowCenter.Y))
		if (centerTile != nil && centerTile.Type != Door) ||
			(leftTile != nil && leftTile.Type != Wall) ||
			(rightTile != nil && rightTile.Type != Wall) {
			return fmt.Errorf("row is invalid")
		}
	case horizontal:
		topTile := level.getTile(NewPosition2D(rowCenter.X, rowCenter.Y - 1))
		bottomTile := level.getTile(NewPosition2D(rowCenter.X, rowCenter.Y + 1))
		if (centerTile != nil && centerTile.Type != Door) ||
			(topTile != nil && topTile.Type != Wall) ||
			(bottomTile != nil && bottomTile.Type != Wall) {
			return fmt.Errorf("row is invalid")
		}
	default:
		panic("invalid hallway direction")
	}
	return nil
}

// generates the tiles in a single, 3-tile "step" of the hallway
func (level Level) generateHallwayStep(rowCenter Position2D, direction int) {
	if level.getTile(rowCenter) == nil || level.getTile(rowCenter).Type == Wall {
		level.Tiles[rowCenter.X][rowCenter.Y] = GenerateTile(Walkable)
	}
	switch direction {
	case vertical:
		level.Tiles[rowCenter.X - 1][rowCenter.Y] = GenerateTile(Wall)
		level.Tiles[rowCenter.X + 1][rowCenter.Y] = GenerateTile(Wall)
	case horizontal:
		level.Tiles[rowCenter.X][rowCenter.Y - 1] = GenerateTile(Wall)
		level.Tiles[rowCenter.X][rowCenter.Y + 1] = GenerateTile(Wall)
	default:
		panic("invalid hallway direction")
	}
}

// "Caps" the end of a hallway by adding a wall one tile set past the waypoint
func (level Level) capHallwayEnd(startPos Position2D, direction int) {
	switch direction {
	case up:
		level.Tiles[startPos.X + 1][startPos.Y - 1] = GenerateTile(Wall)
		level.Tiles[startPos.X][startPos.Y - 1] = GenerateTile(Wall)
		level.Tiles[startPos.X - 1][startPos.Y - 1] = GenerateTile(Wall)
	case down:
		level.Tiles[startPos.X + 1][startPos.Y + 1] = GenerateTile(Wall)
		level.Tiles[startPos.X][startPos.Y + 1] = GenerateTile(Wall)
		level.Tiles[startPos.X - 1][startPos.Y + 1] = GenerateTile(Wall)
	case right:
		level.Tiles[startPos.X + 1][startPos.Y + 1] = GenerateTile(Wall)
		level.Tiles[startPos.X + 1][startPos.Y] = GenerateTile(Wall)
		level.Tiles[startPos.X + 1][startPos.Y - 1] = GenerateTile(Wall)
	case left:
		level.Tiles[startPos.X - 1][startPos.Y + 1] = GenerateTile(Wall)
		level.Tiles[startPos.X - 1][startPos.Y] = GenerateTile(Wall)
		level.Tiles[startPos.X - 1][startPos.Y - 1] = GenerateTile(Wall)
	default:
		panic("unknown hallway cap direction.")
	}
}

// places an exit on a valid, walkable tile. Else, throws an error
func (level *Level) PlaceExit(exitPos Position2D) error {
	if exitTile := level.getTile(exitPos); exitTile != nil && exitTile.Type == Walkable {
		level.Tiles[exitPos.X][exitPos.Y].Type = LockedExit
		level.Exits = append(level.Exits, exitTile)
	} else {
		return fmt.Errorf("invalid exit location")
	}
	return nil
}

// Places an item on a tile if it does not currently have one
func (level Level) PlaceItem(pos Position2D, itemId int) error {
	if itemTile := level.getTile(pos); itemTile != nil && itemTile.Item == item.NoItem {
		itemTile.Item = itemId
		return nil
	}
	return fmt.Errorf("invalid item placement")
}

// Clears a tile's item if it has one
func (level Level) ClearItem(pos Position2D) {
	if itemTile := level.getTile(pos); itemTile != nil {
		itemTile.Item = item.NoItem
	}
}
/* -------------------------------- Generation Utility Functions -------------------------------- */

// used in room generation to determine what kind of tile should be generated
func generateRoomTile(topLeft Position2D, width int, length int, newTilePos Position2D, doors []Position2D) (*Tile, error) {
	if isPerimeter(topLeft, width, length, newTilePos) {
		if isDoor(newTilePos, doors) {
			return GenerateTile(Door), nil
		} else {
			return GenerateTile(Wall), nil
		}
	} else if isDoor(newTilePos, doors) {
		return nil, fmt.Errorf("invalid Door at %d, %d", newTilePos.X, newTilePos.Y)
	} else {
		return GenerateTile(Walkable), nil
	}
}

// is this tile a perimeter tile of a room?
func isPerimeter(topLeft Position2D, width int, length int, newTilePos Position2D) bool {
	return newTilePos.X == topLeft.X || newTilePos.Y == topLeft.Y ||
		newTilePos.X == topLeft.X + width - 1 || newTilePos.Y == topLeft.Y + length - 1
}

// is the given position included in the array of Door positions?
func isDoor(tilePos Position2D, doors []Position2D) bool {
	for _, door := range doors {
		if tilePos.Equals(door) {
			return true
		}
	}
	return false
}

func getRoomBottomRight(topLeft Position2D, width int, length int) Position2D {
	return NewPosition2D(topLeft.X + width, topLeft.Y + length)
}

func allocateLevelTiles(w int, l int) [][]*Tile {
	room := make([][]*Tile, w)
	for i := range room {
		room[i] = make([]*Tile, l)
	}
	return room
}

// ensures all hallway endpoints are valid (at right angles)
func validateWaypoints(start Position2D, end Position2D, waypoints []Position2D) error {
	invalidError := fmt.Errorf("waypoints are not all at right angles")
	if len(waypoints) == 0 && !(end.X == start.X || end.Y == start.Y) {
		return invalidError
	}
	for idx := range waypoints {
		if idx == 0 {
			if !(waypoints[idx].X == start.X || waypoints[idx].Y == start.Y) {
				return invalidError
			}
		} else if waypoints[idx].X == waypoints[idx - 1].X || waypoints[idx].Y == waypoints[idx - 1].Y {
			if !(waypoints[idx].X == end.X || waypoints[idx].Y == end.Y) {
				return invalidError
			}
		} else {
			return invalidError
		}
	}
	return nil
}