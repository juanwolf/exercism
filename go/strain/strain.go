package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
//

type predicate[T any] func(T) bool

func filter[T any](in []T, fn predicate[T]) ([]T, []T) {
	keep := []T{}
	discard := []T{}
	for _, v := range in {
		if fn(v) {
			keep = append(keep, v)
		} else {
			discard = append(discard, v)
		}
	}
	return keep, discard
}

func Discard[T any](in []T, fn predicate[T]) []T {
	_, discard := filter(in, fn)
	return discard
}

func Keep[T any](in []T, fn predicate[T]) []T {
	keep, _ := filter(in, fn)
	return keep
}
