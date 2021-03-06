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

// helper to split a string into N slices
func splitToSlice(s string, N int) []string {
	if N < 2 {
		return []string{s}
	}
	chunkSize := int(len(s)/N) + 1
	var s2 []string
	for i := 0; i < N-1; i++ {
		s2 = append(s2, s[i*chunkSize:(i+1)*chunkSize])
	}
	s2 = append(s2, s[(N-1)*chunkSize:])
	return s2
}
