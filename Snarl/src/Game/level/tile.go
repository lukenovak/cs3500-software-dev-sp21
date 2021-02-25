package level

import "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/item"

const (
	Wall         = 0
	Walkable     = 1
	Door         = 2
	LockedExit   = 3
	UnlockedExit = 4
)

type Tile struct {
	Type        int
	Item        *item.Item
}


// generates a tile with no object at the given position
func GenerateTile(tileType int) *Tile {
	return &Tile{
		Type: tileType,
		Item: nil,
	}
}

func (t Tile) Equals(secondTile Tile) bool {
	return t.Type == secondTile.Type && t.Item.Type == secondTile.Item.Type
}