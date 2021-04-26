package level

// Tile type constants
const (
	Wall = iota
	Walkable
	Door
)

// Item type constants
const (
	NoItem = iota
	KeyID
	LockedExit
	UnlockedExit
)

// Tile represents the smallest unit of a level. Each tile corresponds to one position and is said to be in the room
// with the given RoomId. It may contain an Item, but is not guaranteed to, so always check the Item field for a nil
// pointer
type Tile struct {
	Type   int
	RoomId int
	Item   *Item
}

// Item represents a game item, and has a Type. It may be expanded to contain more data in the future.
type Item struct {
	Type int
}

// GenerateTile generates a tile with no object at the given position
func GenerateTile(tileType int, roomId int) *Tile {
	return &Tile{
		Type:   tileType,
		RoomId: roomId,
		Item:   nil,
	}
}

// NewKey creates a new Item with the key ID type
func NewKey() Item {
	return Item{
		Type: KeyID,
	}
}

// Equals returns true if the two given tiles are the same
func (t Tile) Equals(secondTile Tile) bool {
	return t.Type == secondTile.Type &&
		!(t.Item == nil && secondTile.Item != nil) &&
		!(t.Item != nil && secondTile.Item == nil) &&
		(t.Item == nil && secondTile.Item == nil ||
			(t.Item.Type == secondTile.Item.Type))
}

// TypeAsString returns a string representation of this item's type
func (i Item) TypeAsString() string {
	switch i.Type {
	case KeyID:
		return "key"
	case UnlockedExit, LockedExit:
		return "exit"
	default:
		return "unknown-item"
	}
}

