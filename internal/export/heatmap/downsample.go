package heatmap

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/downsample"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

// DownsampleData for a heatmap.
func DownsampleData(data *Heatmap, matrices *types.Matrices, downsampleTreshold int) {
	if downsample.Should(matrices.Abundance, downsampleTreshold) {
		startingSize := len(matrices.Abundance)
		matrices.Abundance = downsample.Matrix(matrices.Abundance, downsampleTreshold)
		downsampleScale := float64(len(matrices.Abundance)) / float64(startingSize)

		data.Annotations.List = adjustAnnotations(data.Annotations.List, downsampleScale)
		data.Markers.List = adjustMarkers(data.Markers.List, downsampleScale)
	}
}

func adjustAnnotations(annotations map[string]types.Annotation, scale float64) (adjusted map[string]types.Annotation) {
	if len(annotations) == 0 {
		return
	}

	adjusted = make(map[string]types.Annotation, 0)
	for key, annotation := range annotations {
		adjusted[key] = types.Annotation{
			Text: annotation.Text,
			Position: types.AnnotationPosition{
				X: annotation.Position.X * scale,
				Y: annotation.Position.Y * scale,
			},
		}
	}

	return
}

func adjustMarkers(markers map[string]types.Marker, scale float64) (adjusted map[string]types.Marker) {
	if len(markers) == 0 {
		return
	}

	adjusted = make(map[string]types.Marker, 0)
	for key, marker := range markers {
		adjusted[key] = types.Marker{
			Height: int(math.Round(float64(marker.Height) * scale)),
			Width:  int(math.Round(float64(marker.Width) * scale)),
			X:      marker.X * scale,
			Y:      marker.Y * scale,
		}
	}

	return
}
