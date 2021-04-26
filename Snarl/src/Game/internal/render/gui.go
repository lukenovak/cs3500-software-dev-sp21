package render

import (
	"fyne.io/fyne/v2"
	canvas2 "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"image/color"
)

const (
	fontSize = 16
)

// GuiState renders a given state in the given Fyne window. It is the only externally exposed rendering function, and
// will be called at each update to update the current rendered level
func GuiState(stateLevelTiles [][]*level.Tile, statePlayers []actor.Actor, stateAdversaries []actor.Actor, gameWindow fyne.Window) {
	levelSize := level.NewPosition2D(len(stateLevelTiles), len(stateLevelTiles[0]))
	levelTiles := renderGuiLevel(stateLevelTiles)
	renderGuiActors(levelTiles, statePlayers, levelSize)
	renderGuiActors(levelTiles, stateAdversaries, levelSize)
	windowContainer := container.New(layout.NewGridLayout(levelSize.Row))
	for _, renderedTile := range levelTiles {
		windowContainer.Add(renderedTile)
	}
	gameWindow.SetContent(windowContainer)
	gameWindow.Show()
}

// helper for GuiState, renders a tile layout
func renderGuiLevel(levelToRender [][]*level.Tile) []*fyne.Container {
	levelSize := level.NewPosition2D(len(levelToRender), len(levelToRender[0]))
	tileContainers := make([]*fyne.Container, len(levelToRender[0])*len(levelToRender))
	for row := range levelToRender {
		for col := range levelToRender {
			tileContainers[calc1DPosition(level.NewPosition2D(row, col), levelSize)] = renderGuiTile(levelToRender[row][col])
		}
	}
	return tileContainers
}

// renderGuiTile creates a single tile to be rendered
func renderGuiTile(tileToRender *level.Tile) *fyne.Container {
	var rectColor color.Color
	var containerContent fyne.CanvasObject

	// local function for rendering text
	newTileText := func(text string) fyne.CanvasObject {
		textRender := canvas2.Text{
			Color:     color.RGBA{R: 180, G: 180, B: 180},
			Text:      text,
			Alignment: fyne.TextAlignCenter,
			TextSize:  fontSize,
			TextStyle: fyne.TextStyle{Bold: true},
		}
		return &textRender
	}

	// create the tile based on the tile's type. Render 0 if nil
	if tileToRender == nil {
		rectColor = color.RGBA{R: 0, G: 0, B: 0}
	} else {
		switch tileToRender.Type {
		case level.Wall:
			rectColor = color.RGBA{R: 180, G: 180, B: 0}
		case level.Walkable:
			rectColor = color.RGBA{R: 20, G: 180, B: 20}
		case level.Door:
			containerContent = newTileText(doorTile)
		default:
			containerContent = newTileText(unknownTile)
		}
	}

	// if we don't have text, we have a color so render a rectangle with that color
	if containerContent == nil {
		containerContent = canvas2.NewRectangle(rectColor)
	}

	tileContainer := container.New(layout.NewMaxLayout())
	tileContainer.Add(containerContent)

	if tileToRender != nil && tileToRender.Item != nil {
		switch tileToRender.Item.Type {
		case level.KeyID:
			text := canvas2.NewText(keyTile, color.Black)
			text.TextSize = 24
			tileContainer.Add(text)
		case level.LockedExit:
			text := canvas2.NewText("L", color.Black)
			text.TextSize = 24
			tileContainer.Add(text)

		case level.UnlockedExit:
			text := canvas2.NewText(unlockedTile, color.Black)
			text.TextSize = 24
			tileContainer.Add(text)
		}
	}

	tileContainer.Resize(fyne.NewSize(100, 100))
	return tileContainer
}

// renders players on the GUI with an already existing grid array of containers
func renderGuiActors(tileContainers []*fyne.Container,
	actors []actor.Actor,
	levelSize level.Position2D) {
	for _, actorToRender := range actors {
		tilePos := calc1DPosition(actorToRender.Position, levelSize)
		tileContainers[tilePos].Add(actorToRender.RenderedObj)
	}
}

// utility function to find the 1d array index of a position in a level
func calc1DPosition(pos level.Position2D, levelSize level.Position2D) int {
	return pos.Row*levelSize.Col + pos.Col
}
