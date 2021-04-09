package level

import (
	"fmt"
)

const (
	vertical   = 0
	horizontal = 1
	up         = 2
	down       = 3
	right      = 4
	left       = 5
)

/* -------------------------------- Level Representation & struct methods -------------------------------- */

// Represents the level and all of its tiles.
type Level struct {
	Tiles         [][]*Tile            // The raw tile data as laid out on a 2d plane
	Size          Position2D           // The size of the room
	RoomDataGraph []RoomGraphNode      // A graph of RoomGraphNode, where index == room id. Useful for room metadata
	Items         map[*Item]Position2D // A map of Items in the level to thier absolute positions in the level
	IsUnlocked    bool                 // Keeps track of whether the doors on this level are ulocked
}

// Generates a level with a nil-initialized 2-d tile array of the given size
func NewEmptyLevel(rows int, columns int) (Level, error) {
	if rows < 1 || columns < 1 {
		return Level{Tiles: nil, Size: NewPosition2D(0, 0)}, fmt.Errorf("invalid level size")
	}
	return Level{
		Tiles:      allocateLevelTiles(rows, columns),
		Size:       NewPosition2D(rows, columns),
		Items:      map[*Item]Position2D{},
		IsUnlocked: false,
	}, nil
}

// expand the Level's 2d slice to match the new required position
func (level *Level) expandLevel(newSize Position2D) {
	for i := level.Size.Col; i < newSize.Col; i++ {
		for j := 0; j < level.Size.Row; j++ {
			level.Tiles[j] = append(level.Tiles[j], nil)
		}
	}
	for i := level.Size.Row; i < newSize.Row; i++ {
		level.Tiles = append(level.Tiles, make([]*Tile, max(newSize.Col, level.Size.Col)))
	}
	level.Size = getMaxPosition(level.Size, newSize)
}

// is this position within the bounds of the level?
func (level Level) isInboundsPosition(pos Position2D) bool {
	return 0 <= pos.Row && pos.Row < level.Size.Row && 0 <= pos.Col && pos.Col < level.Size.Col
}

// Alternate method of retrieving tiles using a Position2D. RETURNS NIL IF POS IS OUT OF BOUNDS!
func (level Level) GetTile(pos Position2D) *Tile {
	if level.isInboundsPosition(pos) {
		return level.Tiles[pos.Row][pos.Col]
	}
	return nil
}

// Gets the actual traversable Tiles that are numSteps number of steps away from the current position
func (level Level) GetWalkableTiles(pos Position2D, numSteps int) []*Tile {
	var walkableTiles []*Tile
	if numSteps > 0 {
		adjacentWalkablePositions := level.getAdjacentWalkablePositions(pos)
		for _, adjPosn := range adjacentWalkablePositions {
			nextStep := level.GetWalkableTiles(adjPosn, numSteps-1)
			for _, tile := range nextStep {
				walkableTiles = append(walkableTiles, tile)
			}
		}
		return walkableTiles
	} else {
		return []*Tile{level.GetTile(pos)}
	}
}

// gets traversable positions that are numSteps number of steps away from the current position
func (level Level) GetWalkableTilePositions(pos Position2D, numSteps int) []Position2D {
	walkablePosns := make(map[Position2D]bool)
	var walkablePosnArray []Position2D
	if numSteps > 0 {
		adjacentWalkablePositions := level.getAdjacentWalkablePositions(pos)
		for _, adjPosn := range adjacentWalkablePositions {
			nextStep := level.GetWalkableTilePositions(adjPosn, numSteps-1)
			for _, posn := range nextStep {
				walkablePosns[posn] = true
			}
			walkablePosns[adjPosn] = true
		}
		// convert map to array
		for pos := range walkablePosns {
			walkablePosnArray = append(walkablePosnArray, pos)
		}
		return walkablePosnArray
	} else {
		return []Position2D{pos}
	}
}

