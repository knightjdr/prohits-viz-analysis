package helper

import (
	"math"
)

// SetRange returns a function that will map a number to an output integer range.
// Checks to make sure the input number is within the input range.
func SetRange(inMin, inMax, outMin, outMax float64) func(inputNum float64) float64 {
	inputRange := inMax - inMin
	outputRange := outMax - outMin
	return func(inputNum float64) float64 {
		num := inputNum
		if inputNum > inMax {
			num = inMax
		} else if inputNum < inMin {
			num = inMin
		}
		return math.Round((((num - inMin) * outputRange) / inputRange) + outMin)
	}
}
