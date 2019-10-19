// Package raindrops converts a number to a raindrop.
package raindrops

import (
	"fmt"
	"sort"
)

// Sounds represents the list of sounds where the key is the factor of when the sound should be produced
type Sounds map[int]string

// Factors returns the list of factors ascending sorted
func (s Sounds) Factors() []int {
	keys := []int{}

	for key, _ := range s {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

// DefaultSounds represents the only Sounds that needs to be interpreted by the Convert function for this exercise.
var DefaultSounds = Sounds{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

/*
Convert returns a string according to the number in parameter.
It adds
 'Pling' if the number has 3 as a factor
 'Plang' if the number has 5 as a factor
 'Plong' if the number has 7 as a factor
or returns the number's digit if none of the above.
*/
func Convert(number int) string {
	result := ""

	for _, factor := range DefaultSounds.Factors() {
		if number%factor == 0 {
			result += DefaultSounds[factor]
		}
	}

	if result == "" {
		return fmt.Sprintf("%d", number)
	}

	return result
}
