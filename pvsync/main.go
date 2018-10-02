// Package main takes a stringified JSON and generates a minimap.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/knightjdr/prohits-viz-analysis/interactive"
	"github.com/knightjdr/prohits-viz-analysis/parse"
	"github.com/knightjdr/prohits-viz-analysis/svg"
)

func main() {
	// Parse flags.
	jsonFile := flag.String("json", "", "JSON file")
	flag.Parse()

	data, err := parse.HeatmapJSON(*jsonFile)
	if err != nil {
		os.Exit(1)
	}

	// Format dataset for svg creator.
	abundance, ratios, scores := parse.FormatMatrix(data)

	// Format parameters for svg.
	parameters := FormatParams(data)

	// Creat dummy row and column names.
	dummyColumns, dummyRows := Dummy(len(data.Rows[0].Data), len(data.Rows))

	// Create svg.
	Map(
		data.ImageType,
		abundance,
		ratios,
		scores,
		dummyColumns,
		dummyRows,
		parameters,
	)

	// Create minimap.
	svg.ConvertMap([]string{fmt.Sprintf("%s.svg", data.ImageType)})

	// Generate URL.
	url := interactive.Pngurl(fmt.Sprintf("minimap/%s.png", data.ImageType))
	fmt.Println(url)
}
