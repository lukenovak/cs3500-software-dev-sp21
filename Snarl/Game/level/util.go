package level

type Position2D struct {
	X int
	Y int
}

func (d Position2D) Equals (cmpPos Position2D) bool {
	return d.X == cmpPos.X && d.Y == cmpPos.Y
}

func NewPosition2D(x int, y int) Position2D {
	return Position2D {
		x,y,
	}
}

func min(x, y  int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}