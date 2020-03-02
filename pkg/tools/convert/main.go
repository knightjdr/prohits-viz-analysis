// Package convert takes a file from ProHits-viz V1 and converts it to V2 JSON.
package convert

import (
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/read/csv"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/convert/settings"
)

// File converts a file from v1 format to v2.
func File() {
	file := parseArguments()
	headerMap := map[string]string{
		"column": "condition",
		"row":    "readout",
		"score":  "score",
		"value":  "abundance",
	}
	csv := csv.Read(file, '\t', headerMap)
	settings := settings.Parse(csv)

	matrices := createMatrices(&csv, settings.ScoreType)

	files.CreateFolders([]string{"interactive", "minimap"})
	createMinimap(matrices, settings)
	createInteractive(matrices, settings)

	mapFolder := filepath.Join(".", "minimap")
	fs.Instance.RemoveAll(mapFolder)
}
