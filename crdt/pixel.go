package crdt

// represents one pixel of a drawing board
type Pixel struct {
	// true if this pixel is filled in, false otherwise
	IsFilled bool
}

// fill in this pixel
func (p *Pixel) Fill() {
	(*p).IsFilled = true
}
