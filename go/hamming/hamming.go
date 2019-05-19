package hamming

import (
	"fmt"
)

// Distance calculates the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Distance string parameters are not the same length")
	}

	diff := 0
	for i, nucleotide := range a {
		if nucleotide != rune(b[i]) {
			diff++
		}
	}
	return diff, nil
}
