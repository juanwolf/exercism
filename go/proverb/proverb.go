package proverb

import (
	"fmt"
)

const (
	forWantTemplate string = "For want of a %s the %s was lost."
	lastVerse       string = "And all for the want of a %s."
)

// Proverb returns the For Want of a Nail proverb
func Proverb(rhymes []string) []string {
	var proverb []string
	for i, rhyme := range rhymes {
		if i == len(rhymes)-1 {
			proverb = append(proverb, fmt.Sprintf(lastVerse, rhymes[0]))
		} else {
			proverb = append(proverb, fmt.Sprintf(forWantTemplate, rhyme, rhymes[i+1]))
		}
	}
	return proverb
}
