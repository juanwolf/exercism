// Package leap provides the API to know if a year is a leap year or not.
package leap

// IsLeapYear should return true if the year passed in parameter is a leap year, false if not.
func IsLeapYear(year int) bool {
	isMultipleOfFour := (year % 4) == 0
	isMultipleOfOneHundred := (year % 100) == 0
	isMultipleOfFourHundred := (year % 400) == 0

	return isMultipleOfFour && !(isMultipleOfOneHundred && !isMultipleOfFourHundred)
}
