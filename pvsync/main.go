// Package main takes a stringified JSON and generates a minimap.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/interactive"
	"github.com/knightjdr/prohits-viz-analysis/parse"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
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
	heatmap := parse.FormatMatrix(data)

	// Format parameters for svg.
	parameters := FormatParams(data)

	// Creat dummy row and column names.
	dummyColumns, dummyRows := Dummy(len(data.Rows[0].Data), len(data.Rows))

	// Create svg.
	content := svg.Heatmap(
		data.ImageType,
		heatmap,
		typedef.Annotations{},
		typedef.Markers{},
		dummyColumns,
		dummyRows,
		true,
		parameters,
	)
	filename := fmt.Sprintf("minimap/%s.svg", data.ImageType)
	afero.WriteFile(fs.Instance, filename, []byte(content), 0644)

	// Create minimap.
	svg.ConvertMap([]string{fmt.Sprintf("%s.svg", data.ImageType)})

	// Generate URL.
	url := interactive.Pngurl(fmt.Sprintf("minimap/%s.png", data.ImageType))
	fmt.Println(url)
}
