// Package diffsquares solves this problem: https://exercism.io/my/solutions/c9aeb9ede95744e38e6e31082be0c151
package diffsquares

// SquareOfSum calculates square of sum of N int
func SquareOfSum(N int) int {
	sm := 0
	for i := 1; i < (N + 1); i++ {
		sm += i
	}
	return sm * sm
}

func SquareOfSum2(N int) int {
	sm := N * (N + 1) / 2
	return sm * sm
}

// SumOfSquares calculates sum of i*i
func SumOfSquares(N int) int {
	sm := 0
	for i := 1; i < (N + 1); i++ {
		sm += i * i
	}
	return sm
}

func SumOfSquares2(N int) int {
	// see http://mathforum.org/library/drmath/view/56920.html
	return N * (N + 1) * (2*N + 1) / 6
}

func diffLoop(N int) int {
	return SquareOfSum(N) - SumOfSquares(N)
}

func diff2(N int) int {
	return SquareOfSum2(N) - SumOfSquares2(N)
}

func diffExact(N int) int {
	return N * (N + 1) * (3*N*N - N - 2) / 12
}

// Difference calculates SquareOfSum - SumOfSquares
func Difference(N int) int {
	// exact solution
	return diffExact(N)
	/*
	   BenchmarkSquareOfSum  	20000000	        87.5 ns/op	       0 B/op	       0 allocs/op
	   BenchmarkSumOfSquares 	20000000	        74.8 ns/op	       0 B/op	       0 allocs/op
	   BenchmarkDifference   	2000000000	         0.43 ns/op	       0 B/op	       0 allocs/op

	*/

	//return diff2(N)
	/*
	   BenchmarkSquareOfSum  	20000000	        71.5 ns/op	       0 B/op	       0 allocs/op
	   BenchmarkSumOfSquares 	20000000	        62.7 ns/op	       0 B/op	       0 allocs/op
	   BenchmarkDifference   	2000000000	         0.63 ns/op	       0 B/op	       0 allocs/op

	*/

	// looping solution
	//return diffLoop(N)
	/*
	   BenchmarkSquareOfSum  	20000000	        71.0 ns/op	       0 B/op	       0 allocs/op
	   BenchmarkSumOfSquares 	20000000	        60.9 ns/op	       0 B/op	       0 allocs/op
	   BenchmarkDifference   	10000000	       137 ns/op	       0 B/op	       0 allocs/op
	*/
}
