// Package scrabble calculates score
package scrabble

import "strings"

// Score calculates scrabble score
func Score(word string) int {
	// initial lookup
	mStringScore := map[string]int{
		"A, E, I, O, U, L, N, R, S, T": 1,
		"D, G":          2,
		"B, C, M, P":    3,
		"F, H, V, W, Y": 4,
		"K":             5,
		"J, X":          8,
		"Q, Z":          10}

	// lookup per letter
	mLetterScore := map[string]int{}
	for k, v := range mStringScore {
		//fmt.Println(k,v)
		for _, i := range strings.Split(k, ",") {
			mLetterScore[strings.TrimSpace(i)] = v
		}
	}

	score := 0
	word = strings.ToUpper(word)
	for _, c := range word {
		score += mLetterScore[string(c)]
	}
	return score
}
