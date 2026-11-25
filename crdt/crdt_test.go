package crdt

import "testing"

// Tests that attempting to add a stroke
// with a coordinate that is out of bounds
// for a board throws an error
func TestCannotStrokeOutOfBounds(t *testing.T) {
	b := Board{[][]Pixel{
		{Pixel{IsFilled: false}, Pixel{IsFilled: false}},
		{Pixel{IsFilled: false}, Pixel{IsFilled: false}},
	}}
	outOfBoundsCoordinate := CoordinatePair{X: 0, Y: 2}
	err := b.AddStroke([]CoordinatePair{outOfBoundsCoordinate})
	expectedMessage := "Coordinate pair (0, 2) is out of bounds for board of dimension (2, 2)"
	if err == nil || err.Error() != expectedMessage {
		t.FailNow()
	}
}

// Tests tat adding a valid stroke changes the 2D slice
// of the Board to have the expected pixels set to true
func TestStrokeChangesInternalSlice(t *testing.T) {
	b := Board{[][]Pixel{
		{Pixel{IsFilled: false}, Pixel{IsFilled: false}},
		{Pixel{IsFilled: false}, Pixel{IsFilled: false}},
	}}
	stroke := []CoordinatePair{
		{X: 0, Y: 0},
		{X: 1, Y: 1},
	}
	err := b.AddStroke(stroke)
	if err != nil {
		t.FailNow()
	}
	expected := Board{[][]Pixel{
		{Pixel{IsFilled: true}, Pixel{IsFilled: false}},
		{Pixel{IsFilled: false}, Pixel{IsFilled: true}},
	}}
	if !b.equals(&expected) {
		t.FailNow()
	}
}
