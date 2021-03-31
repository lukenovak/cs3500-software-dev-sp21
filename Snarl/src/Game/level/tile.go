package level

// Tile type constants
const (
	Wall     = iota
	Walkable = iota
	Door     = iota
)

// Item type constants
const (
	NoItem       = iota
	KeyID        = iota
	LockedExit   = iota
	UnlockedExit = iota
)

type Tile struct {
	Type   int
	RoomId int
	Item   *Item
}
type Item struct {
	Type int
}

// generates a tile with no object at the given position
func GenerateTile(tileType int, roomId int) *Tile {
	return &Tile{
		Type:   tileType,
		RoomId: roomId,
		Item:   nil,
	}
}

func NewKey() Item {
	return Item{
		Type: KeyID,
	}
}

func (t Tile) Equals(secondTile Tile) bool {
	return t.Type == secondTile.Type &&
		!(t.Item == nil && secondTile.Item != nil) &&
		!(t.Item != nil && secondTile.Item == nil) &&
		(t.Item == nil && secondTile.Item == nil ||
			(t.Item.Type == secondTile.Item.Type))
}
