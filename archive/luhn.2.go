// Package luhn implements Luhn algorithm to valid card number
package luhn

import "unicode"

// Valid validates a card number
func Valid(cardNum string) bool {

	if len(cardNum) < 2 {
		return false
	}

	var n int   // int value of a digit
	var sm int  // running sum of digits
	var pos int // track digit position from right

	for i := len(cardNum) - 1; i > -1; i-- {
		if unicode.IsSpace(rune(cardNum[i])) {
			continue
		}

		if !unicode.IsDigit(rune(cardNum[i])) {
			return false
		}

		pos++
		n = int(cardNum[i] - '0')

		if pos%2 == 0 {

			if (2 * n) > 9 {
				n = 2*n - 9
			} else {
				n = 2 * n
			}
		}

		sm += n

	}

	if sm%10 > 0 || (sm == 0 && pos == 1) {
		return false
	}
	return true

}
