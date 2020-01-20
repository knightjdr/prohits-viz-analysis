// Package heatmap creates an interactive heatmap/dotplot file.
package heatmap

import (
	"fmt"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/spf13/afero"
)

// Create a file for the interactive heatmap/dotplot viewer.
func Create(data *Data) {
	file, err := fs.Instance.Create(data.Filename)
	log.CheckError(err, true)
	defer file.Close()

	writeString := write(file)

	writeString("{\n")
	writeString(fmt.Sprintf("\t%s,\n", parseParameters(data.AnalysisType, data.Parameters)))
	writeString(fmt.Sprintf("\t%s,\n", parseSettings(data.Settings)))
	writeString(fmt.Sprintf("\t%s,\n", parseColumns(data.Matrices.Conditions)))
	parseRows(data, writeString)
	writeString(fmt.Sprintf("\t%s\n", parseMinimap(data.Minimap)))
	writeString("}\n")
}

func write(file afero.File) func(string) {
	return func(str string) {
		file.Write([]byte(str))
	}
}
