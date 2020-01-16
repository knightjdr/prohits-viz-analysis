package settings

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

func logFiltering(messages *strings.Builder, settings types.Settings) {
	messages.WriteString(
		fmt.Sprintf(
			"Filtering\n"+
				"- minimum %s requirement: %d\n"+
				"- parsimonius %s inclusion: %t\n\n",
			settings.Condition,
			settings.MinConditions,
			settings.Readout,
			settings.ParsimoniousReadoutFiltering,
		),
	)
}

func logBiclustering(messages *strings.Builder, settings types.Settings) {
	if settings.Clustering == "biclustering" {
		if settings.BiclusteringApprox {
			messages.WriteString("- approximate biclustering was performed\n")
		} else {
			messages.WriteString("- biclustering was performed\n")
		}
	}
}

func logHierarchical(messages *strings.Builder, settings types.Settings) {
	if settings.Clustering == "hierarchical" {
		messages.WriteString(
			fmt.Sprintf(
				"- hierarchical clustering was performed\n- distance metric: %s\n- linkage method: %s\n",
				settings.Distance,
				settings.ClusteringMethod,
			),
		)
	}
}

func logNoClustering(messages *strings.Builder, settings types.Settings) {
	if settings.Clustering == "none" {
		if settings.ConditionClustering == "none" && settings.ReadoutClustering == "hierarchical" {
			messages.WriteString(
				fmt.Sprintf(
					"- %s were hierarchically clustered\n- distance metric: %s\n- linkage method: %s\n",
					settings.Readout,
					settings.Distance,
					settings.ClusteringMethod,
				),
			)
		} else if settings.ReadoutClustering == "none" && settings.ConditionClustering == "hierarchical" {
			messages.WriteString(
				fmt.Sprintf(
					"- %s were hierarchically clustered\n- distance metric: %s\n- linkage method: %s\n",
					settings.Condition,
					settings.Distance,
					settings.ClusteringMethod,
				),
			)
		} else {
			messages.WriteString("- no clustering was performed\n")
		}
	}
}

func logClusteringOptimization(messages *strings.Builder, settings types.Settings) {
	if settings.Clustering == "hierarchical" || settings.ConditionClustering == "hierarchical" || settings.ReadoutClustering == "hierarchical" {
		if settings.ClusteringOptimize {
			messages.WriteString("- leaf clusters were optimized\n")
		} else {
			messages.WriteString("- leaf clusters were not optimized\n")
		}
	}
}
