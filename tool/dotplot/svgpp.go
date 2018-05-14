package dotplot

import (
	"io/ioutil"

	"github.com/knightjdr/prohits-viz-analysis/svg"
)

// SvgBB draws a prey prey heatmap.
func SvgPP(dist [][]float64, sorted []string, colorSpace string) {
	// Heatmap params.
	params := map[string]interface{}{
		"colLabel":         "Preys",
		"colorSpace":       colorSpace,
		"invert":           true,
		"maximumAbundance": float64(1),
		"rowLabel":         "Preys",
	}
	heatmap := svg.Heatmap(dist, sorted, sorted, params)
	ioutil.WriteFile("svg/prey-prey.svg", []byte(heatmap), 0644)
	return
}
