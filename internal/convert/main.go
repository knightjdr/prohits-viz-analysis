// Package convert takes a file from ProHits-viz V1 and converts it to V2 JSON.
package convert

import (
	"path/filepath"

	"github.com/knightjdr/prohits-viz-analysis/internal/convert/settings"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/read/csv"
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

	/*
		// Generate interactive file
		interactiveData := interactive.Data{
			Filename:   fmt.Sprintf("interactive/%s.json", imageType),
			ImageType:  imageType,
			Matrices:   matrices,
			Minimap:    "minimap/minimap.png",
			Parameters: parameters,
		}
		interactive.ParseHeatmap(&interactiveData)
	*/

	mapFolder := filepath.Join(".", "minimap")
	fs.Instance.RemoveAll(mapFolder)
}
