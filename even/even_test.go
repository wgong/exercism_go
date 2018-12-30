package even

import "testing"

func TestEven(t *testing.T) {

	if !Even(2) {
		t.Log("2 should be even!")
		t.Fail()
	}

	for _, tc := range testCases {
		got := Even(tc.iNum)

		if got != tc.bEven {
			t.Fatalf("Even(%q) = %v, expect %v.",
				tc.iNum, got, tc.bEven)
		}
	}

}

func BenchmarkEven(b *testing.B) {
	// bench combined time to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Even(tc.iNum)
		}
	}
}
