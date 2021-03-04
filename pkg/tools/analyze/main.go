// Package analyze runs main analysis programs at ProHits-viz.
package analyze

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/data/filter"
	"github.com/knightjdr/prohits-viz-analysis/pkg/data/parser"
	"github.com/knightjdr/prohits-viz-analysis/pkg/data/transform"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/arguments"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/cc"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/correlation"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/dotplot"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/scv"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/settings"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/specificity"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/validate/data"
)

// Run the analysis program.
func Run() {
	analysis := arguments.Parse()

	parser.Read(analysis, false)
	data.Validate(analysis, []string{"readoutLength"})
	transform.Abundance(analysis)
	filter.Process(analysis)
	data.Validate(
		analysis,
		[]string{
			"data",
			"minConditions",
			"readout",
			"score",
		},
	)

	settings.Log(analysis.Settings)

	switch analysis.Settings.Type {
	case "condition-condition":
		cc.Generate(analysis)
	case "correlation":
		correlation.Generate(analysis)
	case "dotplot":
		dotplot.Generate(analysis)
	case "scv":
		scv.Generate(analysis)
	case "specificity":
		specificity.Generate(analysis)
	}
}
