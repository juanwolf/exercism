// Package raindrops converts a number to a raindrop.
package raindrops

import (
	"fmt"
)

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
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return fmt.Sprintf("%d", number)
	}

	return result
}
