package float

import (
	"math"
)

// GetRange returns a function that will map a number to an output integer range.
func GetRange(inMin, inMax, outMin, outMax float64) func(float64) float64 {
	inputRange := inMax - inMin
	outputRange := outMax - outMin
	return func(value float64) float64 {
		num := value
		if value > inMax {
			num = inMax
		} else if value < inMin {
			num = inMin
		}
		return math.Round((((num - inMin) * outputRange) / inputRange) + outMin)
	}
}
