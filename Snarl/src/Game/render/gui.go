package render

import (
	"fmt"
	"fyne.io/fyne/v2"
	canvas2 "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"image/color"
)

func GuiState(stateLevel *level.Level, statePlayers []*actor.Actor, gameWindow fyne.Window) {
	levelTiles := renderGuiPlayers(renderGuiLevel(*stateLevel), statePlayers, stateLevel.Size)
	windowContainer := container.New(layout.NewGridLayout(stateLevel.Size.X))
	for _, renderedTile := range levelTiles {
		windowContainer.Add(renderedTile)
	}
	gameWindow.SetContent(windowContainer)
	gameWindow.ShowAndRun()
}

func renderGuiLevel(levelToRender level.Level) []*fyne.Container {
	var canvasObjects []fyne.CanvasObject
	for y := range levelToRender.Tiles[0] {
		for x := range levelToRender.Tiles {
			canvasObjects = append(canvasObjects, renderGuiTile(levelToRender.Tiles[x][y]))
		}
	}
	rectContainer := make([]*fyne.Container, len(canvasObjects))
	for i, canvasObj := range canvasObjects {
		canvasObj.Resize(fyne.NewSize(100, 50))
		tileContainer := container.New(layout.NewMaxLayout())
		tileContainer.Add(canvasObj)
		tileContainer.Resize(fyne.NewSize(100, 50))
		rectContainer[i] = tileContainer
	}
	return rectContainer
}

func renderGuiTile(tileToRender *level.Tile) fyne.CanvasObject {
	var rectColor color.Color
	if tileToRender == nil {
		rectColor = color.RGBA{R:180, G:180, B:180}
	} else if tileToRender.Type == level.Wall {
		rectColor = color.RGBA{R:180, G:180, B:0}
	} else if tileToRender.Type == level.Walkable {
		rectColor = color.RGBA{R:20, G:180, B:20}
	} else {
		rectColor = color.RGBA{R:180, G:180, B:180}
		text := canvas2.Text{
			Color: rectColor,
			Text: doorTile,
			Alignment: fyne.TextAlignCenter,
			TextSize: 24,
			TextStyle: fyne.TextStyle{Bold: true,},
		}
		return &text
	}
	return canvas2.NewRectangle(rectColor)
}

// renders players on the GUI with an already existing grid array of containers
func renderGuiPlayers(tileContainers []*fyne.Container, players []*actor.Actor, levelSize level.Position2D) []*fyne.Container {
	for playerNum, player := range players {
		tilePos := calc1DPosition(player.Position, levelSize)
		tileContainers[tilePos].Add(canvas2.NewCircle(color.RGBA{R: 150, G: 150, B: 150, A: 255}))
		tileContainers[tilePos].Add(canvas2.NewText(fmt.Sprintf("P%d", playerNum + 1), color.RGBA{R: 0, G: 0, B: 0, A: 255}))

	}
	return tileContainers
}

// utility function to find the 1d array index of a position in a level
func calc1DPosition(pos level.Position2D, levelSize level.Position2D) int {
	return pos.Y * levelSize.X + pos.X
}