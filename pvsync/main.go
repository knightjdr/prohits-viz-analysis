// Package main takes a stringified JSON and generates a minimap.
package main

import (
	"fmt"
	"os"

	"github.com/knightjdr/prohits-viz-analysis/interactive"
	"github.com/knightjdr/prohits-viz-analysis/svg"
	"github.com/knightjdr/prohits-viz-analysis/tool/dotplot"
)

func main() {
	// Parse flags.
	data, err := ParseJSON()
	if err != nil {
		os.Exit(1)
	}

	// Format dataset for svg creator.
	abundance, ratios, scores := FormatData(data)

	// Format parameters for svg.
	params := FormatParams(data)

	// Creat dummy row and column names.
	dummyColumns, dummyRows := Dummy(len(data.Rows[0]), len(data.Rows))

	// Create svg.
	image := "heatmap"
	if data.ImageType == "dotplot" {
		image = "dotplot"
		dotplot.SvgDotplot(
			abundance,
			ratios,
			scores,
			dummyColumns,
			dummyRows,
			data.Invert,
			params,
		)
	} else {
		dotplot.SvgHeatmap(
			abundance,
			dummyColumns,
			dummyRows,
			data.FillColor,
			data.MaximumAbundance,
			data.Invert,
		)
	}

	// Create minimap.
	svg.ConvertMap([]string{fmt.Sprintf("%s.svg", image)})

	// Generate URL.
	url := interactive.Pngurl(fmt.Sprintf("minimap/%s.png", image))
	fmt.Println(url)
}
