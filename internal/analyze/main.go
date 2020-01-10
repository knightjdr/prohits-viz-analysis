// Package analyze runs main analysis programs at ProHits-viz.
package analyze

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/analyze/arguments"
	"github.com/knightjdr/prohits-viz-analysis/internal/analyze/dotplot"
	"github.com/knightjdr/prohits-viz-analysis/internal/analyze/settings"
	"github.com/knightjdr/prohits-viz-analysis/internal/analyze/validate/data"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/data/filter"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/data/parser"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/data/transform"
)

// Run the analysis program.
func Run() {
	analysis := arguments.Parse()

	parser.Read(analysis, false)
	filter.Process(analysis)
	data.Validate(analysis)
	transform.Abundance(analysis)

	settings.Log(analysis.Settings)

	switch analysis.Settings.Type {
	case "dotplot":
		dotplot.Generate(analysis)
	}
}
