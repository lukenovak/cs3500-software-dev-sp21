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
	Item        int
}


// generates a tile with no object at the given position
func GenerateTile(tileType int) *Tile {
	return &Tile{
		Type:        tileType,
		Item: 0,
	}
}

func (t Tile) Equals(secondTile Tile) bool {
	return t.Type == secondTile.Type && t.Item == secondTile.Item
}