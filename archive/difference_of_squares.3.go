// Package diffsquares solves this problem: https://exercism.io/my/solutions/c9aeb9ede95744e38e6e31082be0c151
package diffsquares

// SquareOfSum calculates square of sum of N int
func SquareOfSum(N int) int {
	sm := N * (N + 1) / 2
	return sm * sm
}

// SumOfSquares calculates sum of i*i
func SumOfSquares(N int) int {
	// see http://mathforum.org/library/drmath/view/56920.html
	return N * (N + 1) * (2*N + 1) / 6
}

// Difference calculates SquareOfSum - SumOfSquares
func Difference(N int) int {
	return SquareOfSum(N) - SumOfSquares(N)
}
