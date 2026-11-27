package crdt

import (
	"fmt"
)

// A drawing board
type Board struct {
	Width  int
	Height int
	// set containing coordinates of any points that are colored in
	FilledPixels Set[CoordinatePair]
}

// Returns true if the given coordinate pair
// is within this board's size limits, false otherwise.
// ASSUMES THE BOARD IS RECTANGULAR
func (b *Board) hasCoordinate(c *CoordinatePair) bool {
	if (c.X < 0) || (b.Width <= c.X) {
		return false
	}
	if (c.Y < 0) || (b.Height <= c.Y) {
		return false
	}
	return true
}

// Fills in the specified points on the board.
func (b *Board) AddStroke(coors []CoordinatePair) error {
	for i := 0; i < len(coors); i++ {
		coor := coors[i]
		if (*b).hasCoordinate(&coor) {
			(*b).FilledPixels.Add(coor)
		} else {
			return fmt.Errorf(
				"Coordinate pair (%d, %d) is out of bounds for board of dimension (%d, %d)",
				coor.X, coor.Y,
				(*b).Width, (*b).Height,
			)
		}
	}
	return nil
}

func (b *Board) equals(other *Board) bool {
	// check if dimensions are equal
	if ((*b).Height != (*other).Height) || ((*b).Width != (*other).Width) {
		return false
	}
	// check if number of filled pixels is equal
	if (*b).FilledPixels.Length() != (*other).FilledPixels.Length() {
		return false
	}
	// check if coordinates of filled pixels are equal
	isEqual := true
	for p := range (*b).FilledPixels.All() {
		isEqual = isEqual && (*other).FilledPixels.Contains(p)
	}
	return isEqual
}

// Remaining G-Set operations
func (b *Board) Lookup(coor CoordinatePair) bool {
	// Checks whether this board has the given coordinate
	// in its list of coordinates that are colored in
	return (*b).FilledPixels.Contains(coor)
}

func Compare(s *Board, t *Board) bool {
	res := true
	for p := range s.FilledPixels.All() {
		res = res && t.Lookup(p)
	}
	return res
}

// Returns a new board where a pixel is filled in
// if it is filled in on either of board s or board t.
//
// Returns an error and an empty board if the two given boards
// do not have the same dimensions.
func Merge(s Board, t *Board) (error, Board) {
	// takes a copy of s instead of a reference so that we can use
	// the created/passed copy as the start of the new board

	// Cannot merge boards of different dimensions
	if (s.Width != (*t).Width) || (s.Height != (*t).Height) {
		error := fmt.Errorf(
			"Cannot merge boards of different dimensions (%d, %d) and (%d, %d)",
			s.Width, s.Height,
			(*t).Width, (*t).Height,
		)
		return error, Board{}
	}

	diff := []CoordinatePair{}
	for p := range (*t).FilledPixels.All() {
		if !s.Lookup(p) {
			diff = append(diff, p)
		}
	}
	s.AddStroke(diff)
	return nil, s
}
