package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts letter frequency
func ConcurrentFrequency(words []string) FreqMap {

	chanRes := make(chan FreqMap, len(words))

	for _, w := range words {
		go func(s string) {
			chanRes <- Frequency(s)
		}(w)
	}

	mRes := FreqMap{}
	for range words {
		for k, v := range <-chanRes {
			mRes[k] += v
		}
	}
	return mRes

}
