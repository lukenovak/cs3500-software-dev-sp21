package level

const (
	Wall         = 0
	Walkable     = 1
	Door         = 2
	lockedExit   = 3
	unlockedExit = 4
)

const (
	noObject = 0
)

type Tile struct {
	Type int
	Coordinates Position2D
	Object int
}

func (t Tile) IsPosition(d Position2D) bool {
	return t.Coordinates.Equals(d)
}

// generates a tile with no object at the given position
func GenerateTile(tileType int, xPosn int, yPosn int) *Tile {
	return &Tile{
		Type:        tileType,
		Coordinates: Position2D{
			X: xPosn,
			Y: yPosn,
		},
		Object:      0,
	}
}

func (t Tile) Equals(secondTile Tile) bool {
	return t.IsPosition(secondTile.Coordinates) && t.Type == secondTile.Type
}