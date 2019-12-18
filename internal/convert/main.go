// Package convert takes a file from ProHits-viz V1 and converts it to V2 JSON.
package convert

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/convert/settings"
	"github.com/knightjdr/prohits-viz-analysis/pkg/read/csv"
)

// File converts a file from v1 format to v2.
func File() {
	file := parseArguments()
	csv := csv.Read(file, '\t')
	settings := settings.Parse(csv)

	matrices := createMatrices(&csv, settings.ScoreType)

	/*
		// Create folders
		helper.CreateFolders([]string{"interactive", "minimap"})

		// Generate minimap
		mapData := minimap.Data{
			Filename:   "minimap/minimap",
			ImageType:  imageType,
			Matrices:   matrices,
			Parameters: parameters,
		}
		minimap.Write(&mapData)

		// Generate interactive file
		interactiveData := interactive.Data{
			Filename:   fmt.Sprintf("interactive/%s.json", imageType),
			ImageType:  imageType,
			Matrices:   matrices,
			Minimap:    "minimap/minimap.png",
			Parameters: parameters,
		}
		interactive.ParseHeatmap(&interactiveData)

		// Remove minimap folder.
		mapFolder := filepath.Join(".", "minimap")
		fs.Instance.RemoveAll(mapFolder)
	*/
}
