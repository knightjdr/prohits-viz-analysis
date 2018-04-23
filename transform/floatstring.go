package transform

import "strconv"

// FloatToString converts a float to a string
func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}
