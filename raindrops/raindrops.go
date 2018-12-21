// Package raindrops translates key number(s) to phrase
package raindrops

import (
	"fmt"
	"math"
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
	mFactors := factors(N)
	lNumbers := []int{3, 5, 7}
	mWords := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}

	phrase := ""
	for _, k := range lNumbers {
		if _, ok := mFactors[k]; ok {
			phrase += mWords[k]
		}
	}
	if phrase == "" {
		phrase = fmt.Sprintf("%d", N)
	}

	return phrase
}
