package stats

// MeanFloat calculates the mean of a []float64
func MeanFloat(slice []float64) float64 {
	if len(slice) == 0 {
		return 0
	}

	mean := float64(0)
	for _, value := range slice {
		mean += value
	}

	return mean / float64(len(slice))
}

// MeanInt calculates the mean of a []int
func MeanInt(slice []int) float64 {
	if len(slice) == 0 {
		return 0
	}

	mean := float64(0)
	for _, value := range slice {
		mean += float64(value)
	}

	return mean / float64(len(slice))
}
