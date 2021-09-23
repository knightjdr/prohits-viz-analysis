package specificity

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Format specificity data", func() {
	It("should format data for plot", func() {
		data := map[string]map[string]map[string]float64{
			"a": {
				"x": {"abundance": 10, "score": 0.01, "specificity": 0.67},
				"y": {"abundance": 20, "score": 0.01, "specificity": 2.67},
			},
			"b": {
				"x": {"abundance": 30, "score": 0, "specificity": 6},
			},
			"c": {
				"y": {"abundance": 15, "score": 0.02, "specificity": 1.5},
				"z": {"abundance": 25, "score": 0.01, "specificity": math.Inf(1)},
			},
		}
		settings := types.Settings{
			PrimaryFilter: 0.01,
			ScoreType:     "lte",
		}

		expected := map[string][]types.ScatterPoint{
			"a": {
				{Color: "#dfcd06", Label: "x", X: 10, Y: 0.67},
				{Color: "#dfcd06", Label: "y", X: 20, Y: 2.67},
			},
			"b": {
				{Color: "#dfcd06", Label: "x", X: 30, Y: 6},
			},
			"c": {
				{Color: "#6e97ff", Label: "z", X: 25, Y: 100},
			},
		}
		Expect(formatDataForPlot(data, settings)).To(Equal(expected))
	})

	It("should format data for plot included negative values", func() {
		data := map[string]map[string]map[string]float64{
			"a": {
				"x": {"abundance": -10, "score": 0.01, "specificity": -0.67},
				"y": {"abundance": 20, "score": 0.01, "specificity": 2.67},
			},
			"b": {
				"x": {"abundance": -30, "score": 0, "specificity": 6},
			},
			"c": {
				"y": {"abundance": 15, "score": 0.02, "specificity": 1.5},
				"z": {"abundance": 25, "score": 0.01, "specificity": math.Inf(1)},
			},
		}
		settings := types.Settings{
			PrimaryFilter: 0.01,
			ScoreType:     "lte",
		}

		expected := map[string][]types.ScatterPoint{
			"a": {
				{Color: "#dfcd06", Label: "x", X: -10, Y: -0.67},
				{Color: "#dfcd06", Label: "y", X: 20, Y: 2.67},
			},
			"b": {
				{Color: "#dfcd06", Label: "x", X: -30, Y: 6},
			},
			"c": {
				{Color: "#6e97ff", Label: "z", X: 25, Y: 100},
			},
		}
		Expect(formatDataForPlot(data, settings)).To(Equal(expected))
	})

	Describe("filter data", func() {
		It("should return data filtered by score", func() {
			data := map[string]map[string]map[string]float64{
				"a": {
					"x": {"abundance": 10, "score": 0.01, "specificity": 0.67},
					"y": {"abundance": 20, "score": 0.01, "specificity": 2.67},
				},
				"b": {
					"x": {"abundance": 30, "score": 0, "specificity": 6},
				},
				"c": {
					"y": {"abundance": 15, "score": 0.02, "specificity": 1.5},
					"z": {"abundance": 25, "score": 0.01, "specificity": math.Inf(1)},
				},
			}
			settings := types.Settings{
				PrimaryFilter: 0.01,
				ScoreType:     "lte",
			}

			expected := map[string]map[string]map[string]float64{
				"a": {
					"x": {"abundance": 10, "score": 0.01, "specificity": 0.67},
					"y": {"abundance": 20, "score": 0.01, "specificity": 2.67},
				},
				"b": {
					"x": {"abundance": 30, "score": 0, "specificity": 6},
				},
				"c": {
					"z": {"abundance": 25, "score": 0.01, "specificity": math.Inf(1)},
				},
			}
			Expect(filterData(data, settings)).To(Equal(expected))
		})
	})

	Describe("max per condition", func() {
		It("should return the max (excluding infinity) per condition", func() {
			data := map[string]map[string]map[string]float64{
				"a": {
					"x": {"abundance": 10, "score": 0.01, "specificity": 0.67},
					"y": {"abundance": 20, "score": 0.01, "specificity": 2.67},
				},
				"b": {
					"x": {"abundance": 30, "score": 0, "specificity": 6},
				},
				"c": {
					"y": {"abundance": 15, "score": 0.02, "specificity": 1.5},
					"z": {"abundance": 25, "score": 0.01, "specificity": math.Inf(1)},
				},
				"d": {
					"w": {"abundance": 25, "score": 0.01, "specificity": math.Inf(1)},
				},
			}

			expected := map[string]float64{
				"a": 2.67,
				"b": 6,
				"c": 1.5,
				"d": 100,
			}
			Expect(defineMaxPerCondition(data)).To(Equal(expected))
		})
	})

	Describe("format as scatter points", func() {
		It("should return a slice of scatter points for each condition", func() {
			data := map[string]map[string]map[string]float64{
				"a": {
					"x": {"abundance": 10, "score": 0.01, "specificity": 0.67},
					"y": {"abundance": 20, "score": 0.01, "specificity": 2.67},
				},
				"b": {
					"x": {"abundance": 30, "score": 0, "specificity": 6},
				},
				"c": {
					"y": {"abundance": 15, "score": 0.02, "specificity": 1.5},
					"z": {"abundance": 25, "score": 0.01, "specificity": math.Inf(1)},
				},
			}
			maxPerCondition := map[string]float64{
				"a": 2.67,
				"b": 6,
				"c": 1.5,
			}

			expected := map[string][]types.ScatterPoint{
				"a": {
					{Color: "#dfcd06", Label: "x", X: 10, Y: 0.67},
					{Color: "#dfcd06", Label: "y", X: 20, Y: 2.67},
				},
				"b": {
					{Color: "#dfcd06", Label: "x", X: 30, Y: 6},
				},
				"c": {
					{Color: "#dfcd06", Label: "y", X: 15, Y: 1.5},
					{Color: "#6e97ff", Label: "z", X: 25, Y: 1.5},
				},
			}
			Expect(formatAsScatterPoints(data, maxPerCondition)).To(Equal(expected))
		})
	})
})
