package hierarchical

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
)

func writeMatrices(data *sortedData) {
	writeMatrix(data.matrices.Abundance, data.matrices.Conditions, data.matrices.Readouts, "other/data-transformed.txt")
	writeMatrix(data.matrices.Ratio, data.matrices.Conditions, data.matrices.Readouts, "other/data-transformed-ratios.txt")
}

func writeMatrix(matrix [][]float64, conditions, readouts []string, filename string) {
	file, err := fs.Instance.Create(filename)
	log.CheckError(err, false)
	if err != nil {
		return
	}
	defer file.Close()

	var buffer strings.Builder
	buffer.WriteString(fmt.Sprintf("%s\n", strings.Join(append([]string{""}, conditions...), "\t")))

	for i, row := range matrix {
		buffer.WriteString(readouts[i])
		for _, value := range row {
			buffer.WriteString(fmt.Sprintf("\t%0.2f", value))
		}
		buffer.WriteString("\n")
	}
	file.WriteString(buffer.String())
}
