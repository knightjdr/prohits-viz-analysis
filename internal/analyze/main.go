// Package analyze runs main analysis programs at ProHits-viz.
package analyze

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/analyze/arguments"
	"github.com/knightjdr/prohits-viz-analysis/internal/analyze/validate/data"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/data/filter"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/data/parser"
)

// Run the analysis program.
func Run() {
	analysis := arguments.Parse()

	parser.Read(analysis, false)
	filter.Process(analysis)
	data.Validate(analysis)

	/*
		// Transform readout abundances.
		dataset.FileData = transform.Readouts(dataset)

		// Perform analysis
		tool.Start(&dataset)
	*/
}
