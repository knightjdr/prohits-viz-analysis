package transform

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"

func normalize(analysis *types.Analysis) {
	if analysis.Settings.Normalization == "readout" {
		normalizeByReadout(analysis)
	}
	if analysis.Settings.Normalization == "total" {
		normalizeByTotalSum(analysis)
	}
}
