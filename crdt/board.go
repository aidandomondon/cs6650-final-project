package crdt

import "fmt"

// a drawing board
type Board struct {
	// 2D slice of pixels
	Pixels [][]Pixel
}

func (b *Board) getWidth() int {
	return len((*b).Pixels[0])
}

func (b *Board) getHeight() int {
	return len((*b).Pixels)
}

// returns true if the given coordinate pair
// is within this board's size limits, false otherwise.
// ASSUMES THE BOARD IS RECTANGULAR
func (b *Board) hasCoordinate(c *CoordinatePair) bool {
	if c.X < 0 || b.getWidth() <= c.X {
		return false
	}
	if c.Y < 0 || b.getHeight() <= c.Y {
		return false
	}
	return true
}

// Fills in the specified points on the board.
func (b *Board) AddStroke(coors []CoordinatePair) error {
	for i := 0; i < len(coors); i++ {
		coor := coors[i]
		if (*b).hasCoordinate(&coor) {
			(*b).Pixels[coor.Y][coor.X].Fill()
		} else {
			return fmt.Errorf(
				"Coordinate pair (%d, %d) is out of bounds for board of dimension (%d, %d)",
				coor.X, coor.Y,
				(*b).getWidth(), (*b).getHeight(),
			)
		}
	}
	return nil
}
