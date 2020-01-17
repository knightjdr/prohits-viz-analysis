package correlation

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/sort"
)

// Modified from https://github.com/scipy/scipy/blob/v0.14.0/scipy/stats/stats.py#L2608
// For tied ranks, the average method is used.
func spearman(dataX, dataY []float64) float64 {
	n := int(math.Min(float64(len(dataX)), float64(len(dataY))))
	x := dataX[0:n]
	y := dataY[0:n]

	xRanks := rankArray(x)
	yRanks := rankArray(y)

	return pearson(xRanks, yRanks)
}

func rankArray(x []float64) []float64 {
	n := len(x)
	ranks := make([]float64, n)
	order := sort.ArgsortFloat(x)

	dupcount := 0
	tieRank := float64(0)
	for i := 0; i < n; i++ {
		inext := i + 1
		if i == n-1 || x[order[i]] != x[order[inext]] {
			tieRank = float64(inext) - 0.5*float64(dupcount)
			for j := i - dupcount; j < inext; j++ {
				ranks[order[j]] = tieRank
				dupcount = 0
			}
		} else {
			dupcount++
		}
	}

	return ranks
}
