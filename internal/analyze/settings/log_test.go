package settings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	. "github.com/knightjdr/prohits-viz-analysis/internal/analyze/settings"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

var _ = Describe("Log", func() {
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
			Condition:          "conditionColumn",
			Control:            "controlColumn",
			Distance:           "canberra",
			Files:              []string{"/folder/file1.txt", "file2.txt"},
			LogBase:            "2",
			MinAbundance:       5.00,
			Normalization:      "total",
			PrimaryFilter:      0.01,
			Readout:            "readoutColumn",
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
			"- condition: conditionColumn\n" +
			"- readout: readoutColumn\n" +
			"- score: scoreColumn\n" +
			"- control: controlColumn\n" +
			"- readout length: readoutLengthColumn\n\n" +
			"Readout abundance transformations\n" +
			"- control subtraction was performed\n" +
			"- readout length normalization was performed\n" +
			"- condition normalization was performed using total abundance\n" +
			"- data was log-transformed with base 2\n\n" +
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
