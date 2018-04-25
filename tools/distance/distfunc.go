package distance

import (
	"errors"
	"math"
)

// DistFunc returns a function for calculate the distance between two vectors.
func DistFunc(metric string) func(x []float64, y []float64) (dist float64, err error) {
	if metric == "binary" {
		binary := func(x []float64, y []float64) (dist float64, err error) {
			if len(x) != len(y) {
				err = errors.New("Vectors for calculating distance must be equal")
				return
			}
			denominator := float64(len(x))
			numerator := float64(0)
			for i := range x {
				if x[i] > 0 && y[i] > 0 {
					numerator++
				}
			}
			dist = 1 - (numerator / denominator)
			return
		}
		return binary
	} else if metric == "canberra" {
		canberra := func(x []float64, y []float64) (dist float64, err error) {
			if len(x) != len(y) {
				err = errors.New("Vectors for calculating distance must be equal")
				return
			}
			dist = 0
			for i := range x {
				dist += math.Abs(x[i]-y[i]) / math.Abs(x[i]+y[i])
			}
			return
		}
		return canberra
	} else if metric == "jaccard" {
		jaccard := func(x []float64, y []float64) (dist float64, err error) {
			if len(x) != len(y) {
				err = errors.New("Vectors for calculating distance must be equal")
				return
			}
			denominator := float64(0)
			numerator := float64(0)
			for i := range x {
				numerator += math.Min(x[i], y[i])
				denominator += math.Max(x[i], y[i])
			}
			dist = 1 - (numerator / denominator)
			return
		}
		return jaccard
	} else if metric == "manhattan" {
		manhattan := func(x []float64, y []float64) (dist float64, err error) {
			if len(x) != len(y) {
				err = errors.New("Vectors for calculating distance must be equal")
				return
			}
			dist = 0
			for i := range x {
				dist += math.Abs(x[i] - y[i])
			}
			return
		}
		return manhattan
	} else if metric == "maximum" {
		maximum := func(x []float64, y []float64) (dist float64, err error) {
			if len(x) != len(y) {
				err = errors.New("Vectors for calculating distance must be equal")
				return
			}
			dist = 0
			for i := range x {
				diff := math.Abs(x[i] - y[i])
				if diff > dist {
					dist = diff
				}
			}
			return
		}
		return maximum
	}
	// Euclidean by default
	euclidean := func(x []float64, y []float64) (dist float64, err error) {
		if len(x) != len(y) {
			err = errors.New("Vectors for calculating distance must be equal")
			return
		}
		dist = 0
		for i := range x {
			dist += math.Pow(x[i]-y[i], 2)
		}
		dist = math.Sqrt(dist)
		return
	}
	return euclidean
}
