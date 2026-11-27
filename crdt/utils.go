package crdt

import "iter"

// convenience struct defining a coordinate pair.
// to specify a location (a pixel) on a board.
type CoordinatePair struct {
	X int
	Y int
}

// Collection with no duplicates.
// Implemented using a map to get constant time membership checking
type Set[T comparable] struct {
	contents map[T]struct{}
}

func NewSet[T comparable](items []T) Set[T] {
	newSet := Set[T]{contents: make(map[T]struct{})}
	for _, t := range items {
		newSet.contents[t] = struct{}{}
	}
	return newSet
}

// Returns true if this Set contains the given item t.
func (s *Set[T]) Contains(t T) bool {
	_, ok := (*s).contents[t]
	return ok
}

// Adds the given item to the set in-place.
func (s *Set[T]) Add(t T) {
	(*s).contents[t] = struct{}{}
}

// Returns an iterator over the contents of this set.
func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for t := range s.contents {
			// If yield returns false, the caller used 'break' or 'return'
			if !yield(t) {
				return
			}
		}
	}
}

// Returns the number of elements in this set.
func (s *Set[T]) Length() int {
	return len((*s).contents)
}
