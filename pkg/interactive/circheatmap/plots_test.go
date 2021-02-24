package circheatmap

import (
	"encoding/json"
	"fmt"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plots scv", func() {
	It("should parse plot and return string", func() {
		plots := []types.CircHeatmap{
			{
				Name: "conditionA",
				Readouts: []types.CircHeatmapReadout{
					{
						Known: true,
						Label: "readoutX",
						Segments: map[string]types.RoundedSegment{
							"attribute1": 1,
							"attribute2": 2,
							"attribute3": 3,
						},
					},
					{
						Known: false,
						Label: "readoutY",
						Segments: map[string]types.RoundedSegment{
							"attribute1": 4,
							"attribute2": 5,
							"attribute3": 6,
						},
					},
				},
			},
		}

		expectedString, _ := json.Marshal(plots)
		expected := fmt.Sprintf("\"plots\": %s", expectedString)
		Expect(parsePlots(plots)).To(Equal(expected))
	})
})
