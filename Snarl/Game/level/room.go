package level

import (
	"fmt"
)

type Room struct {
	TopLeft Position2D
	Tiles [][]*Tile
}

// Given a starting coordinate and a room's dimensions, generate a rectangular room with the given doors
func GenerateRectangularRoom(topLeft Position2D, width int, length int, doors []Position2D) (Room, error) {
	var err error
	roomTiles := allocateRoomTiles(width, length)
	if width < 3 || length < 3 {
		return Room{originPosition, nil}, fmt.Errorf("invalid room dimensions")
	}
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			roomTiles[i][j], err = generateRoomTile(topLeft, width, length, NewPosition2D(topLeft.X + i, topLeft.Y + j), doors)
			if err != nil {
				return Room{originPosition, nil}, err
			}
		}
	}
	return Room{
		TopLeft: topLeft,
		Tiles: roomTiles,
	}, nil
}

// generates a room of width 1 that follows the path from start to end.
func GenerateHallway(start Position2D, end Position2D, waypoints []Position2D) (Room, error) {
	return Room{originPosition, nil}, nil
}

// makes space for the room to be allocated
func allocateRoomTiles(w int, l int) [][]*Tile {
	room := make([][]*Tile, w)
	for i := range room {
		room[i] = make([]*Tile, l)
	}
	return room
}

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

// is this tile a perimeter tile?
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