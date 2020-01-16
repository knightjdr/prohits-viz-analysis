package hierarchical

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/data/filter"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	"github.com/knightjdr/prohits-viz-analysis/pkg/stats"
)

// CreateCytoscape files.
func CreateCytoscape(fileData []map[string]string, data *SortedData, settings types.Settings) {
	writeDistanceCytoscape(data.ConditionDist, data.Matrices.Conditions, data.Matrices.Conditions, settings.Condition)
	writeDistanceCytoscape(data.ReadoutDist, data.Matrices.Readouts, data.Matrices.Readouts, settings.Readout)
	writeFileDataCytoscape(fileData, settings)
}

func writeDistanceCytoscape(matrix [][]float64, source, target []string, filehandle string) {
	if len(matrix) > 0 {
		file, err := fs.Instance.Create(fmt.Sprintf("cytoscape/%[1]s-%[1]s-cytoscape.txt", filehandle))
		log.CheckError(err, false)
		if err != nil {
			return
		}
		defer file.Close()

		var buffer strings.Builder
		buffer.WriteString("source\ttarget\tdistance\n")

		noSource := len(source)
		for i, row := range matrix {
			for j := i + 1; j < noSource; j++ {
				value := float.RemoveTrailingZeros(row[j])
				buffer.WriteString(fmt.Sprintf("%s\t%s\t%s\n", source[i], target[j], value))
			}
		}
		file.WriteString(buffer.String())
	}
}

func writeFileDataCytoscape(data []map[string]string, settings types.Settings) {
	file, err := fs.Instance.Create(fmt.Sprintf("cytoscape/%s-%s-cytoscape.txt", settings.Condition, settings.Readout))
	log.CheckError(err, false)
	if err != nil {
		return
	}
	defer file.Close()

	var buffer strings.Builder
	buffer.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", settings.Condition, settings.Readout, settings.Abundance, settings.Score))

	filter := filter.DefineScoreFilter(settings)
	for _, row := range data {
		score, _ := strconv.ParseFloat(row["score"], 64)
		if filter(score) {
			abundance := stats.MeanString(strings.Split(row["abundance"], "|"))
			buffer.WriteString(fmt.Sprintf("%s\t%s\t%0.2f\t%s\n", row["condition"], row["readout"], abundance, row["score"]))
		}
	}

	file.WriteString(buffer.String())
}
