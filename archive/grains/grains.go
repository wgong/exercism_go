package grains

import "errors"

// Total adds up all grains
func Total() uint64 {
	t, err := Square(64)
	if err == nil {
		return 2*t - 1
	}
	return 0
}

// Square calculates number of grains at a square
func Square(seq int) (uint64, error) {
	switch {
	case seq < 1 || seq > 64:
		return 0, errors.New("Invalid input")
	case seq == 1:
		return 1, nil
	default:
		return 2 << (uint(seq) - 2), nil
	}
}
