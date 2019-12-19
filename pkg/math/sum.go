package math

// SumFloat a slice of float64.
func SumFloat(values []float64) float64 {
	sum := float64(0)

	for _, value := range values {
		sum += value
	}

	return sum
}
