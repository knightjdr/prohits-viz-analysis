package scv

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

var _ = Describe("Create scv legend", func() {
	It("should create legend with all metric types", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("svg", 0755)

		data := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutA": {
					"Abundance":     10,
					"LogFoldChange": 5,
					"Specificity":   3,
				},
				"readoutC": {
					"Abundance":     20,
					"LogFoldChange": -10,
					"Specificity":   5,
				},
			},
			"conditionB": {
				"readoutA": {
					"Abundance":     10,
					"LogFoldChange": -7,
					"Specificity":   0.1,
				},
				"readoutB": {
					"Abundance":     20,
					"LogFoldChange": 10,
					"Specificity":   1,
				},
			},
		}
		settings := types.Settings{
			Abundance:      "Abundance",
			AbundanceCap:   50,
			MinAbundance:   2,
			OtherAbundance: []string{"LogFoldChange"},
			ProteinTissues: []string{"U2OS"},
			RnaTissues:     []string{"HEK-293", "HeLa"},
			Specificity:    true,
		}

		expectedElements := types.CircHeatmapLegend{
			{
				Attribute: "Abundance",
				Color:     "blue",
				Filter:    2,
				Max:       50,
				Min:       2,
			},
			{
				Attribute: "LogFoldChange",
				Color:     "blueRed",
				Filter:    2,
				Max:       50,
				Min:       -50,
			},
			{
				Attribute: "Specificity",
				Color:     "blue",
				Filter:    0,
				Max:       50,
				Min:       0,
			},
			{
				Attribute: "Protein expression - U2OS",
				Color:     "red",
				Filter:    0,
				Max:       7,
				Min:       0,
			},
			{
				Attribute: "RNA expression - HEK-293",
				Color:     "green",
				Filter:    0,
				Max:       50,
				Min:       0,
			},
			{
				Attribute: "RNA expression - HeLa",
				Color:     "green",
				Filter:    0,
				Max:       50,
				Min:       0,
			},
		}

		actualElements := createLegend(data, settings)
		Expect(actualElements).To(Equal(expectedElements))
	})
})
