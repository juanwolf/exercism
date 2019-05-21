// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
	"sort"
)

// Kind represents a type of triangle
type Kind int

const (
	// NaT represents not a triangle
	NaT = iota
	// Equ represents an equilateral triangle
	Equ = iota
	// Iso represents a Isosceles triangle
	Iso = iota
	// Deg represents a Degenerate triangle (with zero area)
	Deg = iota
	// Sca represents a Scalene triangle (all sides of different lengths)
	Sca = iota // scalene
)

func (k Kind) String() string {
	switch k {
	case NaT:
		return "Not a Triangle"
	case Equ:
		return "Equilateral"
	case Iso:
		return "Isosceles"
	// Degenerate cases not supported by original tests
	case Deg:
		return "Degenerate"
	case Sca:
		return "Scalene"
	}
	return "Not Defined"
}

// KindFromSides returns the type of triangle it is.
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Float64s(sides)

	if math.IsInf(sides[2], 1) || !(sides[0]+sides[1] >= sides[2] && sides[0] > 0) {
		return NaT
	}

	if sides[0] == sides[1] && sides[1] == sides[2] {
		return Equ
	}

	if sides[0] == sides[1] || sides[1] == sides[2] {
		return Iso
	}

	// Degenerate case
	//if sides[0]+sides[1] == sides[2] {
	//	return Deg
	//}

	return Sca
}
