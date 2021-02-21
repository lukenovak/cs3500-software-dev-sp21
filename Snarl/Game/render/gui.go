package render

import (
	"fyne.io/fyne/v2"
	canvas2 "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level"
	"image/color"
)

func GUILevel(levelToRender level.Level) *fyne.Container {
	var rectangles []*canvas2.Rectangle
	for x := range levelToRender.Tiles {
		for y := range levelToRender.Tiles[x] {

			var rectColor color.Color
			if levelToRender.Tiles[x][y] != nil && levelToRender.Tiles[x][y].Type == level.Wall {
				rectColor = color.RGBA{R:180, G:180, B:0}
			} else if levelToRender.Tiles[x][y] != nil && levelToRender.Tiles[x][y].Type == level.Walkable {
				rectColor = color.RGBA{R:20, G:180, B:20}
			} else {
				rectColor = color.RGBA{R:180, G:180, B:180}
			}
			rectangles = append(rectangles, canvas2.NewRectangle(rectColor))
		}
	}
	rectangles = reverseRectArray(rectangles)
	rectContainer := container.New(layout.NewGridLayout(levelToRender.Size.X))
	for _, rectangle := range rectangles {
		rectContainer.Add(rectangle)
	}
	return rectContainer
}

func reverseRectArray(rectArray []*canvas2.Rectangle) []*canvas2.Rectangle {
	reversedArray := make([]*canvas2.Rectangle, len(rectArray))
	for i, j := 0, len(rectArray) - 1; i < len(rectArray); i, j = i + 1, j - 1 {
		reversedArray[i] = rectArray[j]
	}
	return reversedArray
}