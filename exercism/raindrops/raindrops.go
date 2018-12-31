// Package raindrops translates factor to phrase
package raindrops

import "strconv"

// Convert checks if given number has factor (3,5,7) then output phrase
func Convert(N int) string {
	phrase := ""
	if N%3 == 0 {
		phrase += "Pling"
	}
	if N%5 == 0 {
		phrase += "Plang"
	}
	if N%7 == 0 {
		phrase += "Plong"
	}
	if phrase == "" {
		phrase = strconv.Itoa(N)
	}
	return phrase
}
