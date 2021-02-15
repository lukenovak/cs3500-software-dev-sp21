package level

type Position2D struct {
	X int
	Y int
}

var originPosition = NewPosition2D(0, 0)

func (d Position2D) Equals (cmpPos Position2D) bool {
	return d.X == cmpPos.X && d.Y == cmpPos.Y
}

func NewPosition2D(x int, y int) Position2D {
	return Position2D {
		x,y,
	}
}

// returns the position with the highest coordinates from each position
func getMaxPosition(pos Position2D, pos2 Position2D) Position2D {
	return NewPosition2D(max(pos.X, pos2.X), max(pos.Y, pos2.Y))
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}