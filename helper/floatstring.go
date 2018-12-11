package helper

import "strconv"

// FloatToString converts a float64 to a string.
func FloatToString(num float64) (str string) {
	str = strconv.FormatFloat(num, 'f', -1, 64)
	return
}