// gets traversable positions in the level that are adjacent to the given position
func (level Level) getAdjacentWalkablePositions(pos Position2D) []Position2D {
	var walkablePositions []Position2D

	if leftTile := level.GetTile(NewPosition2D(pos.Row-1, pos.Col)); leftTile != nil && (leftTile.Type == Walkable || leftTile.Type == Door) {
		walkablePositions = append(walkablePositions, NewPosition2D(pos.Row-1, pos.Col))
	}
	if rightTile := level.GetTile(NewPosition2D(pos.Row+1, pos.Col)); rightTile != nil && (rightTile.Type == Walkable || rightTile.Type == Door) {
		walkablePositions = append(walkablePositions, NewPosition2D(pos.Row+1, pos.Col))
	}
	if upTile := level.GetTile(NewPosition2D(pos.Row, pos.Col+1)); upTile != nil && (upTile.Type == Walkable || upTile.Type == Door) {
		walkablePositions = append(walkablePositions, NewPosition2D(pos.Row, pos.Col+1))
	}
	if downTile := level.GetTile(NewPosition2D(pos.Row, pos.Col-1)); downTile != nil && (downTile.Type == Walkable || downTile.Type == Door) {
		walkablePositions = append(walkablePositions, NewPosition2D(pos.Row, pos.Col-1))
	}
	return walkablePositions
}

func (level Level) GetTiles(origin Position2D, size Position2D) [][]*Tile {
	selection := allocateLevelTiles(size.Row, size.Col)

	for i := 0; i < size.Row; i++ {
		for j := 0; j < size.Col; j++ {
			selection[i][j] = level.Tiles[i+origin.Row][j+origin.Col]
		}
	}

	return selection
}

func (level Level) GetItems() ([]*Item, []Position2D) {
	items := make([]*Item, 0)
	positions := make([]Position2D, 0)

	for item, pos := range level.Items {
		items = append(items, item)
		positions = append(positions, pos)
	}

	return items, positions
}

/* -------------------------------- Room + Hallway Generation -------------------------------- */

// adds a Room's tiles to a Level, and expands the Level if necessary
func (level *Level) GenerateRectangularRoom(topLeft Position2D, rows int, cols int, doors []Position2D) error {
	newRoomId := len(level.RoomDataGraph)

	bottomRight := NewPosition2D(topLeft.Row+rows, topLeft.Col+cols)
	level.expandLevel(getMaxPosition(level.Size, bottomRight))
	if rows < 3 || cols < 3 {
		return fmt.Errorf("invalid room dimensions")
	}
	err := level.checkRoomValidity(topLeft, rows, cols)
	if err != nil {
		return err
	}

	// tile generation for the room
	for i := topLeft.Row; i < topLeft.Row+rows; i++ {
		for j := topLeft.Col; j < topLeft.Col+cols; j++ {
			level.Tiles[i][j], err = generateRoomTile(topLeft, rows, cols, NewPosition2D(i, j), doors, newRoomId)
			if err != nil {
				return err
			}
		}
	}

	// add the new room to the graph
	level.RoomDataGraph = append(level.RoomDataGraph, &RoomData{
		Id:      newRoomId,
		TopLeft: topLeft,
		Size:    NewPosition2D(rows, cols),
	})
	return nil
}

// given a top coordinate and a room layout, make the given rectangular room
func (level *Level) GenerateRectangularRoomWithLayout(topLeft Position2D, rows int, cols int, layout [][]int) error {
	newRoomId := len(level.RoomDataGraph)

	bottomRight := NewPosition2D(topLeft.Row+rows, topLeft.Col+cols)
	level.expandLevel(getMaxPosition(level.Size, bottomRight))
	if rows < 3 || cols < 3 || rows != len(layout) || cols != len(layout[0]) {
		return fmt.Errorf("invalid room layout")
	}
	err := level.checkRoomValidity(topLeft, rows, cols)
	if err != nil {
		return err
	}
	for row := 0; row < rows; row++ {
		for column := 0; column < cols; column++ {
			level.Tiles[topLeft.Row+row][topLeft.Col+column] = GenerateTile(layout[row][column], newRoomId)
		}
	}

	// add the new room to the graph
	level.RoomDataGraph = append(level.RoomDataGraph, &RoomData{
		Id:      newRoomId,
		TopLeft: topLeft,
		Size:    NewPosition2D(rows, cols),
	})

	return nil
}

