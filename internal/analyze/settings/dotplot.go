package settings

import (
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
)

func logDotplotSettings(messages *strings.Builder, settings types.Settings) {
	logDotplotAbundance(messages, settings)
	logDotplotScoring(messages, settings)
	logDotplotClustering(messages, settings)
}

func logDotplotAbundance(messages *strings.Builder, settings types.Settings) {
	messages.WriteString(fmt.Sprintf("Abundance\n- minimum abundance required: %s\n", float.RemoveTrailingZeros(settings.MinAbundance)))
	if settings.AbundanceCap > 0 {
		messages.WriteString(fmt.Sprintf("- abundances were capped at %s for visualization\n", float.RemoveTrailingZeros(settings.AbundanceCap)))
	}
	messages.WriteString("\n")
}

func logDotplotScoring(messages *strings.Builder, settings types.Settings) {
	messages.WriteString("Scoring\n")

	if settings.ScoreType == "gte" {
		messages.WriteString("- larger scores are better\n")
	} else {
		messages.WriteString("- smaller scores are better\n")
	}

	messages.WriteString(fmt.Sprintf("- primary filter: %s\n", float.RemoveTrailingZeros(settings.PrimaryFilter)))

	if settings.SecondaryFilter != 0 {
		messages.WriteString(fmt.Sprintf("- secondary filter: %s\n", float.RemoveTrailingZeros(settings.SecondaryFilter)))
	}

	messages.WriteString("\n")
}

func logDotplotClustering(messages *strings.Builder, settings types.Settings) {
	messages.WriteString("Clustering\n")
	logBiclustering(messages, settings)
	logHierarchical(messages, settings)
	logNoClustering(messages, settings)
	logClusteringOptimization(messages, settings)
	messages.WriteString("\n")
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
					"- readouts were hierarchically clustered\n- distance metric: %s\n- linkage method: %s\n",
					settings.Distance,
					settings.ClusteringMethod,
				),
			)
		} else if settings.ReadoutClustering == "none" && settings.ConditionClustering == "hierarchical" {
			messages.WriteString(
				fmt.Sprintf(
					"- conditions were hierarchically clustered\n- distance metric: %s\n- linkage method: %s\n",
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
