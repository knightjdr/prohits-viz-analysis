package scv

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Format data to plot struct", func() {
	It("should create struct when knowness requested", func() {
		data := map[string]map[string]map[string]float64{
			"conditionB": {
				"readoutY": {"abundance": 2},
				"readoutZ": {"abundance": 6},
			},
			"conditionA": {
				"readoutA": {"abundance": 1},
				"readoutX": {"abundance": 2},
				"readoutY": {"abundance": 4},
			},
		}
		known := map[string]map[string]bool{
			"conditionA": {"readoutA": true, "readoutX": true, "readoutY": false},
			"conditionB": {"readoutY": true, "readoutZ": false},
		}
		legend := types.CircHeatmapLegend{
			{Attribute: "abundance", Color: "blue", Max: 50, Min: 0},
		}
		settings := types.Settings{
			Known: "interaction",
		}

		expected := []types.CircHeatmap{
			{
				Name: "conditionA",
				Readouts: []types.CircHeatmapReadout{
					{
						Known: true,
						Label: "readoutX",
						Segments: map[string]types.RoundedSegment{
							"abundance": 2,
						},
					},
					{
						Known: true,
						Label: "readoutA",
						Segments: map[string]types.RoundedSegment{
							"abundance": 1,
						},
					},
					{
						Known: false,
						Label: "readoutY",
						Segments: map[string]types.RoundedSegment{
							"abundance": 4,
						},
					},
				},
			},
			{
				Name: "conditionB",
				Readouts: []types.CircHeatmapReadout{
					{
						Known: true,
						Label: "readoutY",
						Segments: map[string]types.RoundedSegment{
							"abundance": 2,
						},
					},
					{
						Known: false,
						Label: "readoutZ",
						Segments: map[string]types.RoundedSegment{
							"abundance": 6,
						},
					},
				},
			},
		}

		Expect(createPlotData(data, known, legend, settings)).To(Equal(expected))
	})

	It("should create struct when knowness not requested", func() {
		data := map[string]map[string]map[string]float64{
			"conditionB": {
				"readoutY": {"abundance": 2},
				"readoutZ": {"abundance": 6},
			},
			"conditionA": {
				"readoutA": {"abundance": 5},
				"readoutX": {"abundance": 8},
				"readoutY": {"abundance": 6},
			},
		}
		var known map[string]map[string]bool = nil
		legend := types.CircHeatmapLegend{
			{Attribute: "abundance", Color: "blue", Max: 50, Min: 0},
		}
		settings := types.Settings{
			Known: "",
		}

		expected := []types.CircHeatmap{
			{
				Name: "conditionA",
				Readouts: []types.CircHeatmapReadout{
					{
						Known: false,
						Label: "readoutX",
						Segments: map[string]types.RoundedSegment{
							"abundance": 8,
						},
					},
					{
						Known: false,
						Label: "readoutY",
						Segments: map[string]types.RoundedSegment{
							"abundance": 6,
						},
					},
					{
						Known: false,
						Label: "readoutA",
						Segments: map[string]types.RoundedSegment{
							"abundance": 5,
						},
					},
				},
			},
			{
				Name: "conditionB",
				Readouts: []types.CircHeatmapReadout{
					{
						Known: false,
						Label: "readoutZ",
						Segments: map[string]types.RoundedSegment{
							"abundance": 6,
						},
					},
					{
						Known: false,
						Label: "readoutY",
						Segments: map[string]types.RoundedSegment{
							"abundance": 2,
						},
					},
				},
			},
		}

		Expect(createPlotData(data, known, legend, settings)).To(Equal(expected))
	})
})
