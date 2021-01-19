package scatter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plots", func() {
	It("should parse plot and return string", func() {
		plots := []types.ScatterPlot{
			{
				Labels: types.ScatterAxesLabels{X: "conditionX", Y: "conditionY"},
				Name:   "condition-condition",
				Points: []types.ScatterPoint{
					{Label: "readoutA", X: 1.020001, Y: 3.00, Color: "#0066cc"},
					{Label: "readoutB", X: 0.00, Y: 4.00, Color: "#99ccff"},
					{Label: "readoutC", X: 2.00, Y: 0.00, Color: "#99ccff"},
				},
			},
		}

		expectedString := "\"plots\": [{\"labels\":{\"x\":\"conditionX\",\"y\":\"conditionY\"},\"name\":\"condition-condition\",\"points\":[{\"color\":\"#0066cc\",\"label\":\"readoutA\",\"x\":1.02,\"y\":3.00},{\"color\":\"#99ccff\",\"label\":\"readoutB\",\"x\":0.00,\"y\":4.00},{\"color\":\"#99ccff\",\"label\":\"readoutC\",\"x\":2.00,\"y\":0.00}]}]"
		Expect(parsePlots(plots)).To(Equal(expectedString))
	})
})
