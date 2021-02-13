package level

import (
	"fmt"
)

type Room [][]*Tile

// Given a starting coordinate and a room's dimensions, generate a rectangular room with the given doors
func GenerateRectangularRoom(topLeft Position2D, width int, length int, doors []Position2D) (Room, error) {
	var err error
	room := allocateRoom(width, length)
	if width < 3 || length < 3 {
		return nil, fmt.Errorf("invalid room dimensions")
	}
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			room[i][j], err = generateRoomTile(topLeft, width, length, NewPosition2D(topLeft.X + i, topLeft.Y + j), doors)
			if err != nil {
				return nil, err
			}
		}
	}
	return room, nil
}


// makes space for the room to be allocated
func allocateRoom(w int, l int) [][]*Tile {
	room := make([][]*Tile, w)
	for i := range room {
		room[i] = make([]*Tile, l)
	}
	return room
}

func generateRoomTile(topLeft Position2D, width int, length int, newTilePos Position2D, doors []Position2D) (*Tile, error) {
	if isPerimeter(topLeft, width, length, newTilePos) {
		if isDoor(newTilePos, doors) {
			return GenerateTile(door, newTilePos.X, newTilePos.Y), nil
		} else {
			return GenerateTile(wall, newTilePos.X, newTilePos.Y), nil
		}
	} else if isDoor(newTilePos, doors) {
		return nil, fmt.Errorf("invalid door at %d, %d", newTilePos.X, newTilePos.Y)
	} else {
		return GenerateTile(walkable, newTilePos.X, newTilePos.Y), nil
	}
}

// is this tile a perimeter tile?
func isPerimeter(topLeft Position2D, width int, length int, newTilePos Position2D) bool {
	return newTilePos.X == topLeft.X || newTilePos.Y == topLeft.Y ||
		newTilePos.X == topLeft.X + width - 1 || newTilePos.Y == topLeft.Y + length - 1
}

// is the given position included in the array of door positions?
func isDoor(tilePos Position2D, doors []Position2D) bool {
	isDoor := false
	for _, door := range doors {
		isDoor = isDoor || tilePos.Equals(door)
	}
	return isDoor
}