// Checks to see that this room is valid (it does not overlap with another room)
func (level Level) checkRoomValidity(topLeft Position2D, rows int, cols int) error {
	for i := topLeft.Row; i < topLeft.Row+rows; i++ {
		for j := topLeft.Col; j < topLeft.Col+cols; j++ {
			if level.Tiles[i][j] != nil {
				return fmt.Errorf("invalid room placement. check that your room does not overlap with another room")
			}
		}
	}
	return nil
}

// generates a "hallway", which is a start and end point, with an ordered list of waypoints
func (level *Level) GenerateHallway(start Position2D, end Position2D, waypoints []Position2D) error {
	newHallwayId := len(level.RoomDataGraph)

	err := level.validateHallway(start, end, waypoints)
	if err != nil {
		return err
	}
	currPos := start

	level.expandLevel(getListMaxPosition(waypoints))

	// go through the waypoints and generate all the necessary tiles
	for _, waypoint := range waypoints {
		level.generateBetweenWaypoints(&currPos, waypoint, true, newHallwayId)
	}

	// the last "sub-hallway" to the end is special because it doesn't get capped
	if !currPos.Equals(end) {
		level.generateBetweenWaypoints(&currPos, end, false, newHallwayId)
	}
	level.GetTile(end).Type = Door

	// add the new room to the graph
	level.RoomDataGraph = append(level.RoomDataGraph, &HallData{
		Id:        newHallwayId,
		Start:     start,
		End:       end,
		Waypoints: waypoints,
	})

	// Add the room connections (we know this is a valid hallway so we gan get rooms directly)
	level.RoomDataGraph[newHallwayId].ConnectNode(level.RoomDataGraph[level.GetTile(start).RoomId])
	level.RoomDataGraph[newHallwayId].ConnectNode(level.RoomDataGraph[level.GetTile(end).RoomId])

	return nil
}

// generates the "hallway" segments between hallway waypoints
func (level Level) generateBetweenWaypoints(startPos *Position2D, endPos Position2D, shouldCapEnd bool, hallwayId int) {

	generateCol := func(amountToIterate int, orientation int, direction int) {
		for !startPos.Equals(endPos) {
			startPos.Col += amountToIterate
			level.generateHallwayStep(*startPos, orientation, hallwayId)
		}
		if shouldCapEnd {
			level.capHallwayEnd(*startPos, direction, hallwayId)
		}
	}

	generateRow := func(amountToIterate int, orientation int, direction int) {
		for !startPos.Equals(endPos) {
			startPos.Row += amountToIterate
			level.generateHallwayStep(*startPos, orientation, hallwayId)
		}
		if shouldCapEnd {
			level.capHallwayEnd(*startPos, direction, hallwayId)
		}
	}

	if endPos.Row == startPos.Row && endPos.Col > startPos.Col {
		generateCol(1, horizontal, right)
	} else if endPos.Row == startPos.Row && endPos.Col < startPos.Col {
		generateCol(-1, horizontal, left)
	} else if endPos.Row > startPos.Row {
		generateRow(1, vertical, down)
	} else {
		generateRow(-1, vertical, up)
	}
}

