// Package scatter creates an interactive scatter plot file.
package scatter

import (
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/spf13/afero"
)

// Create a file for the interactive scatter plot viewer.
func Create(data *Data) {
	file, err := fs.Instance.Create(data.Filename)
	log.CheckError(err, true)
	defer file.Close()

	writeString := write(file)

	writeString("{\n")
	writeString(fmt.Sprintf("\t%s,\n", parseLegend(data.Legend)))
	writeString(fmt.Sprintf("\t%s,\n", parseParameters(data.AnalysisType, data.Parameters)))
	writeString(fmt.Sprintf("\t%s,\n", parseSettings(data.Settings)))
	writeString(fmt.Sprintf("\t%s\n", parsePlots(data.Plots)))
	writeString("}\n")
}

func write(file afero.File) func(string) {
	return func(str string) {
		file.Write([]byte(str))
	}
}
