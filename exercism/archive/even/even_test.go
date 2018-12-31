package even

import "testing"

func TestEven(t *testing.T) {
	for _, tc := range testCases {
		got, err := Even(tc.iNum)

		if got != tc.bEven {
			t.Fatalf("Even(%q) = %s, expect %s.",
				tc.iNum, got, tc.bEven)
		}

		// we do not expect error
		if err != nil {
			t.Fatalf("Even(%q) returned error: %v when expecting none.",
				tc.s1, err)
		}

	}
}
