package hamming

import (
	"errors"
)

// Distance calculates the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
	runesA, runesB := []rune(a), []rune(b)
	if len(runesA) != len(runesB) {
		return 0, errors.New("distance string parameters are not the same length")
	}

	diff := 0
	for i, nucleotide := range runesA {
		if nucleotide != runesB[i] {
			diff++
		}
	}

	return diff, nil
}
