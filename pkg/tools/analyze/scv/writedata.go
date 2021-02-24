package scv

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

func writeData(data map[string]map[string]map[string]float64, known map[string]map[string]bool, legend types.CircHeatmapLegend, settings types.Settings) {
	file, err := fs.Instance.Create("other/scv-data.txt")
	log.CheckError(err, false)
	if err != nil {
		return
	}
	defer file.Close()

	var buffer strings.Builder

	writeDataHeader(&buffer, legend, settings)

	writeKnownness := getKnownessDataWriter(settings.Known, known)

	conditions := getAndSortConditions(data)
	for _, condition := range conditions {
		readouts := getAndSortReadouts(data[condition])
		for _, readout := range readouts {
			buffer.WriteString(fmt.Sprintf("%s\t%s", condition, readout))
			for _, legendElement := range legend {
				buffer.WriteString(fmt.Sprintf("\t%.2f", data[condition][readout][legendElement.Attribute]))
			}
			writeKnownness(&buffer, condition, readout)
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

func getKnownessDataWriter(known string, knownMap map[string]map[string]bool) func(*strings.Builder, string, string) {
	if known != "" && knownMap != nil {
		return func(buffer *strings.Builder, condition string, readout string) {
			buffer.WriteString(fmt.Sprintf("\t%t", knownMap[condition][readout]))
		}
	}
	return func(buffer *strings.Builder, condition string, readout string) {
	}
}
