// Package hamming implements http://rosalind.info/problems/hamm/
package hamming

import "errors"

// Distance counts how many strands are different
func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("input lengths mismatch")
    }
    
    dif := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            dif++
        }
    }
    
    return dif, nil
}
