// Package raindrops translates key number(s) to phrase
package raindrops

import (
	"math"
	"strconv"
)

// use map to store factors
func factors(N int) map[int]int {
	mFact := map[int]int{1: 1, N: 1}
	for i := 2; i < int(math.Sqrt(float64(N)))+1; i++ {
		if N%i == 0 {
			mFact[i] = 1
			mFact[int(N/i)] = 1
		}
	}
	return mFact
}

// Convert translates number to phrase
func Convert(N int) string {
	phrase, mFactors := "", factors(N)

	if mFactors[3] > 0 {
		phrase += "Pling"
	}
	if mFactors[5] > 0 {
		phrase += "Plang"
	}
	if mFactors[7] > 0 {
		phrase += "Plong"
	}
	if phrase == "" {
		phrase = strconv.Itoa(N)
	}

	return phrase
}
