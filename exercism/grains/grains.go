package grains

import "fmt"

// Total adds up all grains
func Total() uint64 {
	const sum uint64 = 18446744073709551615
	return sum
}

// Square calculates number of grains at a square
func Square(seq int) (uint64, error) {
	if seq < 1 || seq > 64 {
		return 0, fmt.Errorf("bad square: %d", seq)
	}
	return 1 << (uint(seq) - 1), nil
}
