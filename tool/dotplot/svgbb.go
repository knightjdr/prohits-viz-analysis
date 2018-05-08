package dotplot

import (
	"io/ioutil"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/svg"
)

// SvgBB draws a bait bait heatmap.
func SvgBB(dist [][]float64, unsorted, sorted []string, colorSpace string) {
	// Normalize distance matrix to 1.
	maxDist := float64(0)
	normalizedDist := dist
	for _, row := range normalizedDist {
		for _, dist := range row {
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	for i, row := range normalizedDist {
		for j, dist := range row {
			normalizedDist[i][j] = dist / maxDist
		}
	}

	// Sort matrix.
	sortedMatrix, _ := hclust.Sort(normalizedDist, unsorted, sorted, "column")
	sortedMatrix, _ = hclust.Sort(sortedMatrix, unsorted, sorted, "row")

	// Heatmap params.
	params := map[string]interface{}{
		"colLabel":         "Baits",
		"colorSpace":       colorSpace,
		"invert":           true,
		"maximumAbundance": float64(1),
		"rowLabel":         "Baits",
	}
	heatmap := svg.Heatmap(sortedMatrix, sorted, sorted, params)
	ioutil.WriteFile("svg/bait-bait.svg", []byte(heatmap), 0644)
	return
}
