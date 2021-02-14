package render

import "github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level"

const (
	wallTile = "▓"
	walkableTile = "░"
	doorTile = "D"
	unknownTile = "?"
)

func RenderRoom(room level.Room) string {
	render := ""
	for i := 0; i < len(room.Tiles[0]); i++ {
		for j := range room.Tiles {
			render = render + renderTile(room.Tiles[j][i])
		}
		render = render + "\n"
	}
	return render
}

func renderTile(tile *level.Tile) string {
	switch tile.Type {
	case level.Wall:
		return wallTile
	case level.Walkable:
		return walkableTile
	case level.Door:
		return doorTile
	default:
		return unknownTile

	}
}