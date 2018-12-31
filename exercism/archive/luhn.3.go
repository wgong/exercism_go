// Package luhn implements Luhn algorithm to valid card number
package luhn

import "unicode"

// Valid validates a card number
func Valid(cardNum string) bool {
	if len(cardNum) < 2 {
		return false
	}
	var sm, pos int
	for i := len(cardNum) - 1; i > -1; i-- {
		if unicode.IsSpace(rune(cardNum[i])) {
			continue
		}
		if !unicode.IsDigit(rune(cardNum[i])) {
			return false
		}
		pos++                      // track digit position from right
		n := int(cardNum[i] - '0') // int value of a digit
		if pos%2 == 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sm += n // running sum of digits
	}
	if sm%10 > 0 || (sm == 0 && pos == 1) {
		return false
	}
	return true
}
