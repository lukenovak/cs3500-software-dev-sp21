package render

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level"
	"testing"
)

func TestRenderRoom(t *testing.T) {
	room, _ := level.GenerateRectangularRoom(level.NewPosition2D(0, 0), 3, 3, nil)
	render := RenderRoom(room)
	exampleRender :=
			"▓▓▓\n" +
			"▓░▓\n" +
			"▓▓▓\n"

	if render != exampleRender {
		t.Fail()
	} else {
		print(render)
	}
}