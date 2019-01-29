package circheatmap

import "github.com/knightjdr/prohits-viz-analysis/typedef"

func readoutMetrics(parameters typedef.Parameters) map[string]string {
	columns := make(map[string]string)
	columns["abundance"] = parameters.Abundance

	if len(parameters.OtherAbundance) > 0 {
		for _, abundance := range parameters.OtherAbundance {
			columns[abundance] = abundance
		}
	}

	return columns
}