// Checks to see that the requested hallway is valid
func (level Level) validateHallway(start Position2D, end Position2D, waypoints []Position2D) error {
	err := validateWaypoints(start, end, waypoints)
	if err != nil {
		return err
	}

	if level.GetTile(start).Type != Door || level.GetTile(end).Type != Door {
		return fmt.Errorf("invalid hallway")
	}
	currPos := start

	validateHorizontal := func(waypoint Position2D, amountToIterate int) error {
		for !currPos.Equals(waypoint) {
			if err = level.validateHallwayStep(currPos, horizontal); err != nil {
				return err
			}
			currPos.Col += amountToIterate
		}
		if err = level.validateHallwayStep(NewPosition2D(currPos.Row, currPos.Col+amountToIterate), horizontal); err != nil {
			return err
		}
		return nil
	}

	validateVertical := func(waypoint Position2D, amountToIterate int) error {
		for !currPos.Equals(waypoint) {
			if err = level.validateHallwayStep(currPos, vertical); err != nil {
				return err
			}
			currPos.Row += amountToIterate
		}
		if err = level.validateHallwayStep(NewPosition2D(currPos.Row+amountToIterate, currPos.Col), vertical); err != nil {
			return err
		}
		return nil
	}

	// go through the waypoints and validate all the necessary tiles
	for _, waypoint := range append(waypoints) {
		if waypoint.Row == currPos.Row && waypoint.Col > currPos.Col { // moving right
			err = validateHorizontal(waypoint, 1)
		} else if waypoint.Row == currPos.Row && waypoint.Col < currPos.Col { // moving left
			err = validateHorizontal(waypoint, -1)
		} else if waypoint.Row > currPos.Row { // moving right
			err = validateVertical(waypoint, 1)
		} else if waypoint.Col < currPos.Col { // moving left
			err = validateVertical(waypoint, -1)
		}

		if err != nil {
			return err
		}
	}

	return err
}

// Checks to see that this "step" in the hallway is valid
func (level Level) validateHallwayStep(rowCenter Position2D, direction int) error {
	validate := func(doorTile *Tile, fstAdjWallTile *Tile, sndAdjWallTile *Tile) error {
		if (doorTile != nil && doorTile.Type != Door) ||
			(fstAdjWallTile != nil && fstAdjWallTile.Type != Wall) ||
			(sndAdjWallTile != nil && sndAdjWallTile.Type != Wall) {
			return fmt.Errorf("row is invalid")
		} else {
			return nil
		}
	}

	centerTile := level.GetTile(rowCenter)
	switch direction {
	case horizontal:
		return validate(centerTile,
			level.GetTile(NewPosition2D(rowCenter.Row-1, rowCenter.Col)),
			level.GetTile(NewPosition2D(rowCenter.Row+1, rowCenter.Col)))
	case vertical:
		return validate(centerTile,
			level.GetTile(NewPosition2D(rowCenter.Row, rowCenter.Col-1)),
			level.GetTile(NewPosition2D(rowCenter.Row, rowCenter.Col+1)))
	default:
		panic("invalid hallway direction")
	}
}

// generates the tiles in a single, 3-tile "step" of the hallway
func (level Level) generateHallwayStep(rowCenter Position2D, direction int, hallwayId int) {
	if level.GetTile(rowCenter) == nil || level.GetTile(rowCenter).Type == Wall {
		level.Tiles[rowCenter.Row][rowCenter.Col] = GenerateTile(Walkable, hallwayId)
	}
	switch direction {
	case horizontal:
		level.Tiles[rowCenter.Row-1][rowCenter.Col] = GenerateTile(Wall, hallwayId)
		level.Tiles[rowCenter.Row+1][rowCenter.Col] = GenerateTile(Wall, hallwayId)
	case vertical:
		level.Tiles[rowCenter.Row][rowCenter.Col-1] = GenerateTile(Wall, hallwayId)
		level.Tiles[rowCenter.Row][rowCenter.Col+1] = GenerateTile(Wall, hallwayId)
	default:
		panic("invalid hallway direction")
	}
}

// "Caps" the end of a hallway by adding a wall one tile set past the waypoint
func (level Level) capHallwayEnd(startPos Position2D, direction int, hallwayId int) {

	switch direction {
	case left:
		level.Tiles[startPos.Row+1][startPos.Col-1] = GenerateTile(Wall, hallwayId)
		level.Tiles[startPos.Row][startPos.Col-1] = GenerateTile(Wall, hallwayId)
		level.Tiles[startPos.Row-1][startPos.Col-1] = GenerateTile(Wall, hallwayId)
	case right:
		level.Tiles[startPos.Row+1][startPos.Col+1] = GenerateTile(Wall, hallwayId)
		level.Tiles[startPos.Row][startPos.Col+1] = GenerateTile(Wall, hallwayId)
		level.Tiles[startPos.Row-1][startPos.Col+1] = GenerateTile(Wall, hallwayId)
	case down:
		level.Tiles[startPos.Row+1][startPos.Col+1] = GenerateTile(Wall, hallwayId)
		level.Tiles[startPos.Row+1][startPos.Col] = GenerateTile(Wall, hallwayId)
		level.Tiles[startPos.Row+1][startPos.Col-1] = GenerateTile(Wall, hallwayId)
	case up:
		level.Tiles[startPos.Row-1][startPos.Col+1] = GenerateTile(Wall, hallwayId)
		level.Tiles[startPos.Row-1][startPos.Col] = GenerateTile(Wall, hallwayId)
		level.Tiles[startPos.Row-1][startPos.Col-1] = GenerateTile(Wall, hallwayId)
	default:
		panic("unknown hallway cap direction.")
	}
}

