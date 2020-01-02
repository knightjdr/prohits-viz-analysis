package minimap

import "github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"

func downsampleIfNeeded(data *Data) {
	if downsample.Should(data.Matrices.Abundance, data.DownsampleThreshold) {
		data.Matrices.Abundance = downsample.Matrix(data.Matrices.Abundance, data.DownsampleThreshold)
	}
}
