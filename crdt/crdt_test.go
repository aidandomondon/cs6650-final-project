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
