package scatter

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Legend", func() {
	It("should parse legend and return string", func() {
		plots := []types.ScatterPlot{
			{
				Labels: types.ScatterAxesLabels{X: "conditionX", Y: "conditionY"},
				Name:   "condition-condition",
				Points: []types.ScatterPoint{
					{Label: "readoutA", X: 1, Y: 3, Color: "#0066cc"},
					{Label: "readoutB", X: 0, Y: 4, Color: "#99ccff"},
					{Label: "readoutC", X: 2, Y: 0, Color: "#99ccff"},
				},
			},
		}

		expectedString, _ := json.Marshal(plots)
		expected := fmt.Sprintf("\"plots\": %s", expectedString)
		Expect(parsePlots(plots)).To(Equal(expected))
	})
})
