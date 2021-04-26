package level

import "math"

// Position2D is the base position type that contains a row coordinate and a column coordinate
type Position2D struct {
	Row int
	Col int
}

// Equals returns true if the called position and given position are the same.
func (d Position2D) Equals(cmpPos Position2D) bool {
	return d.Row == cmpPos.Row && d.Col == cmpPos.Col
}

// AddPosition adds two Position2D together by adding their rows and columns
func (d Position2D) AddPosition(addend Position2D) Position2D {
	return NewPosition2D(d.Row+addend.Row, d.Col+addend.Col)
}

// GetManhattanDistance returns the manhattan distance, in steps, between the called point and given point
func (d Position2D) GetManhattanDistance(cmpPos Position2D) int {
	return max(d.Row-cmpPos.Row, cmpPos.Row-d.Row) + max(d.Col-cmpPos.Col, cmpPos.Col-d.Col)
}

// NewPosition2D constructs a Position2D given a row and column
func NewPosition2D(row int, col int) Position2D {
	return Position2D{
		row, col,
	}
}

// returns the position with the highest coordinates from each position
func getMaxPosition(pos Position2D, pos2 Position2D) Position2D {
	return NewPosition2D(max(pos.Row, pos2.Row), max(pos.Col, pos2.Col))
}

// getListMaxPostion returns a Position2D with the max row and max column of all the positions in the list
func getListMaxPosition(posList []Position2D) Position2D {
	currMax := NewPosition2D(int(math.Inf(-1)), int(math.Inf(-1)))
	for _, pos := range posList {
		currMax = getMaxPosition(currMax, pos)
	}
	return currMax
}

// utility function to get the max of two integers (not native to go!!)
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
