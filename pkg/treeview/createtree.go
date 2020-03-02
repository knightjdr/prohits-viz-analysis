package treeview

import (
	"fmt"
	"strings"

	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/math"
	"github.com/spf13/afero"
)

func createATR(data Data) {
	if len(data.Trees.Column) > 0 {
		file, err := fs.Instance.Create(fmt.Sprintf("%s.atr", data.Filename))
		log.CheckError(err, true)
		defer file.Close()

		label := getTreeLabeler(data.Names.Columns, data.Names.UnsortedColumns, "ARRY")
		writeTreeNodes(file, data.Trees.Column, label)
	}
}

func createGTR(data Data) {
	if len(data.Trees.Row) > 0 {
		file, err := fs.Instance.Create(fmt.Sprintf("%s.gtr", data.Filename))
		log.CheckError(err, true)
		defer file.Close()

		label := getTreeLabeler(data.Names.Rows, data.Names.UnsortedRows, "GENE")
		writeTreeNodes(file, data.Trees.Row, label)
	}
}

func writeTreeNodes(file afero.File, dendrogram []hclust.SubCluster, label func(int) string) {
	var buffer strings.Builder

	correlations := convertDistanceToCorrelation(dendrogram)

	for nodeIndex, cluster := range dendrogram {
		aName := label(cluster.Leafa)
		bName := label(cluster.Leafb)
		buffer.WriteString(fmt.Sprintf("NODE%dX\t%s\t%s\t%0.5f\n", nodeIndex+1, aName, bName, correlations[nodeIndex]))
	}

	file.WriteString(buffer.String())
}

func convertDistanceToCorrelation(dendrogram []hclust.SubCluster) []float64 {
	correlation := make([]float64, len(dendrogram))

	// Get branch heights and add one to prevent division by zero in next step.
	heights := hclust.GetNodeHeight(dendrogram)
	for i := range heights {
		heights[i]++
	}

	min := math.MinSliceFloat(heights)
	for i, height := range heights {
		correlation[i] = min / height
	}

	return correlation
}
