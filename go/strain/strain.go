package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
//

type predicate[T any] func(T) bool

func Discard[T any](in []T, fn predicate[T]) []T {
	res := []T{}
	for _, v := range in {
		if !fn(v) {
			res = append(res, v)
		}
	}
	return res
}

func Keep[T any](in []T, fn predicate[T]) []T {
	res := []T{}
	for _, v := range in {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}
