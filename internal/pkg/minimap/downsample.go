package minimap

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"

func downsampleIfNeeded(data *Data) [][]float64 {
	if downsample.Should(data.Matrices.Abundance, data.DownsampleThreshold) {
		return downsample.Matrix(data.Matrices.Abundance, data.DownsampleThreshold)
	}
	return data.Matrices.Abundance
}
