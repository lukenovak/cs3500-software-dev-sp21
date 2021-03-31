package render

import "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"

const (
	wallTile     = "▓"
	walkableTile = "░"
	doorTile     = "D"
	lockedTile   = "¤"
	unlockedTile = "U"
	unknownTile  = "?"
)

func ASCIILevel(levelToRender level.Level) string {
	render := ""
	for i := 0; i < len(levelToRender.Tiles[0]); i++ {
		for j := range levelToRender.Tiles {
			render = render + renderTile(levelToRender.Tiles[j][i])
		}
		render = render + "\n"
	}
	return render
}

func renderTile(tile *level.Tile) string {
	if tile == nil {
		return " "
	}
	switch tile.Type {
	case level.Wall:
		return wallTile
	case level.Walkable:
		return walkableTile
	case level.Door:
		return doorTile
	default:
		switch tile.Item.Type {
		case level.LockedExit:
			return lockedTile
		case level.UnlockedExit:
			return unlockedTile
		default:
			return unknownTile
		}
	}
}
