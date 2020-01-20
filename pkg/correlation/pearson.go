package correlation

import (
	goMath "math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/stats"
)

// Pearson performs correlation using the Pearson statistic.
func Pearson(x, y []float64) float64 {
	n := int(goMath.Min(float64(len(x)), float64(len(y))))

	mean := map[string]float64{
		"x": stats.MeanFloat(x[0:n]),
		"y": stats.MeanFloat(y[0:n]),
	}

	numerator := float64(0)
	denominatorA := float64(0)
	denominatorB := float64(0)
	for i := 0; i < n; i++ {
		numerator += (x[i] - mean["x"]) * (y[i] - mean["y"])
		denominatorA += goMath.Pow(x[i]-mean["x"], 2)
		denominatorB += goMath.Pow(y[i]-mean["y"], 2)
	}

	return numerator / (goMath.Sqrt(denominatorA) * goMath.Sqrt(denominatorB))
}
