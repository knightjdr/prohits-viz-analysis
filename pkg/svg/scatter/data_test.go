package scatter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Format data", func() {
	Describe("ticks", func() {
		Describe("upper (max) linear tick", func() {
			It("should define tick", func() {
				tests := []map[string]float64{
					{"max": 9.2, "step": 1},
					{"max": 62, "step": 10},
					{"max": 157, "step": 100},
				}
				expected := []float64{10, 70, 200}

				for i, test := range tests {
					Expect(getUpperLinearTick(test["max"], test["step"])).To(Equal(expected[i]))
				}
			})
		})

		Describe("calculate linear ticks", func() {
			It("should return ticks", func() {
				minmax := map[string]map[string]float64{
					"x": {"max": 9.2},
					"y": {"max": 72},
				}

				expected := Ticks{
					X: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
					Y: []float64{0, 10, 20, 30, 40, 50, 60, 70, 80},
				}
				Expect(calculateLinearTicks(minmax)).To(Equal(expected))
			})
		})

		Describe("upper (max) log tick", func() {
			It("should define tick for base 2", func() {
				tests := []float64{1.5, 3, 12, 275}
				expected := []float64{2, 4, 16, 512}
				for i, test := range tests {
					Expect(getUpperLogTick("2", test)).To(Equal(expected[i]))
				}
			})

			It("should define tick for base 10", func() {
				tests := []float64{1.5, 3, 12, 275}
				expected := []float64{10, 10, 100, 1000}
				for i, test := range tests {
					Expect(getUpperLogTick("10", test)).To(Equal(expected[i]))
				}
			})
		})

		Describe("lower (min) log tick", func() {
			It("should define tick for base 2", func() {
				tests := []float64{1, 1.25, 3}
				expected := []float64{0.125, 0.25, 1}
				for i, test := range tests {
					Expect(getLowerLogTick("2", test)).To(Equal(expected[i]))
				}
			})

			It("should define tick for base 10", func() {
				tests := []float64{5, 20, 1}
				expected := []float64{0.1, 1, 0.01}
				for i, test := range tests {
					Expect(getLowerLogTick("10", test)).To(Equal(expected[i]))
				}
			})
		})

		Describe("calculate log ticks", func() {
			It("should return ticks for base 2", func() {
				logBase := "2"
				minmax := map[string]map[string]float64{
					"x": {"max": 9.2, "min": 1},
					"y": {"max": 27, "min": 2},
				}

				expected := Ticks{
					X: []float64{0.125, 0.25, 0.5, 1, 2, 4, 8, 16},
					Y: []float64{1, 2, 4, 8, 16, 32},
				}
				Expect(calculateLogTicks(logBase, minmax)).To(Equal(expected))
			})

			It("should return ticks for base 10", func() {
				logBase := "10"
				minmax := map[string]map[string]float64{
					"x": {"max": 9.2, "min": 1},
					"y": {"max": 27, "min": 2},
				}

				expected := Ticks{
					X: []float64{0.01, 0.1, 1, 10},
					Y: []float64{0.1, 1, 10, 100},
				}
				Expect(calculateLogTicks(logBase, minmax)).To(Equal(expected))
			})
		})

		Describe("define ticks from data", func() {
			It("should define linear ticks", func() {
				logBase := "none"
				plot := []types.ScatterPoint{
					{X: 1, Y: 0.01},
					{X: 27, Y: 102},
					{X: 16, Y: 32},
				}

				expected := Ticks{
					X: []float64{0, 10, 20, 30},
					Y: []float64{0, 100, 200},
				}
				Expect(defineTicks(plot, logBase)).To(Equal(expected))
			})

			It("should define log base 2 ticks", func() {
				logBase := "2"
				plot := []types.ScatterPoint{
					{X: 1, Y: 0.01},
					{X: 27, Y: 102},
					{X: 16, Y: 32},
				}

				expected := Ticks{
					X: []float64{0.125, 0.25, 0.5, 1, 2, 4, 8, 16, 32},
					Y: []float64{0.125, 0.25, 0.5, 1, 2, 4, 8, 16, 32, 64, 128},
				}
				Expect(defineTicks(plot, logBase)).To(Equal(expected))
			})

			It("should define log base 10 ticks", func() {
				logBase := "10"
				plot := []types.ScatterPoint{
					{X: 1, Y: 0.01},
					{X: 27, Y: 102},
					{X: 16, Y: 32},
				}

				expected := Ticks{
					X: []float64{0.01, 0.1, 1, 10, 100},
					Y: []float64{0.01, 0.1, 1, 10, 100, 1000},
				}
				Expect(defineTicks(plot, logBase)).To(Equal(expected))
			})
		})
	})

	Describe("Scale data", func() {
		It("should scale linear data", func() {
			axisLength := float64(100)
			scatter := Scatter{
				LogBase: "none",
				Plot: []types.ScatterPoint{
					{Label: "a", X: 5, Y: 10},
					{Label: "b", X: 15, Y: 5},
					{Label: "c", X: 10, Y: 25},
				},
				Ticks: Ticks{
					X: []float64{0, 10, 20},
					Y: []float64{0, 10, 20, 30},
				},
			}

			expected := Scatter{
				LogBase: "none",
				Plot: []types.ScatterPoint{
					{Label: "a", X: 25, Y: 33.33},
					{Label: "b", X: 75, Y: 16.67},
					{Label: "c", X: 50, Y: 83.33},
				},
				Ticks: Ticks{
					X:      []float64{0, 50, 100},
					XLabel: []string{"0", "10", "20"},
					Y:      []float64{0, 33.33, 66.67, 100},
					YLabel: []string{"0", "10", "20", "30"},
				},
			}
			scaleData(&scatter, axisLength)
			Expect(scatter).To(Equal(expected))
		})
	})
})
