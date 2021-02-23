package level


// TODO: Make this not hard coded (will be done at a later date)
func GenerateNewLevel(size Position2D) (*Level, error){
	newLevel, err := NewEmptyLevel(size.X, size.Y)
	if err != nil {
		panic(err)
	}
	// first room from 0,0 to 5, 4
	err = newLevel.GenerateRectangularRoom(NewPosition2D(0,0),
		6,
		5,
		[]Position2D{NewPosition2D(1, 4), NewPosition2D(5, 3)})
	if err != nil {
		return nil, err
	}

	// second room from 9, 9 to 14, 16
	err = newLevel.GenerateRectangularRoom(NewPosition2D(9,9),
		6,
		8,
		[]Position2D{NewPosition2D(9, 13)})
	if err != nil {
		return nil, err
	}

	// Third room from 20, 21 to 28, 29
	err = newLevel.GenerateRectangularRoom(NewPosition2D(20,21),
		9,
		9,
		[]Position2D{NewPosition2D(20, 25)})
	if err != nil {
		return nil, err
	}

	// connecting hallways
	hallwayWaypoints := []Position2D{{7, 3}, {7,13}}
	err = newLevel.GenerateHallway(NewPosition2D(5, 3), NewPosition2D(9, 13), hallwayWaypoints)
	if err != nil {
		return nil, err
	}

	hallwayWaypoints = []Position2D{{1, 25}}
	err = newLevel.GenerateHallway(NewPosition2D(1, 4), NewPosition2D(20, 25), hallwayWaypoints)
	if err != nil {
		return nil, err
	}

	return &newLevel, nil
}