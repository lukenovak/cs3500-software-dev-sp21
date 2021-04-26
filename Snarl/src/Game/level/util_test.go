package level

import "testing"

func TestNewPosition2D(t *testing.T) {
	testPosn := Position2D{0, 0}
	newPosn := NewPosition2D(0, 0)

	if !testPosn.Equals(newPosn) {
		t.Fail()
	}
}

func TestPosition2D_Equals(t *testing.T) {
	testPosn := NewPosition2D(0, 0)
	equalPosn := NewPosition2D(0, 0)
	unequalPosn := NewPosition2D(1, 0)

	if !testPosn.Equals(equalPosn) {
		t.Fail()
	}

	if testPosn.Equals(unequalPosn) {
		t.Fail()
	}
}

func TestPosition2D_AddPosition(t *testing.T) {
	testPosn := NewPosition2D(0, 0)
	addPosn := NewPosition2D(34, 11)

	if !testPosn.AddPosition(addPosn).Equals(addPosn) {
		t.Fail()
	}
}
