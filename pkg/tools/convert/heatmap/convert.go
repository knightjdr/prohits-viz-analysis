// Package heatmap converts an interactive file from ProHits-viz V1 to V2 format
package heatmap

import (
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/read/csv"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/convert/heatmap/settings"
)

func Convert(filename string) {
	headerMap := map[string]string{
		"column": "condition",
		"row":    "readout",
		"score":  "score",
		"value":  "abundance",
	}
	csv := csv.ReadToSliceViaHeader(filename, '\t', headerMap)
	settings := settings.Parse(csv)

	matrices := createMatrices(&csv, settings.ScoreType)
	files.CreateFolders([]string{"interactive", "minimap"})
	createMinimap(matrices, settings)
	createInteractive(matrices, settings)

	mapFolder := filepath.Join(".", "minimap")
	fs.Instance.RemoveAll(mapFolder)
}
