package level

import "fmt"

const (
	vertical = 0
	horizontal = 1
)

type Level struct {
	Tiles [][]*Tile
	Size Position2D
}

func NewEmptyLevel(width int, length int) Level {
	return Level {
		Tiles: allocateLevelTiles(width, length),
		Size: NewPosition2D(width, length),
	}
}

func (l Level) getTile(pos Position2D) *Tile {
	return l.Tiles[pos.X][pos.Y]
}

// adds a Room's tiles to a Level, and expands the Level if necessary
func (l Level) GenerateRectangularRoom(topLeft Position2D, width int, length int, doors []Position2D) error {
	bottomRight := getRoomBottomRight(topLeft, width, length)
	expandLevel(&l, getMaxPosition(l.Size, bottomRight))
	if width < 3 || length < 3 {
		return fmt.Errorf("invalid room dimensions")
	}
	err := l.checkRoomValidity(topLeft, width, length)
	if err != nil {
		return err
	}
	for i := topLeft.X; i < topLeft.X + width; i++ {
		for j := topLeft.Y; j < topLeft.Y + length; j++ {
			l.Tiles[i][j], err = generateRoomTile(topLeft, width, length, NewPosition2D(i, j), doors)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// generates a "hallway", which is a start and end point, with an ordered list of waypoints
func (l Level) GenerateHallway(start Position2D, end Position2D, waypoints []Position2D) error {
	err := l.validateHallway(start, end, waypoints)
	if err != nil {
		return err
	}
	currPos := start

	// go through the waypoints and generate all the necessary tiles
	for _, waypoint := range waypoints {
		l.generateBetweenWaypoints(&currPos, waypoint, true)
	}

	if !currPos.Equals(end) {
		l.generateBetweenWaypoints(&currPos, end, false)
	}
	l.Tiles[end.X][end.Y] = GenerateTile(Door, end.X, end.Y)
	return nil
}

func (l Level) generateBetweenWaypoints(startPos *Position2D, endPos Position2D, shouldCapEnd bool) {
	if endPos.X == startPos.X && endPos.Y > startPos.Y { // moving down
		for !startPos.Equals(endPos) {
			startPos.Y += 1
			l.generateHallwayRow(*startPos, vertical)
		}
		if shouldCapEnd {
			l.Tiles[startPos.X + 1][startPos.Y + 1] = GenerateTile(Wall, startPos.X + 1, startPos.Y + 1)
			l.Tiles[startPos.X][startPos.Y + 1] = GenerateTile(Wall, startPos.X, startPos.Y + 1)
			l.Tiles[startPos.X - 1][startPos.Y + 1] = GenerateTile(Wall, startPos.X - 1, startPos.Y + 1)
		}
	} else if endPos.X == startPos.X && endPos.Y < startPos.Y { // moving up
		for !startPos.Equals(endPos) {
			startPos.Y -= 1
			l.generateHallwayRow(*startPos, vertical)
		}
		if shouldCapEnd {
			l.Tiles[startPos.X + 1][startPos.Y - 1] = GenerateTile(Wall, startPos.X + 1, startPos.Y - 1)
			l.Tiles[startPos.X][startPos.Y - 1] = GenerateTile(Wall, startPos.X, startPos.Y - 1)
			l.Tiles[startPos.X - 1][startPos.Y - 1] = GenerateTile(Wall, startPos.X - 1, startPos.Y - 1)
		}
	} else if endPos.X > startPos.X { // moving right
		for !startPos.Equals(endPos) {
			startPos.X += 1
			l.generateHallwayRow(*startPos, horizontal)
		}
		if shouldCapEnd {
			l.Tiles[startPos.X + 1][startPos.Y + 1] = GenerateTile(Wall, startPos.X + 1, startPos.Y + 1)
			l.Tiles[startPos.X + 1][startPos.Y] = GenerateTile(Wall, startPos.X + 1, startPos.Y)
			l.Tiles[startPos.X + 1][startPos.Y - 1] = GenerateTile(Wall, startPos.X + 1, startPos.Y - 1)
		}
	} else if endPos.Y < startPos.Y { // moving left
		for !startPos.Equals(endPos) {
			startPos.X -= 1
			l.generateHallwayRow(*startPos, horizontal)
		}
		if shouldCapEnd {
			l.Tiles[startPos.X - 1][startPos.Y + 1] = GenerateTile(Wall, startPos.X - 1, startPos.Y + 1)
			l.Tiles[startPos.X - 1][startPos.Y] = GenerateTile(Wall, startPos.X - 1, startPos.Y)
			l.Tiles[startPos.X - 1][startPos.Y - 1] = GenerateTile(Wall, startPos.X - 1, startPos.Y - 1)
		}
	}
}

func (l Level) checkRoomValidity(topLeft Position2D, width int, length int) error {
	for i := topLeft.X; i < topLeft.X + width; i++ {
		for j := topLeft.Y; j < topLeft.Y + length; j++ {
			if l.Tiles[i][j] != nil {
				return fmt.Errorf("invalid room placement. check that your room does not overlap with another room")
			}
		}
	}
	return nil
}

func (l Level) validateHallway(start Position2D, end Position2D, waypoints []Position2D) error {
	err := validateWaypoints(start, end, waypoints)
	if err != nil {
		return err
	}

	currPos := start

	// go through the waypoints and validate all the necessary tiles
	for _, waypoint := range append(waypoints) {
		if waypoint.X == currPos.X && waypoint.Y > currPos.Y { // moving up
			for !currPos.Equals(waypoint) {
				if err = l.validateHallwayRow(currPos, vertical); err != nil {
					return err
				}
				currPos.Y += 1
			}
			if err = l.validateHallwayRow(NewPosition2D(currPos.X, currPos.Y+1), vertical); err != nil {
				return err
			}
		} else if waypoint.X == currPos.X && waypoint.Y < currPos.Y { // moving down
			for !currPos.Equals(waypoint) {
				if err = l.validateHallwayRow(currPos, vertical); err != nil {
					return err
				}
				currPos.Y -= 1
			}
			if err = l.validateHallwayRow(NewPosition2D(currPos.X, currPos.Y-1), vertical); err != nil {
				return err
			}
		} else if waypoint.X > currPos.X { // moving right
			for !currPos.Equals(waypoint) {
				if err = l.validateHallwayRow(currPos, horizontal); err != nil {
					return err
				}
				currPos.X += 1
			}
			if err = l.validateHallwayRow(NewPosition2D(currPos.X+1, currPos.Y), horizontal); err != nil {
				return err
			}
		} else if waypoint.Y < currPos.Y { // moving left
			for !currPos.Equals(waypoint) {
				if err = l.validateHallwayRow(currPos, horizontal); err != nil {
					return err
				}
				currPos.X -= 1
			}
			if err = l.validateHallwayRow(NewPosition2D(currPos.X-1, currPos.Y), horizontal); err != nil {
				return err
			}
		}
	}

	return nil
}

func (l Level) validateHallwayRow(rowCenter Position2D, direction int) error {
	centerTile := l.getTile(rowCenter)
	switch direction {
	case vertical:
		leftTile := l.getTile(NewPosition2D(rowCenter.X - 1, rowCenter.Y))
		rightTile := l.getTile(NewPosition2D(rowCenter.X + 1, rowCenter.Y))
		if (centerTile != nil && centerTile.Type != Door) ||
			(leftTile != nil && leftTile.Type != Wall) ||
			(rightTile != nil && rightTile.Type != Wall) {
			return fmt.Errorf("row is invalid")
		}
	case horizontal:
		topTile := l.getTile(NewPosition2D(rowCenter.X, rowCenter.Y - 1))
		bottomTile := l.getTile(NewPosition2D(rowCenter.X, rowCenter.Y + 1))
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

func (l Level) generateHallwayRow(rowCenter Position2D, direction int) {
	if l.getTile(rowCenter) == nil || l.getTile(rowCenter).Type == Wall {
		l.Tiles[rowCenter.X][rowCenter.Y] = GenerateTile(Walkable, rowCenter.X, rowCenter.Y)
	}
	switch direction {
	case vertical:
		l.Tiles[rowCenter.X - 1][rowCenter.Y] = GenerateTile(Wall, rowCenter.X - 1, rowCenter.Y)
		l.Tiles[rowCenter.X + 1][rowCenter.Y] = GenerateTile(Wall, rowCenter.X + 1, rowCenter.Y)
	case horizontal:
		l.Tiles[rowCenter.X][rowCenter.Y - 1] = GenerateTile(Wall, rowCenter.X, rowCenter.Y - 1)
		l.Tiles[rowCenter.X][rowCenter.Y + 1] = GenerateTile(Wall, rowCenter.X, rowCenter.Y + 1)
	default:
		panic("invalid hallway direction")
	}
}

// used in room generation to determine what kind of tile should be generated
func generateRoomTile(topLeft Position2D, width int, length int, newTilePos Position2D, doors []Position2D) (*Tile, error) {
	if isPerimeter(topLeft, width, length, newTilePos) {
		if isDoor(newTilePos, doors) {
			return GenerateTile(Door, newTilePos.X, newTilePos.Y), nil
		} else {
			return GenerateTile(Wall, newTilePos.X, newTilePos.Y), nil
		}
	} else if isDoor(newTilePos, doors) {
		return nil, fmt.Errorf("invalid Door at %d, %d", newTilePos.X, newTilePos.Y)
	} else {
		return GenerateTile(Walkable, newTilePos.X, newTilePos.Y), nil
	}
}

// is this tile a perimeter tile of a room?
func isPerimeter(topLeft Position2D, width int, length int, newTilePos Position2D) bool {
	return newTilePos.X == topLeft.X || newTilePos.Y == topLeft.Y ||
		newTilePos.X == topLeft.X + width - 1 || newTilePos.Y == topLeft.Y + length - 1
}

// is the given position included in the array of Door positions?
func isDoor(tilePos Position2D, doors []Position2D) bool {
	isDoor := false
	for _, door := range doors {
		if tilePos.Equals(door) {
			return true
		}
	}
	return isDoor
}

// expand the Level's 2d slice to match the new required position
func expandLevel(level *Level, newSize Position2D) {
	for i := level.Size.Y; i < newSize.Y; i++ {
		for j := 0; j < level.Size.X; j++ {
			level.Tiles[j] = append(level.Tiles[j], nil)
		}
	}
	for i := level.Size.X; i < newSize.X; i++ {
		level.Tiles = append(level.Tiles, make([]*Tile, level.Size.Y))
	}
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