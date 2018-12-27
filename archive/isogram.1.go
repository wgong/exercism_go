// Package isogram detects isogram word
package isogram

import "unicode"

// IsIsogram checks if word is isogram
func IsIsogram(word string) bool {
	var mDup = map[rune]int{}
	for _, c := range word {
		c = unicode.ToUpper(c)
		if c == ' ' || c == '-' {
			continue
		}
		mDup[c]++
		if mDup[c] > 1 {
			return false
		}
	}
	return true
}
