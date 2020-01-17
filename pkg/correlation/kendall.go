package correlation

import (
	goMath "math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"github.com/knightjdr/prohits-viz-analysis/pkg/sort"
)

// Modified from https://github.com/scipy/scipy/blob/v1.4.1/scipy/stats/stats.py#L3974-L4169
func kendall(dataX, dataY []float64) float64 {
	n := int(goMath.Min(float64(len(dataX)), float64(len(dataY))))
	x := dataX[0:n]
	y := dataY[0:n]

	yIndices := sort.ArgsortFloat(y)
	x = sort.ByIndicesFloat(x, yIndices)
	y = sort.ByIndicesFloat(y, yIndices)
	yRank := getDenseRank(y)

	xIndices := sort.ArgsortFloat(x)
	x = sort.ByIndicesFloat(x, xIndices)
	yRank = sort.ByIndicesInt(yRank, xIndices)
	xRank := getDenseRank(x)

	xtie := countRankTie(xRank)
	ytie := countRankTie(yRank)

	tot := (n * (n - 1)) / 2
	if xtie == tot || ytie == tot {
		return 0
	}

	cnt := getDiffNonTies(xRank, yRank)
	ntie := countTies(cnt)
	dis := discordant(xRank, yRank)

	conMinusDis := float64(tot - xtie - ytie + ntie - 2*dis)
	tau := conMinusDis / goMath.Sqrt(float64(tot-xtie)) / goMath.Sqrt(float64(tot-ytie))
	tau = goMath.Min(1., goMath.Max(-1., tau))

	return tau
}

func getDenseRank(x []float64) []int {
	n := len(x)
	isDiff := make([]int, n)
	isDiff[0] = 1

	for i := 0; i < n-1; i++ {
		if x[i+1] != x[i] {
			isDiff[i+1] = 1
		} else {
			isDiff[i+1] = 0
		}
	}

	previous := 0
	sum := make([]int, n)
	for i, value := range isDiff {
		sum[i] = previous + value
		previous = sum[i]
	}

	return sum
}

func countRankTie(ranks []int) int {
	max := math.MaxSliceInt(ranks)
	bins := make([]int, max+1)

	for _, rank := range ranks {
		bins[rank]++
	}

	ties := make([]int, 0)
	for _, sum := range bins {
		if sum > 1 {
			ties = append(ties, sum)
		}
	}

	return countTies(ties)
}

func countTies(x []int) int {
	count := 0

	for _, value := range x {
		count += value * (value - 1) / 2
	}

	return count
}

func getDiffNonTies(x, y []int) []int {
	n := len(x)
	nonTies := make([]int, 0)
	nonTies = append(nonTies, 0)

	index := 1
	for i := 0; i < n-1; i++ {
		if x[i+1] != x[i] || y[i+1] != y[i] {
			nonTies = append(nonTies, index)
		}
		index++
	}
	nonTies = append(nonTies, index)

	diff := make([]int, len(nonTies)-1)
	for i := range diff {
		diff[i] = nonTies[i+1] - nonTies[i]
	}

	return diff
}

// Modified from kendall_dis https://github.com/scipy/scipy/blob/master/scipy/stats/_stats.pyx.
func discordant(x, y []int) int {
	sup := 1 + math.MaxSliceInt(y)
	arr := make([]int, sup+((sup-1)>>14))

	dis := 0
	i := 0
	index := 0
	k := 0
	size := len(x)

	for i < size {
		for k < size && x[i] == x[k] {
			dis += i
			index = y[k]
			for index != 0 {
				dis -= arr[index+(index>>14)]
				index = index & (index - 1)
			}
			k++
		}
		for i < k {
			index = y[i]
			for index < sup {
				arr[index+(index>>14)]++
				index += index & -index
			}
			i++
		}
	}

	return dis
}
