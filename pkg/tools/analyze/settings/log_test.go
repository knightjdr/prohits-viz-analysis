package settings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	. "github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/settings"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Log", func() {
	It("should log correlation settings", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		settings := types.Settings{
			Abundance:                    "abundanceColumn",
			Clustering:                   "hierarchical",
			ClusteringMethod:             "complete",
			ClusteringOptimize:           true,
			Condition:                    "Bait",
			ConditionAbundanceFilter:     5,
			ConditionScoreFilter:         0.01,
			Control:                      "controlColumn",
			Correlation:                  "pearson",
			CytoscapeCutoff:              0.7,
			Distance:                     "canberra",
			Files:                        []string{"/folder/file1.txt", "file2.txt"},
			IgnoreSourceTargetMatches:    true,
			LogBase:                      "2",
			MinConditions:                2,
			MockConditionAbundance:       true,
			Normalization:                "total",
			ParsimoniousReadoutFiltering: true,
			ReadoutAbundanceFilter:       10,
			ReadoutScoreFilter:           0.05,
			Readout:                      "Prey",
			ReadoutLength:                "readoutLengthColumn",
			Score:                        "scoreColumn",
			ScoreType:                    "lte",
			Type:                         "correlation",
			UseReplicates:                true,
		}

		expected := "Analysis type: correlation\n\n" +
			"Files used\n" +
			"- file1.txt\n" +
			"- file2.txt\n\n" +
			"Columns used\n" +
			"- abundance: abundanceColumn\n" +
			"- condition: Bait\n" +
			"- readout: Prey\n" +
			"- score: scoreColumn\n" +
			"- control: controlColumn\n" +
			"- readout length: readoutLengthColumn\n\n" +
			"Prey abundance transformations\n" +
			"- control subtraction was performed\n" +
			"- Prey length normalization was performed\n" +
			"- Bait normalization was performed using total abundance\n" +
			"- data was log-transformed with base 2\n" +
			"- abundance values were mocked for Bait(s) with missing values\n\n" +
			"Filtering\n" +
			"- minimum Bait requirement: 2\n" +
			"- parsimonius Prey inclusion: true\n\n" +
			"Abundance\n" +
			"- minimum abundance for Bait correlation: 5\n" +
			"- minimum abundance for Prey correlation: 10\n\n" +
			"Scoring\n" +
			"- smaller scores are better\n" +
			"- score filter for Bait correlation: 0.01\n" +
			"- score filter for Prey correlation: 0.05\n\n" +
			"Correlation\n" +
			"- correlation method: pearson\n" +
			"- treat replicates as separate data points: true\n" +
			"- ignore source genes in pairwise correlations: true\n" +
			"- cytoscape cutoff: 0.7\n\n" +
			"Clustering\n" +
			"- hierarchical clustering was performed\n" +
			"- distance metric: canberra\n" +
			"- linkage method: complete\n" +
			"- leaf clusters were optimized\n\n"

		Log(settings)
		actual, _ := afero.ReadFile(fs.Instance, "log.txt")
		Expect(string(actual)).To(Equal(expected))
	})

	It("should log dotplot settings", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		settings := types.Settings{
			Abundance:          "abundanceColumn",
			AbundanceCap:       10.00,
			Clustering:         "hierarchical",
			ClusteringMethod:   "complete",
			ClusteringOptimize: true,
			Condition:          "Bait",
			Control:            "controlColumn",
			Distance:           "canberra",
			Files:              []string{"/folder/file1.txt", "file2.txt"},
			LogBase:            "2",
			MinAbundance:       5.00,
			Normalization:      "total",
			PrimaryFilter:      0.01,
			Readout:            "Prey",
			ReadoutLength:      "readoutLengthColumn",
			Score:              "scoreColumn",
			ScoreType:          "lte",
			SecondaryFilter:    0.05,
			Type:               "dotplot",
		}

		expected := "Analysis type: dotplot\n\n" +
			"Files used\n" +
			"- file1.txt\n" +
			"- file2.txt\n\n" +
			"Columns used\n" +
			"- abundance: abundanceColumn\n" +
			"- condition: Bait\n" +
			"- readout: Prey\n" +
			"- score: scoreColumn\n" +
			"- control: controlColumn\n" +
			"- readout length: readoutLengthColumn\n\n" +
			"Prey abundance transformations\n" +
			"- control subtraction was performed\n" +
			"- Prey length normalization was performed\n" +
			"- Bait normalization was performed using total abundance\n" +
			"- data was log-transformed with base 2\n\n" +
			"Filtering\n" +
			"- minimum Bait requirement: 0\n" +
			"- parsimonius Prey inclusion: false\n\n" +
			"Abundance\n" +
			"- minimum abundance required: 5\n" +
			"- abundances were capped at 10 for visualization\n\n" +
			"Scoring\n" +
			"- smaller scores are better\n" +
			"- primary filter: 0.01\n" +
			"- secondary filter: 0.05\n\n" +
			"Clustering\n" +
			"- hierarchical clustering was performed\n" +
			"- distance metric: canberra\n" +
			"- linkage method: complete\n" +
			"- leaf clusters were optimized\n\n"

		Log(settings)
		actual, _ := afero.ReadFile(fs.Instance, "log.txt")
		Expect(string(actual)).To(Equal(expected))
	})
})
