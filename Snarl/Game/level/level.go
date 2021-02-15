package level

import "fmt"

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

// adds a Room's tiles to a Level, and expands the Level if necessary
func (l Level) GenerateRectangularRoom(topLeft Position2D, width int, length int, doors []Position2D) error {
	var err error
	if width < 3 || length < 3 {
		return fmt.Errorf("invalid room dimensions")
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
		isDoor = isDoor || tilePos.Equals(door)
	}
	return isDoor
}

// expand the Level's 2d slice to match the new required position
func expandLevel(level *Level, newSize Position2D) {
	panic("not yet implemented")
}

func allocateLevelTiles(w int, l int) [][]*Tile {
	room := make([][]*Tile, w)
	for i := range room {
		room[i] = make([]*Tile, l)
	}
	return room
}