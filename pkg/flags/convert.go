package flags

import (
	"math"
	"strconv"
)

func convertFloat(arg interface{}) float64 {
	var converted float64
	if arg != nil {
		switch arg.(type) {
		default:
			converted = float64(0)
		case float64:
			converted = arg.(float64)
		case int:
			converted = float64(arg.(int))
		case string:
			converted, _ = strconv.ParseFloat(arg.(string), 64)
		}
	}
	return converted
}

func convertInt(arg interface{}) int {
	var converted int
	if arg != nil {
		switch arg.(type) {
		default:
			converted = 0
		case float64:
			converted = int(math.Round(arg.(float64)))
		case int:
			converted = arg.(int)
		case string:
			converted, _ = strconv.Atoi(arg.(string))
		}
	}
	return converted
}

func convertString(arg interface{}) string {
	var converted string
	if arg != nil {
		switch arg.(type) {
		default:
			converted = ""
		case string:
			converted, _ = arg.(string)
		}
	}
	return converted
}
