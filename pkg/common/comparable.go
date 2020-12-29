package common

// Comparable interface compares two reference types.
// It is equal to zero if the two are the same
// It is less the zero if the current is minor of the input one
// > 0 otherwise
type Comparable interface {
	CompareTo(o Comparable) int
}
