package render

import (
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/Game/level"
	"testing"
)

func TestRenderRoom(t *testing.T) {
	testLevel := level.NewEmptyLevel(3, 3)
	err := testLevel.GenerateRectangularRoom(level.NewPosition2D(0, 0), 3, 3, nil)
	if err != nil {
		t.Fatal("unable to generate level")
	}
	render := RenderLevel(testLevel)
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