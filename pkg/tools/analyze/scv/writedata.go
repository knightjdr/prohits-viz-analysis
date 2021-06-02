package scv

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func writeData(plots []types.CircHeatmap, legend types.CircHeatmapLegend, settings types.Settings) {
	file, err := fs.Instance.Create("other/scv-data.txt")
	log.CheckError(err, false)
	if err != nil {
		return
	}
	defer file.Close()

	var buffer strings.Builder

	writeDataHeader(&buffer, legend, settings)

	writeKnownness := getKnownessDataWriter(settings.Known)

	for _, plot := range plots {
		for _, readout := range plot.Readouts {
			buffer.WriteString(fmt.Sprintf("%s\t%s", plot.Name, readout.Label))
			for _, legendElement := range legend {
				buffer.WriteString(fmt.Sprintf("\t%.2f", readout.Segments[legendElement.Attribute]))
			}
			writeKnownness(&buffer, readout.Known)
			buffer.WriteString("\n")
		}
	}

	file.WriteString(buffer.String())
}

func writeDataHeader(buffer *strings.Builder, legend types.CircHeatmapLegend, settings types.Settings) {
	buffer.WriteString(fmt.Sprintf("%s\t%s", settings.Condition, settings.Readout))
	for _, legendElement := range legend {
		buffer.WriteString(fmt.Sprintf("\t%s", legendElement.Attribute))
	}
	if settings.Known != "" {
		buffer.WriteString(fmt.Sprintf("\tknown %s", settings.Known))
	}
	buffer.WriteString("\n")
}

func getKnownessDataWriter(known string) func(*strings.Builder, bool) {
	if known != "" {
		return func(buffer *strings.Builder, isKnown bool) {
			buffer.WriteString(fmt.Sprintf("\t%t", isKnown))
		}
	}
	return func(buffer *strings.Builder, isKnown bool) {
	}
}
