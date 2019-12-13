// Package analyze runs main analysis programs at ProHits-viz.
package analyze

import "github.com/knightjdr/prohits-viz-analysis/internal/analyze/arguments"

// Run the analysis program specified by the "analysisType" argument.
func Run() {
	arguments.Parse()
}
