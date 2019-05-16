// Package twofer is an exercise from exercism.io. The Readme within the package
// will give you a clearer understanding of what this package is about.
package twofer

import (
	"fmt"
)

// TwoFerTemplate is a template of the string expected in the
// ShareWith function
const TwoFerTemplate = "One for %s, one for me."

// ShareWith should return a string with the pattern
// "One for ${name}, on for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf(TwoFerTemplate, name)
}
