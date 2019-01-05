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

	nBuf := len(words)
	resChan := make(chan FreqMap, nBuf)

	for _, w := range words {
		go func(s string) {
			m := FreqMap{}
			for _, r := range s {
				//				if unicode.IsLetter(r) {
				//					m[unicode.ToUpper(r)]++
				//				}
				m[r]++
			}
			resChan <- m
		}(w)
	}

	mResult := FreqMap{}
	for i := 0; i < nBuf; i++ {
		m := <-resChan
		for k, v := range m {
			mResult[k] += v
		}
	}
	return mResult

}
