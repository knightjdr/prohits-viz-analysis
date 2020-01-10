package settings

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log dotplot settings", func() {
	It("should log biclustering settings", func() {
		settings := types.Settings{
			BiclusteringApprox: true,
			Clustering:         "biclustering",
		}

		expected := "Clustering\n" +
			"- approximate biclustering was performed\n\n"

		var messages strings.Builder
		logDotplotSettings(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log hierarchical clustering settings", func() {
		settings := types.Settings{
			Clustering:         "hierarchical",
			ClusteringMethod:   "complete",
			ClusteringOptimize: true,
			Distance:           "canberra",
		}

		expected := "Clustering\n" +
			"- hierarchical clustering was performed\n" +
			"- distance metric: canberra\n" +
			"- linkage method: complete\n" +
			"- leaf clusters were optimized\n\n"

		var messages strings.Builder
		logDotplotSettings(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log no clustering settings", func() {
		settings := types.Settings{
			Clustering:          "none",
			ConditionClustering: "none",
			ReadoutClustering:   "none",
		}

		expected := "Clustering\n" +
			"- no clustering was performed\n\n"

		var messages strings.Builder
		logDotplotSettings(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log biclustering settings", func() {
	It("should log nothing when clustering is not set to biclustering", func() {
		settings := types.Settings{
			Clustering: "hierarchical",
		}

		expected := ""

		var messages strings.Builder
		logBiclustering(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log settings", func() {
		settings := types.Settings{
			BiclusteringApprox: false,
			Clustering:         "biclustering",
		}

		expected := "- biclustering was performed\n"

		var messages strings.Builder
		logBiclustering(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log settings for approximation", func() {
		settings := types.Settings{
			BiclusteringApprox: true,
			Clustering:         "biclustering",
		}

		expected := "- approximate biclustering was performed\n"

		var messages strings.Builder
		logBiclustering(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log hierarchical clustering settings", func() {
	It("should log nothing when clustering is not set to hierarchical", func() {
		settings := types.Settings{
			Clustering: "biclustering",
		}

		expected := ""

		var messages strings.Builder
		logHierarchical(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log settings", func() {
		settings := types.Settings{
			Clustering:       "hierarchical",
			ClusteringMethod: "complete",
			Distance:         "canberra",
		}

		expected := "- hierarchical clustering was performed\n" +
			"- distance metric: canberra\n" +
			"- linkage method: complete\n"

		var messages strings.Builder
		logHierarchical(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log no clustering settings", func() {
	It("should log nothing when clustering is not selected", func() {
		settings := types.Settings{
			Clustering: "hierarchical",
		}

		expected := ""

		var messages strings.Builder
		logNoClustering(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log that no clustering was performed when conditions and readouts are specified", func() {
		settings := types.Settings{
			Clustering:          "none",
			ConditionClustering: "none",
			ReadoutClustering:   "none",
		}

		expected := "- no clustering was performed\n"

		var messages strings.Builder
		logNoClustering(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log that conditions were clustered", func() {
		settings := types.Settings{
			Clustering:          "none",
			ClusteringMethod:    "complete",
			ConditionClustering: "hierarchical",
			Distance:            "canberra",
			ReadoutClustering:   "none",
		}

		expected := "- conditions were hierarchically clustered\n" +
			"- distance metric: canberra\n" +
			"- linkage method: complete\n"

		var messages strings.Builder
		logNoClustering(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	It("should log that readouts were clustered", func() {
		settings := types.Settings{
			Clustering:          "none",
			ClusteringMethod:    "complete",
			ConditionClustering: "none",
			Distance:            "canberra",
			ReadoutClustering:   "hierarchical",
		}

		expected := "- readouts were hierarchically clustered\n" +
			"- distance metric: canberra\n" +
			"- linkage method: complete\n"

		var messages strings.Builder
		logNoClustering(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})
})

var _ = Describe("Log clustering optimization", func() {
	It("should log nothing when hierarchical clustering is not performed", func() {
		settings := types.Settings{
			Clustering: "biclustering",
		}

		expected := ""

		var messages strings.Builder
		logClusteringOptimization(&messages, settings)
		Expect(messages.String()).To(Equal(expected))
	})

	Describe("all data hierarchically clustered", func() {
		It("should log that optimization was performed", func() {
			settings := types.Settings{
				Clustering:         "hierarchical",
				ClusteringOptimize: true,
			}

			expected := "- leaf clusters were optimized\n"

			var messages strings.Builder
			logClusteringOptimization(&messages, settings)
			Expect(messages.String()).To(Equal(expected))
		})

		It("should log that optimization was not performed", func() {
			settings := types.Settings{
				Clustering:         "hierarchical",
				ClusteringOptimize: false,
			}

			expected := "- leaf clusters were not optimized\n"

			var messages strings.Builder
			logClusteringOptimization(&messages, settings)
			Expect(messages.String()).To(Equal(expected))
		})
	})

	Describe("conditions hierarchically clustered", func() {
		It("should log that optimization was performed", func() {
			settings := types.Settings{
				Clustering:          "none",
				ClusteringOptimize:  true,
				ConditionClustering: "hierarchical",
			}

			expected := "- leaf clusters were optimized\n"

			var messages strings.Builder
			logClusteringOptimization(&messages, settings)
			Expect(messages.String()).To(Equal(expected))
		})
	})

	Describe("conditions hierarchically clustered", func() {
		It("should log that optimization was performed", func() {
			settings := types.Settings{
				Clustering:         "none",
				ClusteringOptimize: true,
				ReadoutClustering:  "hierarchical",
			}

			expected := "- leaf clusters were optimized\n"

			var messages strings.Builder
			logClusteringOptimization(&messages, settings)
			Expect(messages.String()).To(Equal(expected))
		})
	})
})
