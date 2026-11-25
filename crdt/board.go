package crdt

// a drawing board
type Board struct {
	// 2D slice of pixels
	Pixels [][]Pixel
}

// Fills in the specified points on the board.
func (b *Board) AddStroke(coors []CoordinatePair) {
	for i := 0; i < len(coors); i++ {
		coor := coors[i]
		(*b).Pixels[coor.Y][coor.X].Fill()
	}
}
