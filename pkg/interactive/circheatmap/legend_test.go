package circheatmap

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Legend scv", func() {
	It("should parse legend and return string", func() {
		legend := types.CircHeatmapLegend{
			{Attribute: "abundance", Color: "blue", Max: 50, Min: 0},
			{Attribute: "foldChange", Color: "blue", Max: 50, Min: 0},
			{Attribute: "Specificity", Color: "blue", Max: 50, Min: 0},
			{Attribute: "Protein expression - U2OS", Color: "red", Max: 7, Min: 0},
			{Attribute: "RNA expression - HEK-293", Color: "green", Max: 50, Min: 0},
			{Attribute: "RNA expression - HeLa", Color: "green", Max: 50, Min: 0},
		}

		expectedString, _ := json.Marshal(legend)
		expected := fmt.Sprintf("\"circles\": {\"order\": %s}", expectedString)
		Expect(parseLegend(legend)).To(Equal(expected))
	})
})
