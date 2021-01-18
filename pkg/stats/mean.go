package stats

import "strconv"

// MeanString calculates the mean of a []string (must be convertable to float).
func MeanString(slice []string) float64 {
	if len(slice) == 0 {
		return 0
	}

	mean := float64(0)
	for _, value := range slice {
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return 0
		}
		mean += parsed
	}

	return mean / float64(len(slice))
}
