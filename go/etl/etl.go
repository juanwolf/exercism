package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var result map[string]int = map[string]int{}
	for k, v := range in {
		for _, character := range v {
			result[strings.ToLower(character)] = k
		}
	}
	return result
}
