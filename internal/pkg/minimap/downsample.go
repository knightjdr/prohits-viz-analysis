package minimap

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"

func downsampleIfNeeded(data *Data) {
	if downsample.Should(data.Matrices.Abundance, 1000) {
		data.Matrices.Abundance = downsample.Matrix(data.Matrices.Abundance, 1000)
	}
}
