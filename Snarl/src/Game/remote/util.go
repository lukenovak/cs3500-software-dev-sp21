package remote

import (
	"bufio"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/internal/render"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"golang.org/x/image/colornames"
)

// BlockingRead reads from a connection, but blocks until we have data in the connection
func BlockingRead(r *bufio.Reader) *[]byte {
	byteChan := make(chan []byte)
	b := make([]byte, 4096)
	go func() {
		for {
			n, _ := r.ReadBytes('\n')
			if len(n) > 0 {
				byteChan <- n
				break
			}
		}
	}()
	for {
		select {
		case b = <-byteChan:
			return &b
		default:
			continue
		}
	}
}

func UpdateGui(updateMessage PlayerUpdateMessage, gameWindow fyne.Window) {

	// converting to tiles
	tiles := make([][]*level.Tile, 0)
	for i, row := range updateMessage.Layout {
		tiles = append(tiles, make([]*level.Tile, len(row)))
		for j, tileType := range row {
			tile := level.Tile{
				Type: tileType,
			}
			tiles[i][j] = &tile
		}
	}

	// add items to tiles
	for _, object := range updateMessage.Objects {
		var item level.Item
		switch object.Type {
		case "key":
			item = level.Item{Type: level.KeyID}
		case "exit":
			item = level.Item{Type: level.LockedExit}
		}
		tiles[object.Position[0]][object.Position[1]].Item = &item
	}

	// generate actor lists
	players := make([]actor.Actor, 0)
	adversaries := make([]actor.Actor, 0)
	convertToRelative := func(pos level.Position2D) level.Position2D {
		updatePosition := updateMessage.Position.ToPos2D()
		return level.NewPosition2D(pos.Row-updatePosition.Row+2, pos.Col-updatePosition.Col+2)
	}
	for _, actorData := range updateMessage.Actors {
		switch actorData.Type {
		case "player":
			players = append(players, actor.Actor{
				Type:        actor.PlayerType,
				Name:        actorData.Name,
				Position:    convertToRelative(actorData.Position.ToPos2D()),
				RenderedObj: canvas.NewCircle(colornames.Cornflowerblue),
			})
		case "zombie":
			adversaries = append(adversaries, actor.Actor{
				Type:        actor.ZombieType,
				Name:        actorData.Name,
				Position:    convertToRelative(actorData.Position.ToPos2D()),
				RenderedObj: canvas.NewCircle(colornames.Crimson),
			})
		case "ghost":
			adversaries = append(adversaries, actor.Actor{
				Type:        actor.GhostType,
				Name:        actorData.Name,
				Position:    convertToRelative(actorData.Position.ToPos2D()),
				RenderedObj: canvas.NewCircle(colornames.Hotpink),
			})
		}
	}

	render.GuiState(tiles, players, adversaries, gameWindow)
}

func (jsonLevel Level) ToGameLevel() level.Level {
	var newLevel, err = level.NewEmptyLevel(4, 4)
	if err != nil {
		panic(err)
	}
	// generate rooms
	for _, room := range jsonLevel.Rooms {
		newOrigin := room.Origin.ToPos2D()
		err = newLevel.GenerateRectangularRoomWithLayout(newOrigin, len(room.Layout), len(room.Layout[0]), room.Layout)
		if err != nil {
			panic(err)
		}
	}

	// generate hallways
	for _, hallway := range jsonLevel.Hallways {
		newFrom := hallway.From.ToPos2D()
		newTo := hallway.To.ToPos2D()
		var newWaypoints []level.Position2D
		for _, point := range hallway.Waypoints {
			newWaypoints = append(newWaypoints, point.ToPos2D())
		}
		err = newLevel.GenerateHallway(newFrom, newTo, newWaypoints)
		if err != nil {
			panic(err)
		}
	}

	return newLevel
}

func LevelToTestLevel(inputLevel level.Level) Level {
	testRooms := make([]Room, 0)
	testHallways := make([]Hallway, 0)

	for _, node := range inputLevel.RoomDataGraph {
		switch node.Type() {
		case "room":
			testRooms = append(testRooms, roomToTestRoom(node.(*level.RoomData), inputLevel))
		case "hallway":
			testHallways = append(testHallways, hallToTestHall(node.(*level.HallData)))
		}
	}

	return Level{
		Rooms:    testRooms,
		Hallways: testHallways,
	}
}

func roomToTestRoom(room *level.RoomData, inputLevel level.Level) Room {
	return Room{
		Type: "room",
		Origin: Point{
			room.TopLeft.Row,
			room.TopLeft.Col,
		},
		Bounds: Bounds{
			Rows:    room.Size.Row,
			Columns: room.Size.Col,
		},
		Layout: tilesToArray(inputLevel.GetTiles(room.TopLeft, room.Size)),
	}
}

func hallToTestHall(hall *level.HallData) Hallway {
	waypoints := make([]Point, 0)
	for _, point := range hall.Waypoints {
		waypoints = append(waypoints, Point{
			point.Row,
			point.Col,
		})
	}
	return Hallway{
		From: Point{
			hall.Start.Row,
			hall.Start.Col,
		},
		To: Point{
			hall.End.Row,
			hall.End.Col,
		},
		Waypoints: waypoints,
	}
}

func tilesToArray(tiles [][]*level.Tile) [][]int {
	output := make([][]int, 0)
	for _, tileRow := range tiles {
		outputRow := make([]int, 0)
		for _, tile := range tileRow {
			if tile == nil {
				outputRow = append(outputRow, 0)
			} else {
				outputRow = append(outputRow, tile.Type)
			}
		}
		output = append(output, outputRow)
	}
	return output
}
