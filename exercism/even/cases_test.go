package even

var testCases = []struct {
	iNum  int
	bEven bool
}{
	{100, true},
	{103, false},
	{0, true},
	{-1, false},
	{-4, true},
	{-41, false},
}
