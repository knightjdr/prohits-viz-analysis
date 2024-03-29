package nocluster

import (
	"github.com/knightjdr/hclust"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/normalize"
	"github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/dotplot/hierarchical"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

type orderParameters struct {
	cluster   string
	inputList []string
	names     []string
	transpose bool
}

func order(matrices *types.Matrices, analysis *types.Analysis) hierarchical.HclustData {
	orderedData := hierarchical.HclustData{
		Dendrogram: map[string][]hclust.SubCluster{
			"condition": {},
			"readout":   {},
		},
		Distance: map[string][][]float64{
			"condition": {},
			"readout":   {},
		},
		NormalizedDistance: map[string][][]float64{
			"condition": {},
			"readout":   {},
		},
		Tree: map[string]hclust.TreeLayout{
			"condition": {},
			"readout":   {},
		},
		UnsortedNames: map[string][]string{
			"condition": {},
			"readout":   {},
		},
	}

	orderDimension("condition", matrices, analysis, &orderedData)
	orderDimension("readout", matrices, analysis, &orderedData)

	return orderedData
}

func orderDimension(orderType string, matrices *types.Matrices, analysis *types.Analysis, data *hierarchical.HclustData) {
	parameters := defineOrderParameters(orderType, matrices, analysis.Settings)
	data.UnsortedNames[orderType] = parameters.names

	if parameters.cluster != "none" {
		var err error
		data.Distance[orderType] = hclust.Distance(matrices.Abundance, analysis.Settings.Distance, parameters.transpose)
		data.Dendrogram[orderType], err = hclust.Cluster(data.Distance[orderType], analysis.Settings.ClusteringMethod)
		log.CheckError(err, true)

		if analysis.Settings.ClusteringOptimize {
			data.Dendrogram[orderType] = hclust.Optimize(data.Dendrogram[orderType], data.Distance[orderType], types.HclustIgnore)
		}

		data.Tree[orderType], err = hclust.Tree(data.Dendrogram[orderType], parameters.names)
		log.CheckError(err, true)

		data.NormalizedDistance[orderType] = normalize.Matrix(data.Distance[orderType])
	} else {
		data.Tree[orderType] = hclust.TreeLayout{
			Order: checkRequestedList(analysis.Data, orderType, parameters.inputList),
		}
	}
}

func defineOrderParameters(orderType string, matrices *types.Matrices, settings types.Settings) orderParameters {
	if orderType == "condition" {
		return orderParameters{
			cluster:   settings.ConditionClustering,
			inputList: settings.ConditionList,
			names:     matrices.Conditions,
			transpose: true,
		}

	}
	return orderParameters{
		cluster:   settings.ReadoutClustering,
		inputList: settings.ReadoutList,
		names:     matrices.Readouts,
		transpose: false,
	}
}
