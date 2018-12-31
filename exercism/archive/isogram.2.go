// Package isogram detects isogram word
package isogram

import "unicode"

// IsIsogram checks if word is isogram
func IsIsogram(word string) bool {
	var mDup = map[rune]bool{}
	for _, c := range word {
		if unicode.IsLetter(c) == false {
			continue
		}
		if mDup[unicode.ToUpper(c)] {
			return false
		} else {
 			mDup[unicode.ToUpper(c)] = true
		}
	}
	return true
}
