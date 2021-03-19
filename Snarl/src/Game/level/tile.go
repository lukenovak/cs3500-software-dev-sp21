package level

const (
	Wall         = 0
	Walkable     = 1
	Door         = 2
	LockedExit   = 3
	UnlockedExit = 4
)

type Tile struct {
	Type   int
	RoomId int
	Item   *Item
}

// generates a tile with no object at the given position
func GenerateTile(tileType int, roomId int) *Tile {
	return &Tile{
		Type:   tileType,
		RoomId: roomId,
		Item:   nil,
	}
}

func (t Tile) Equals(secondTile Tile) bool {
	return t.Type == secondTile.Type &&
		!(t.Item == nil && secondTile.Item != nil) &&
		!(t.Item != nil && secondTile.Item == nil) &&
		(t.Item == nil && secondTile.Item == nil ||
			(t.Item.Type == secondTile.Item.Type))
}
