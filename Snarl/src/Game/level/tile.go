package level

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/item"
)

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
	Item   *item.Item
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
	return t.Type == secondTile.Type && t.Item.Type == secondTile.Item.Type
}

// generates data ablut the tile for use in a testing task
func (t Tile) TileData() map[string]interface{} {
	var tileData map[string]interface{}

	switch t.Type {
	case Wall:
		tileData["traversable"] = false
	case LockedExit, UnlockedExit:
		tileData["object"] = "exit"
		tileData["traversable"] = true
	case Walkable, Door:
		tileData["traversable"] = true
	}

	switch t.Item.Type {
	case item.NoItem:
		tileData["object"] = nil
	case item.KeyID:
		tileData["object"] = "key"
	}

	return tileData
}
