// Package scatter converts an interactive file from ProHits-viz V1 to V2 format
package scatter

import (
	"path"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

// Convert a heatmap or dotplot file to json format.
func Convert(filename string) {
	format := determineFormat(filename)

	files.CreateFolders([]string{"interactive"})
	fileid := strings.Split(path.Base(filename), ".txt")[0]

	legend := []map[string]string{}
	plots := []types.ScatterPlot{}
	settings := types.Settings{}
	if format == 1 {
		plots, settings, legend = readFormat1(filename)
	}
	if format == 2 {
		plots, settings, legend = readFormat2(filename)
	}

	createInteractive(plots, settings, legend, fileid)
}

func determineFormat(filename string) int {
	file, err := fs.Instance.Open(filename)
	log.CheckError(err, true)
	defer file.Close()
	reader := createReader(file)

	header, err := reader.Read()
	log.CheckError(err, true)

	var format int
	if header[0] == "entry" {
		format = 1
	}
	if header[0] == "details:" {
		format = 2
	}

	return format
}