// Unlocks all exits in a level
func (level *Level) UnlockExits() {
	for _, itemPos := range level.Items {
		exitItem := Item{Type: UnlockedExit}
		tileItem := level.GetTile(itemPos).Item
		if tileItem != nil && tileItem.Type == LockedExit {
			level.ClearItem(itemPos)
			level.PlaceItem(itemPos, &exitItem)
		}
	}
	level.IsUnlocked = true
}

// Places an item on a tile if it does not currently have one
func (level Level) PlaceItem(pos Position2D, itemToPlace *Item) error {
	if itemTile := level.GetTile(pos); itemTile != nil && itemTile.Item == nil && itemTile.Type == Walkable {
		itemTile.Item = itemToPlace
		level.Items[itemToPlace] = pos
		return nil
	}
	return fmt.Errorf("invalid item placement")
}

// Clears a tile's item if it has one
func (level Level) ClearItem(pos Position2D) {
	if itemTile := level.GetTile(pos); itemTile != nil && itemTile.Item != nil {
		delete(level.Items, itemTile.Item)
		itemTile.Item = nil
	}
}

/* -------------------------------- Generation Utility Functions -------------------------------- */

// used in room generation to determine what kind of tile should be generated
func generateRoomTile(topLeft Position2D, row int, col int, newTilePos Position2D, doors []Position2D, roomId int) (*Tile, error) {
	if isPerimeter(topLeft, row, col, newTilePos) {
		if isDoor(newTilePos, doors) {
			return GenerateTile(Door, roomId), nil
		} else {
			return GenerateTile(Wall, roomId), nil
		}
	} else if isDoor(newTilePos, doors) {
		return nil, fmt.Errorf("invalid Door at %d, %d", newTilePos.Row, newTilePos.Col)
	} else {
		return GenerateTile(Walkable, roomId), nil
	}
}

// is this tile a perimeter tile of a room?
func isPerimeter(topLeft Position2D, width int, length int, newTilePos Position2D) bool {
	return newTilePos.Row == topLeft.Row || newTilePos.Col == topLeft.Col ||
		newTilePos.Row == topLeft.Row+width-1 || newTilePos.Col == topLeft.Col+length-1
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

// Allocates array space for all of the levels tiles. May produce a sparse array
func allocateLevelTiles(rows int, columns int) [][]*Tile {
	room := make([][]*Tile, rows)
	for i := range room {
		room[i] = make([]*Tile, columns)
	}
	return room
}

// ensures all hallway endpoints are valid (at right angles)
func validateWaypoints(start Position2D, end Position2D, waypoints []Position2D) error {
	invalidError := fmt.Errorf("waypoints are not all at right angles")
	if len(waypoints) == 0 && !(end.Row == start.Row || end.Col == start.Col) {
		return invalidError
	}
	for idx := range waypoints {
		if idx == 0 {
			if !(waypoints[idx].Row == start.Row || waypoints[idx].Col == start.Col) {
				return invalidError
			}
		} else if waypoints[idx].Row == waypoints[idx-1].Row || waypoints[idx].Col == waypoints[idx-1].Col {
			if !(waypoints[idx].Row == end.Row || waypoints[idx].Col == end.Col) {
				return invalidError
			}
		} else {
			return invalidError
		}
	}
	return nil
}
