package render

import (
	"fyne.io/fyne/v2"
	canvas2 "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"image/color"
)

func GUILevel(levelToRender level.Level) *fyne.Container {
	var canvasObjects []fyne.CanvasObject
	for x := range levelToRender.Tiles {
		for y := range levelToRender.Tiles[x] {
			canvasObjects = append(canvasObjects, renderGuiTile(levelToRender.Tiles[x][y]))
		}
	}
	canvasObjects = reverseRectArray(canvasObjects)
	rectContainer := container.New(layout.NewGridLayout(levelToRender.Size.X))
	for _, canvasObj := range canvasObjects {
		rectContainer.Add(canvasObj)
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
			TextSize: 32,
			TextStyle: fyne.TextStyle{Bold: true,},
		}
		return &text
	}
	return canvas2.NewRectangle(rectColor)
}

func reverseRectArray(rectArray []fyne.CanvasObject) []fyne.CanvasObject {
	reversedArray := make([]fyne.CanvasObject, len(rectArray))
	for i, j := 0, len(rectArray) - 1; i < len(rectArray); i, j = i + 1, j - 1 {
		reversedArray[i] = rectArray[j]
	}
	return reversedArray
}