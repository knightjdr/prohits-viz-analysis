// Package main takes a file from ProHits-viz V1 and converts it to V2 JSON.
package main

import (
	"flag"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/filereader/columnparser"
	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/helper"
	"github.com/knightjdr/prohits-viz-analysis/interactive"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
)

func main() {
	// Get file to convert.
	file := flag.String("file", "", "ProHits-viz V1 file")
	flag.Parse()

	// Set file type and column mapping.
	fileType := []string{"text/tab-separated-values"}
	columnMap := map[string]string{
		"condition": "column",
		"params":    "params",
		"readout":   "row",
		"score":     "score",
		"abundance": "value",
	}

	// Parse file and parameters.
	csv := columnparser.ParseCsv([]string{*file}, fileType, columnMap, true)
	imageType, parameters := parseParams(csv)

	// Convert csv to matrices
	matrices := helper.ConditionReadoutMatrix(csv, parameters.ScoreType, false, true)

	// Create folders
	helper.CreateFolders([]string{"interactive", "minimap"})

	// Generate svg for minimap
	content := svg.Heatmap(
		imageType,
		matrices,
		typedef.Annotations{},
		typedef.Markers{},
		matrices.Conditions,
		matrices.Readouts,
		true,
		parameters,
	)
	mapSvg := fmt.Sprintf("%s.svg", imageType)
	mapSvgPath := fmt.Sprintf("minimap/%s.svg", imageType)
	afero.WriteFile(fs.Instance, mapSvgPath, []byte(content), 0644)

	// Generate minimap
	svgMiniList := []string{mapSvg}
	svg.ConvertMap(svgMiniList)

	// Generate interactive file
	mapPngPath := fmt.Sprintf("minimap/%s.png", imageType)
	json := interactive.ParseHeatmap(
		imageType,
		matrices.Abundance,
		matrices.Ratio,
		matrices.Score,
		matrices.Conditions,
		matrices.Readouts,
		parameters.InvertColor,
		parameters,
		mapPngPath,
		parameters.XLabel,
		parameters.YLabel,
	)
	interactivePath := fmt.Sprintf("interactive/%s.json", imageType)
	afero.WriteFile(fs.Instance, interactivePath, []byte(json), 0644)

	// Remove minimap folder.
	// mapFolder := filepath.Join(".", "minimap")
	// fs.Instance.RemoveAll(mapFolder)
}
