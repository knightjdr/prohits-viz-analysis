package scv

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/svg/circheatmap"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Create scv legend", func() {
	It("should create legend with all metric types", func() {
		settings := types.Settings{
			Abundance:      "abundance",
			AbundanceCap:   50,
			OtherAbundance: []string{"foldChange"},
			ProteinTissues: []string{"U2OS"},
			RnaTissues:     []string{"HEK-293", "HeLa"},
			Specificity:    true,
		}

		expectedElements := []circheatmap.LegendElement{
			{Attribute: "abundance", Color: "blue", Max: 50, Min: 0},
			{Attribute: "foldChange", Color: "blue", Max: 50, Min: 0},
			{Attribute: "Specificity", Color: "blue", Max: 50, Min: 0},
			{Attribute: "Protein expression - U2OS", Color: "red", Max: 7, Min: 0},
			{Attribute: "RNA expression - HEK-293", Color: "green", Max: 50, Min: 0},
			{Attribute: "RNA expression - HeLa", Color: "green", Max: 50, Min: 0},
		}

		actualElements := createLegend(settings)
		Expect(actualElements).To(Equal(expectedElements))
	})
})
