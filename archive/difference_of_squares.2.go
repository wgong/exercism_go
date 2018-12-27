// Package diffsquares solves this problem: https://exercism.io/my/solutions/c9aeb9ede95744e38e6e31082be0c151
package diffsquares

// SquareOfSum calculates square of sum of N int
func SquareOfSum(N int) int {
	sm := N * (N + 1) / 2
	return sm * sm
}

/*
func SquareOfSum(N int) int {
	sm := 0
	for i := 1; i < (N + 1); i++ {
		sm += i
	}
	return sm * sm
}
*/

// SumOfSquares calculates sum of i*i
func SumOfSquares(N int) int {
	// see http://mathforum.org/library/drmath/view/56920.html
	return N * (N + 1) * (2*N + 1) / 6
}

/*
func SumOfSquares(N int) int {
	sm := 0
	for i := 1; i < (N + 1); i++ {
		sm += i * i
	}
	return sm
}

func diffExact(N int) int {
	return N * (N + 1) * (3*N*N - N - 2) / 12
}

*/

// Difference calculates SquareOfSum - SumOfSquares
func Difference(N int) int {
	return SquareOfSum(N) - SumOfSquares(N)
}

/*
$ go test -v --bench . --benchmem

goos: linux
goarch: amd64
BenchmarkSquareOfSum  	2000000000	         0.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumOfSquares 	2000000000	         0.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkDifference   	2000000000	         0.62 ns/op	       0 B/op	       0 allocs/op

*/
