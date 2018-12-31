// Package luhn implements Luhn algorithm to valid card number
package luhn

import "regexp"

// Valid validates a card number
func Valid(cardNum string) bool {

	// cleanup card_num
	r, _ := regexp.Compile("[ ]")
	cardNum = r.ReplaceAllString(cardNum, "")

	if len(cardNum) < 2 {
		return false
	}

	r, _ = regexp.Compile("^[0-9]+$")
	if r.MatchString(cardNum) == false {
		return false
	}

	var n, sm int

	for i := len(cardNum) - 1; i > -1; i -= 2 {

		n = int(cardNum[i] - '0')
		sm += n

	}

	for i := len(cardNum) - 2; i > -1; i -= 2 {

		n = 2 * int(cardNum[i]-'0')
		if n > 9 {
			n -= 9
		}
		sm += n

	}

	if sm%10 > 0 {
		return false
	}
	return true

}
