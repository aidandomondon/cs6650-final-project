package crdt

import "testing"

// Tests that attempting to add a stroke
// with a coordinate that is out of bounds
// for a board throws an error
func TestCannotStrokeOutOfBounds(t *testing.T) {
	b := Board{
		Width:        2,
		Height:       2,
		FilledPixels: NewSet([]CoordinatePair{}),
	}
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
	b := Board{
		Width:        2,
		Height:       2,
		FilledPixels: NewSet([]CoordinatePair{}),
	}
	stroke := []CoordinatePair{
		{X: 0, Y: 0},
		{X: 1, Y: 1},
	}
	err := b.AddStroke(stroke)
	if err != nil {
		t.FailNow()
	}
	expected := Board{
		Width:        2,
		Height:       2,
		FilledPixels: NewSet([]CoordinatePair{{0, 0}, {1, 1}}),
	}
	if !b.equals(&expected) {
		t.FailNow()
	}
}

func TestLookupFilledPixel(t *testing.T) {
	coor := CoordinatePair{1, 0}
	b := Board{
		Width: 2, Height: 2,
		FilledPixels: NewSet([]CoordinatePair{coor}),
	}
	if !b.Lookup(coor) {
		t.FailNow()
	}
}

func TestLookupUnfilledPixel(t *testing.T) {
	coor := CoordinatePair{1, 0}
	b := Board{
		Width: 2, Height: 2,
		FilledPixels: NewSet([]CoordinatePair{}),
	}
	if b.Lookup(coor) {
		t.FailNow()
	}
}

// Looking up a coordinate out of board's bounds should return false.
func TestLookupOutOfBoundsPixel(t *testing.T) {
	b := Board{
		Width: 2, Height: 2,
		FilledPixels: NewSet([]CoordinatePair{{0, 0}, {0, 1}, {1, 0}, {1, 1}}),
	}
	outOfBoundsCoor := CoordinatePair{0, 2}
	if b.Lookup(outOfBoundsCoor) {
		t.FailNow()
	}
}

func TestMerge(t *testing.T) {
	a := Board{
		Width: 2, Height: 2,
		FilledPixels: NewSet([]CoordinatePair{{0, 0}, {1, 1}}),
	}
	b := Board{
		Width: 2, Height: 2,
		FilledPixels: NewSet([]CoordinatePair{{0, 1}, {1, 0}}),
	}
	err, c := Merge(a, &b)
	expectedC := Board{
		Width: 2, Height: 2,
		FilledPixels: NewSet([]CoordinatePair{{0, 0}, {0, 1}, {1, 0}, {1, 1}}),
	}
	if (err != nil) || (!c.equals(&expectedC)) {
		t.FailNow()
	}
}

// Merging a board with an empty board returns an identical board
func TestMergeEmpty(t *testing.T) {
	a := Board{
		Width: 2, Height: 2,
		FilledPixels: NewSet([]CoordinatePair{{0, 0}, {1, 1}}),
	}
	b := Board{
		Width: 2, Height: 2,
		FilledPixels: NewSet([]CoordinatePair{}),
	}
	err, c := Merge(a, &b)
	if (err != nil) || (!c.equals(&a)) {
		t.FailNow()
	}
}

func TestMergeIncompatibleDimensions(t *testing.T) {
	a := Board{
		Width: 2, Height: 2,
		FilledPixels: NewSet([]CoordinatePair{{0, 0}, {1, 1}}),
	}
	b := Board{
		Width: 2, Height: 3,
		FilledPixels: NewSet([]CoordinatePair{{0, 0}, {1, 1}}),
	}
	err, c := Merge(a, &b)
	if (err == nil) || (c.equals(&a)) {
		t.FailNow()
	}
}
