// Package scrabble calculates score
package scrabble

import (
	"strings"
)

var mScore = map[string]int{
	"A, E, I, O, U, L, N, R, S, T": 1,
	"D, G":          2,
	"B, C, M, P":    3,
	"F, H, V, W, Y": 4,
	"K":             5,
	"J, X":          8,
	"Q, Z":          10}

var bMapDefined = false

func defineMap() {
	for k, v := range mScore {
		for _, i := range strings.Split(k, ",") {
			mScore[strings.TrimSpace(i)] = v
		}
	}
	// remove initial entry
	for k := range mScore {
		if len(k) > 1 {
			delete(mScore, k)
		}
	}
	bMapDefined = true
}

// Score calculates scrabble score
func Score(word string) int {
	if bMapDefined != true {
		defineMap()
	}

	score := 0
	word = strings.ToUpper(word)
	for _, c := range word {
		score += mScore[string(c)]
	}
	return score
}
