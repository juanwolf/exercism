// Package raindrops converts a number to a raindrop.
package raindrops

import (
	"fmt"
)

// Sound represents a sound and the factor associated to it.
type Sound struct {
	Factor int
	Sound  string
}

// SoundList is the list of Sound supported in this exercise.
var SoundList = []Sound{
	Sound{
		Factor: 3,
		Sound:  "Pling",
	},
	Sound{
		Factor: 5,
		Sound:  "Plang",
	},
	Sound{
		Factor: 7,
		Sound:  "Plong",
	},
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

	for _, sound := range SoundList {
		if number%sound.Factor == 0 {
			result += sound.Sound
		}
	}

	if result == "" {
		return fmt.Sprintf("%d", number)
	}

	return result
}
