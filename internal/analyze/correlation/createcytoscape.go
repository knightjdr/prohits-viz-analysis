package correlation

import (
	"bufio"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
)

func createCytoscapeFiles(conditionData, readoutData *correlationData, settings types.Settings) {
	createCytoscape(conditionData, settings.CytoscapeCutoff, settings.Condition)
	createCytoscape(readoutData, settings.CytoscapeCutoff, settings.Readout)
}

func createCytoscape(data *correlationData, cutoff float64, label string) {
	file, err := fs.Instance.Create(fmt.Sprintf("cytoscape/%[1]s-%[1]s.txt", label))
	log.CheckError(err, true)
	defer file.Close()

	buffer := bufio.NewWriter(file)
	buffer.WriteString(fmt.Sprintf("%[1]s\t%[1]s\tcorrelation\n", label))

	n := len(data.matrix)
	for i, row := range data.matrix {
		for j := i + 1; j < n; j++ {
			correlation := row[j]
			if correlation >= cutoff {
				corrString := float.RemoveTrailingZeros(correlation)
				buffer.WriteString(fmt.Sprintf("%s\t%s\t%s\n", data.sortedLabels[i], data.sortedLabels[j], corrString))
			}
		}
	}

	buffer.Flush()
}
