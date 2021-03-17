package level

import "math"

type Position2D struct {
	Row int
	Col int
}

func (d Position2D) Equals(cmpPos Position2D) bool {
	return d.Row == cmpPos.Row && d.Col == cmpPos.Col
}

func (d Position2D) AddPosition(addend Position2D) Position2D {
	return NewPosition2D(d.Row+addend.Row, d.Col+addend.Col)
}

func (d Position2D) GetManhattanDistance(cmpPos Position2D) int {
	return max(d.Row-cmpPos.Row, cmpPos.Row-d.Row) + max(d.Col-cmpPos.Col, cmpPos.Col-d.Col)
}

func NewPosition2D(row int, col int) Position2D {
	return Position2D{
		row, col,
	}
}

// returns the position with the highest coordinates from each position
func getMaxPosition(pos Position2D, pos2 Position2D) Position2D {
	return NewPosition2D(max(pos.Row, pos2.Row), max(pos.Col, pos2.Col))
}

func getListMaxPosition(posList []Position2D) Position2D {
	currMax := NewPosition2D(int(math.Inf(-1)), int(math.Inf(-1)))
	for _, pos := range posList {
		currMax = getMaxPosition(currMax, pos)
	}
	return currMax
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
