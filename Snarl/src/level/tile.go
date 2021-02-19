package level

const (
	Wall         = 0
	Walkable     = 1
	Door         = 2
	LockedExit   = 3
	UnlockedExit = 4
)

type Tile struct {
	Type        int
	Coordinates Position2D
	Item        int
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
		Item: 0,
	}
}

func (t Tile) Equals(secondTile Tile) bool {
	return t.IsPosition(secondTile.Coordinates) && t.Type == secondTile.Type
}