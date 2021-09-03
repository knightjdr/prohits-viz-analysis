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
			OtherAbundance: []string{"LogFoldChange"},
			ProteinTissues: []string{"U2OS"},
			RnaTissues:     []string{"HEK-293", "HeLa"},
			Specificity:    true,
		}

		expectedElements := types.CircHeatmapLegend{
			{Attribute: "Abundance", Color: "blue", Max: 50, Min: 0},
			{Attribute: "LogFoldChange", Color: "blueRed", Max: 50, Min: -50},
			{Attribute: "Specificity", Color: "blue", Max: 50, Min: 0},
			{Attribute: "Protein expression - U2OS", Color: "red", Max: 7, Min: 0},
			{Attribute: "RNA expression - HEK-293", Color: "green", Max: 50, Min: 0},
			{Attribute: "RNA expression - HeLa", Color: "green", Max: 50, Min: 0},
		}

		actualElements := createLegend(data, settings)
		Expect(actualElements).To(Equal(expectedElements))
	})
})
