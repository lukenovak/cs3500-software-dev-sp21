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

func GuiState(stateLevelTiles [][]*level.Tile, statePlayers []actor.Actor, stateAdversaries []actor.Actor, gameWindow fyne.Window) {
	levelSize := level.NewPosition2D(len(stateLevelTiles), len(stateLevelTiles[0]))
	levelTiles := renderGuiLevel(stateLevelTiles)
	renderGuiActors(levelTiles, statePlayers, levelSize, renderPlayer)
	renderGuiActors(levelTiles, stateAdversaries, levelSize, renderAdversary)
	windowContainer := container.New(layout.NewGridLayout(levelSize.Row))
	for _, renderedTile := range levelTiles {
		windowContainer.Add(renderedTile)
	}
	gameWindow.SetContent(windowContainer)
	gameWindow.Show()
}

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

// creates a single tile to be rendered
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
		case level.LockedExit:
			containerContent = newTileText("L")
		case level.UnlockedExit:
			containerContent = newTileText(unlockedTile)
		default:
			containerContent = newTileText(unknownTile)
		}
	}

	if containerContent == nil {
		containerContent = canvas2.NewRectangle(rectColor)
	}

	tileContainer := container.New(layout.NewMaxLayout())
	tileContainer.Add(containerContent)

	if tileToRender != nil && tileToRender.Item != nil {
		switch tileToRender.Item.Type {
		case level.KeyID:
			tileContainer.Add(canvas2.NewText("K", color.Black))
		}
	}

	tileContainer.Resize(fyne.NewSize(100, 100))
	return tileContainer
}

// renders players on the GUI with an already existing grid array of containers
func renderGuiActors(tileContainers []*fyne.Container,
	actors []actor.Actor,
	levelSize level.Position2D,
	renderFunc func(*fyne.Container)) {
	for _, actorToRender := range actors {
		tilePos := calc1DPosition(actorToRender.Position, levelSize)
		renderFunc(tileContainers[tilePos])
	}
}

// utility function to find the 1d array index of a position in a level
func calc1DPosition(pos level.Position2D, levelSize level.Position2D) int {
	return pos.Row*levelSize.Col + pos.Col
}

/* ----------------------- Actor Render functions ----------------------------- */

func renderPlayer(baseContainer *fyne.Container) {
	baseContainer.Add(canvas2.NewCircle(color.RGBA{R: 150, G: 150, B: 150, A: 255}))
}

func renderAdversary(baseContainer *fyne.Container) {
	baseContainer.Add(canvas2.NewCircle(color.RGBA{R: 200, G: 100, B: 100, A: 255}))
}
