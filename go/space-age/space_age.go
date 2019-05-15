package space

import (
	"fmt"
	"log"
	"strconv"
)

// Planet is a string representing a planet. Earth, Pluto, etc...
type Planet string

const (
	Earth   Planet = "Earth"
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"

	EarthYearSeconds = 31557600
)

// OrbitalPeriodSeconds returns the orbital period for the current planet in seconds
func (p Planet) OrbitalPeriodSeconds() float64 {
	yearRatioFromEarth := map[Planet]float64{
		Earth:   1,
		Mercury: 0.2408467,
		Venus:   0.61519726,
		Mars:    1.8808158,
		Jupiter: 11.862615,
		Saturn:  29.447498,
		Uranus:  84.016846,
		Neptune: 164.79132,
	}

	return yearRatioFromEarth[p] * EarthYearSeconds
}

// Age returns the age of the person on a specific planet for the number of seconds passed in argument
func Age(seconds float64, planet Planet) float64 {
	orbitalPeriod := seconds / planet.OrbitalPeriodSeconds()
	truncatedResult, err := strconv.ParseFloat(fmt.Sprintf("%.2f", orbitalPeriod), 2)
	if err != nil {
		log.Fatal("can't parse", orbitalPeriod, "into a float:", err)
	}
	return truncatedResult
}
