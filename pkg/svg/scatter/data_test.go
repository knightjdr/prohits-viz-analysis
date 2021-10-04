package scatter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Format data", func() {
	Describe("define axis boundaries", func() {
		Describe("linear axes", func() {
			It("should find suitable min/max for positive data", func() {
				logBase := "none"
				plot := []types.ScatterPoint{
					{Label: "a", X: 5, Y: 17},
					{Label: "b", X: 45, Y: 127},
					{Label: "c", X: 23, Y: 11},
				}

				expected := boundaries{
					x: boundary{min: 0, max: 50},
					y: boundary{min: 0, max: 200},
				}

				Expect(defineAxisBoundaries(plot, logBase)).To(Equal(expected))
			})

			It("should find suitable min/max for negative data", func() {
				logBase := "none"
				plot := []types.ScatterPoint{
					{Label: "a", X: -5, Y: -17},
					{Label: "b", X: -45, Y: -127},
					{Label: "c", X: -23, Y: -11},
				}

				expected := boundaries{
					x: boundary{min: -50, max: 0},
					y: boundary{min: -200, max: 0},
				}

				Expect(defineAxisBoundaries(plot, logBase)).To(Equal(expected))
			})

			It("should find suitable min/max for positive and negative data", func() {
				logBase := "none"
				plot := []types.ScatterPoint{
					{Label: "a", X: 5, Y: 17},
					{Label: "b", X: 45, Y: 127},
					{Label: "c", X: -23, Y: -11},
				}

				expected := boundaries{
					x: boundary{min: -30, max: 50},
					y: boundary{min: -100, max: 200},
				}

				Expect(defineAxisBoundaries(plot, logBase)).To(Equal(expected))
			})

			It("should find suitable min/max for positive and negative data on different scales", func() {
				logBase := "none"
				plot := []types.ScatterPoint{
					{Label: "a", X: 5, Y: 17},
					{Label: "b", X: 45, Y: 527},
					{Label: "c", X: -123, Y: -11},
				}

				expected := boundaries{
					x: boundary{min: -200, max: 100},
					y: boundary{min: -100, max: 600},
				}

				Expect(defineAxisBoundaries(plot, logBase)).To(Equal(expected))
			})
		})

		Describe("log axes", func() {
			Describe("base 2", func() {
				It("should find suitable min/max for positive data", func() {
					logBase := "2"
					plot := []types.ScatterPoint{
						{Label: "a", X: 5, Y: 17},
						{Label: "b", X: 45, Y: 127},
						{Label: "c", X: 23, Y: 0.6},
					}

					expected := boundaries{
						x: boundary{min: 0.5, max: 64},
						y: boundary{min: 0.5, max: 128},
					}

					Expect(defineAxisBoundaries(plot, logBase)).To(Equal(expected))
				})

				It("should find suitable min/max for negative data", func() {
					logBase := "2"
					plot := []types.ScatterPoint{
						{Label: "a", X: -5, Y: -17},
						{Label: "b", X: -45, Y: -127},
						{Label: "c", X: -23, Y: -0.6},
					}

					expected := boundaries{
						x: boundary{min: -64, max: -0.5},
						y: boundary{min: -128, max: -0.5},
					}

					Expect(defineAxisBoundaries(plot, logBase)).To(Equal(expected))
				})

				It("should find suitable min/max for negative data", func() {
					logBase := "2"
					plot := []types.ScatterPoint{
						{Label: "a", X: -5, Y: 17},
						{Label: "b", X: 45, Y: -127},
						{Label: "c", X: 23, Y: 0.6},
					}

					expected := boundaries{
						x: boundary{min: -8, max: 64},
						y: boundary{min: -128, max: 32},
					}

					Expect(defineAxisBoundaries(plot, logBase)).To(Equal(expected))
				})
			})

			Describe("base 10", func() {
				It("should find suitable min/max for positive data", func() {
					logBase := "10"
					plot := []types.ScatterPoint{
						{Label: "a", X: 5, Y: 17},
						{Label: "b", X: 45, Y: 127},
						{Label: "c", X: 23, Y: 0.6},
					}

					expected := boundaries{
						x: boundary{min: 0.1, max: 100},
						y: boundary{min: 0.1, max: 1000},
					}

					Expect(defineAxisBoundaries(plot, logBase)).To(Equal(expected))
				})

				It("should find suitable min/max for negative data", func() {
					logBase := "10"
					plot := []types.ScatterPoint{
						{Label: "a", X: -5, Y: -17},
						{Label: "b", X: -45, Y: -127},
						{Label: "c", X: -23, Y: -0.6},
					}

					expected := boundaries{
						x: boundary{min: -100, max: -0.1},
						y: boundary{min: -1000, max: -0.1},
					}

					Expect(defineAxisBoundaries(plot, logBase)).To(Equal(expected))
				})

				It("should find suitable min/max for negative data", func() {
					logBase := "10"
					plot := []types.ScatterPoint{
						{Label: "a", X: -5, Y: 17},
						{Label: "b", X: 45, Y: -127},
						{Label: "c", X: 23, Y: 0.6},
					}

					expected := boundaries{
						x: boundary{min: -10, max: 100},
						y: boundary{min: -1000, max: 100},
					}

					Expect(defineAxisBoundaries(plot, logBase)).To(Equal(expected))
				})
			})
		})
	})

	Describe("define axis boundaries", func() {
		Describe("linear axes", func() {
			It("should return ticks for positive values", func() {
				axis := boundaries{
					x: boundary{
						max: 10,
						min: 0,
					},
					y: boundary{
						max: 30,
						min: 0,
					},
				}
				logBase := "none"

				expected := Ticks{
					X: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
					Y: []float64{0, 10, 20, 30},
				}
				Expect(defineTicks(axis, logBase)).To(Equal(expected))
			})

			It("should return ticks for negative values", func() {
				axis := boundaries{
					x: boundary{
						max: 0,
						min: -10,
					},
					y: boundary{
						max: 0,
						min: -30,
					},
				}
				logBase := "none"

				expected := Ticks{
					X: []float64{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0},
					Y: []float64{-30, -20, -10, 0},
				}
				Expect(defineTicks(axis, logBase)).To(Equal(expected))
			})

			It("should return ticks for positive and negative values", func() {
				axis := boundaries{
					x: boundary{
						max: 30,
						min: -10,
					},
					y: boundary{
						max: 10,
						min: -20,
					},
				}
				logBase := "none"

				expected := Ticks{
					X: []float64{-10, 0, 10, 20, 30},
					Y: []float64{-20, -10, 0, 10},
				}
				Expect(defineTicks(axis, logBase)).To(Equal(expected))
			})
		})

		Describe("log base 2 axes", func() {
			It("should return ticks for positive values", func() {
				axis := boundaries{
					x: boundary{
						max: 16,
						min: 1,
					},
					y: boundary{
						max: 8,
						min: 1,
					},
				}
				logBase := "2"

				expected := Ticks{
					X: []float64{1, 2, 4, 8, 16},
					Y: []float64{1, 2, 4, 8},
				}
				Expect(defineTicks(axis, logBase)).To(Equal(expected))
			})

			It("should return ticks for negative values", func() {
				axis := boundaries{
					x: boundary{
						max: -1,
						min: -16,
					},
					y: boundary{
						max: -1,
						min: -8,
					},
				}
				logBase := "2"

				expected := Ticks{
					X: []float64{-16, -8, -4, -2, -1},
					Y: []float64{-8, -4, -2, -1},
				}
				Expect(defineTicks(axis, logBase)).To(Equal(expected))
			})

			It("should return ticks for positive and negative values", func() {
				axis := boundaries{
					x: boundary{
						max: 16,
						min: -8,
					},
					y: boundary{
						max: 8,
						min: -64,
					},
				}
				logBase := "2"

				expected := Ticks{
					X: []float64{-8, -4, -2, -1, -0.5, 0.5, 1, 2, 4, 8, 16},
					Y: []float64{-64, -32, -16, -8, -4, -2, -1, -0.5, 0.5, 1, 2, 4, 8},
				}
				Expect(defineTicks(axis, logBase)).To(Equal(expected))
			})
		})

		Describe("log base 10 axes", func() {
			It("should return ticks for positive values", func() {
				axis := boundaries{
					x: boundary{
						max: 10,
						min: 0.1,
					},
					y: boundary{
						max: 100,
						min: 1,
					},
				}
				logBase := "10"

				expected := Ticks{
					X: []float64{0.1, 1, 10},
					Y: []float64{1, 10, 100},
				}
				Expect(defineTicks(axis, logBase)).To(Equal(expected))
			})

			It("should return ticks for negative values", func() {
				axis := boundaries{
					x: boundary{
						max: -0.1,
						min: -10,
					},
					y: boundary{
						max: -1,
						min: -100,
					},
				}
				logBase := "10"

				expected := Ticks{
					X: []float64{-10, -1, -0.1},
					Y: []float64{-100, -10, -1},
				}
				Expect(defineTicks(axis, logBase)).To(Equal(expected))
			})

			It("should return ticks for positive and negative values", func() {
				axis := boundaries{
					x: boundary{
						max: 10,
						min: -100,
					},
					y: boundary{
						max: 1,
						min: -10,
					},
				}
				logBase := "10"

				expected := Ticks{
					X: []float64{-100, -10, -1, -0.1, 0.1, 1, 10},
					Y: []float64{-10, -1, -0.1, 0.1, 1},
				}
				Expect(defineTicks(axis, logBase)).To(Equal(expected))
			})
		})
	})

	Describe("define axes", func() {
		Describe("linear axes", func() {
			It("should define axes with positive ticks", func() {
				ticks := Ticks{
					X: []float64{0, 10, 20},
					Y: []float64{0, 10, 20, 30},
				}

				expected := Axes{
					X: Line{X1: 0, X2: 20, Y1: 0, Y2: 0},
					Y: Line{X1: 0, X2: 0, Y1: 0, Y2: 30},
				}
				Expect(defineAxes(ticks)).To(Equal(expected))
			})

			It("should define axes with negative ticks", func() {
				ticks := Ticks{
					X: []float64{-20, -10, 0},
					Y: []float64{-30, -20, -10, 0},
				}

				expected := Axes{
					X: Line{X1: -20, X2: 0, Y1: 0, Y2: 0},
					Y: Line{X1: 0, X2: 0, Y1: -30, Y2: 0},
				}
				Expect(defineAxes(ticks)).To(Equal(expected))
			})

			It("should define axes with positive and negative ticks", func() {
				ticks := Ticks{
					X: []float64{-20, -10, 0, 10},
					Y: []float64{-30, -20, -10, 0, 10, 20, 30},
				}

				expected := Axes{
					X: Line{X1: -20, X2: 10, Y1: 0, Y2: 0},
					Y: Line{X1: 0, X2: 0, Y1: -30, Y2: 30},
				}
				Expect(defineAxes(ticks)).To(Equal(expected))
			})
		})

		Describe("log axes", func() {
			It("should define axes with positive ticks", func() {
				ticks := Ticks{
					X: []float64{1, 2, 4},
					Y: []float64{1, 2, 4, 8},
				}

				expected := Axes{
					X: Line{X1: 1, X2: 4, Y1: 1, Y2: 1},
					Y: Line{X1: 1, X2: 1, Y1: 1, Y2: 8},
				}
				Expect(defineAxes(ticks)).To(Equal(expected))
			})

			It("should define axes with negative ticks", func() {
				ticks := Ticks{
					X: []float64{-4, -2, -1},
					Y: []float64{-8, -4, -2, -1},
				}

				expected := Axes{
					X: Line{X1: -4, X2: -1, Y1: -1, Y2: -1},
					Y: Line{X1: -1, X2: -1, Y1: -8, Y2: -1},
				}
				Expect(defineAxes(ticks)).To(Equal(expected))
			})

			It("should define axes with positive and negative ticks", func() {
				ticks := Ticks{
					X: []float64{-4, -2, -1, 1, 2},
					Y: []float64{-8, -4, -2, -1, 1, 2, 4, 8},
				}

				expected := Axes{
					X: Line{X1: -4, X2: 2, Y1: 0, Y2: 0},
					Y: Line{X1: 0, X2: 0, Y1: -8, Y2: 8},
				}
				Expect(defineAxes(ticks)).To(Equal(expected))
			})
		})
	})

	Describe("Scale data", func() {
		Describe("linear axes", func() {
			It("should scale positive data", func() {
				axisLength := float64(100)
				scatter := Scatter{
					Axes: Axes{
						X: Line{X1: 0, X2: 20, Y1: 0, Y2: 0},
						Y: Line{X1: 0, X2: 0, Y1: 0, Y2: 30},
					},
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
					Axes: Axes{
						X: Line{X1: 0, X2: 100, Y1: 100, Y2: 100},
						Y: Line{X1: 0, X2: 0, Y1: 100, Y2: 0},
					},
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

			It("should scale negative data", func() {
				axisLength := float64(100)
				scatter := Scatter{
					Axes: Axes{
						X: Line{X1: -20, X2: 0, Y1: 0, Y2: 0},
						Y: Line{X1: 0, X2: 0, Y1: -30, Y2: 0},
					},
					LogBase: "none",
					Plot: []types.ScatterPoint{
						{Label: "a", X: -5, Y: -10},
						{Label: "b", X: -15, Y: -5},
						{Label: "c", X: -10, Y: -25},
					},
					Ticks: Ticks{
						X: []float64{-20, -10, -0},
						Y: []float64{-30, -20, -10, 0},
					},
				}

				expected := Scatter{
					Axes: Axes{
						X: Line{X1: 0, X2: 100, Y1: 0, Y2: 0},
						Y: Line{X1: 100, X2: 100, Y1: 100, Y2: 0},
					},
					LogBase: "none",
					Plot: []types.ScatterPoint{
						{Label: "a", X: 75, Y: 66.67},
						{Label: "b", X: 25, Y: 83.33},
						{Label: "c", X: 50, Y: 16.67},
					},
					Ticks: Ticks{
						X:      []float64{0, 50, 100},
						XLabel: []string{"-20", "-10", "0"},
						Y:      []float64{0, 33.33, 66.67, 100},
						YLabel: []string{"-30", "-20", "-10", "0"},
					},
				}
				scaleData(&scatter, axisLength)
				Expect(scatter).To(Equal(expected))
			})

			It("should scale positive and negative data", func() {
				axisLength := float64(100)
				scatter := Scatter{
					Axes: Axes{
						X: Line{X1: -20, X2: 10, Y1: 0, Y2: 0},
						Y: Line{X1: 0, X2: 0, Y1: -30, Y2: 10},
					},
					LogBase: "none",
					Plot: []types.ScatterPoint{
						{Label: "a", X: 5, Y: 10},
						{Label: "b", X: -15, Y: 5},
						{Label: "c", X: 10, Y: -25},
					},
					Ticks: Ticks{
						X: []float64{-20, -10, -0, 10},
						Y: []float64{-30, -20, -10, 0, 10},
					},
				}

				expected := Scatter{
					Axes: Axes{
						X: Line{X1: 0, X2: 100, Y1: 25, Y2: 25},
						Y: Line{X1: 66.67, X2: 66.67, Y1: 100, Y2: 0},
					},
					LogBase: "none",
					Plot: []types.ScatterPoint{
						{Label: "a", X: 83.33, Y: 100},
						{Label: "b", X: 16.67, Y: 87.5},
						{Label: "c", X: 100, Y: 12.5},
					},
					Ticks: Ticks{
						X:      []float64{0, 33.33, 66.67, 100},
						XLabel: []string{"-20", "-10", "0", "10"},
						Y:      []float64{0, 25, 50, 75, 100},
						YLabel: []string{"-30", "-20", "-10", "0", "10"},
					},
				}
				scaleData(&scatter, axisLength)
				Expect(scatter).To(Equal(expected))
			})
		})

		Describe("log axes", func() {
			It("should scale positive data", func() {
				axisLength := float64(100)
				scatter := Scatter{
					Axes: Axes{
						X: Line{X1: 1, X2: 4, Y1: 1, Y2: 1},
						Y: Line{X1: 1, X2: 1, Y1: 1, Y2: 8},
					},
					LogBase: "2",
					Plot: []types.ScatterPoint{
						{Label: "a", X: 2, Y: 4},
						{Label: "b", X: 1, Y: 2},
						{Label: "c", X: 3, Y: 6},
					},
					Ticks: Ticks{
						X: []float64{1, 2, 4},
						Y: []float64{1, 2, 4, 8},
					},
				}

				expected := Scatter{
					Axes: Axes{
						X: Line{X1: 0, X2: 100, Y1: 100, Y2: 100},
						Y: Line{X1: 0, X2: 0, Y1: 100, Y2: 0},
					},
					LogBase: "2",
					Plot: []types.ScatterPoint{
						{Label: "a", X: 50, Y: 66.67},
						{Label: "b", X: 0, Y: 33.33},
						{Label: "c", X: 79.25, Y: 86.17},
					},
					Ticks: Ticks{
						X:      []float64{0, 50, 100},
						XLabel: []string{"1", "2", "4"},
						Y:      []float64{0, 33.33, 66.67, 100},
						YLabel: []string{"1", "2", "4", "8"},
					},
				}
				scaleData(&scatter, axisLength)
				Expect(scatter).To(Equal(expected))
			})

			It("should scale negative data", func() {
				axisLength := float64(100)
				scatter := Scatter{
					Axes: Axes{
						X: Line{X1: -4, X2: -1, Y1: -1, Y2: -1},
						Y: Line{X1: -1, X2: -1, Y1: -8, Y2: -1},
					},
					LogBase: "2",
					Plot: []types.ScatterPoint{
						{Label: "a", X: -2, Y: -4},
						{Label: "b", X: -1, Y: -2},
						{Label: "c", X: -3, Y: -6},
					},
					Ticks: Ticks{
						X: []float64{-4, -2, -1},
						Y: []float64{-8, -4, -2, -1},
					},
				}

				expected := Scatter{
					Axes: Axes{
						X: Line{X1: 0, X2: 100, Y1: 0, Y2: 0},
						Y: Line{X1: 100, X2: 100, Y1: 100, Y2: 0},
					},
					LogBase: "2",
					Plot: []types.ScatterPoint{
						{Label: "a", X: 50, Y: 33.33},
						{Label: "b", X: 100, Y: 66.67},
						{Label: "c", X: 20.75, Y: 13.83},
					},
					Ticks: Ticks{
						X:      []float64{0, 50, 100},
						XLabel: []string{"-4", "-2", "-1"},
						Y:      []float64{0, 33.33, 66.67, 100},
						YLabel: []string{"-8", "-4", "-2", "-1"},
					},
				}
				scaleData(&scatter, axisLength)
				Expect(scatter).To(Equal(expected))
			})

			It("should scale positive and negative data", func() {
				axisLength := float64(100)
				scatter := Scatter{
					Axes: Axes{
						X: Line{X1: -4, X2: 2, Y1: 0, Y2: 0},
						Y: Line{X1: 0, X2: 0, Y1: -8, Y2: 4},
					},
					LogBase: "2",
					Plot: []types.ScatterPoint{
						{Label: "a", X: 2, Y: 4},
						{Label: "b", X: 1, Y: -2},
						{Label: "c", X: -3, Y: -6},
					},
					Ticks: Ticks{
						X: []float64{-4, -2, -1, 1, 2},
						Y: []float64{-8, -4, -2, -1, 1, 2, 4},
					},
				}

				expected := Scatter{
					Axes: Axes{
						X: Line{X1: 0, X2: 100, Y1: 41.67, Y2: 41.67},
						Y: Line{X1: 62.5, X2: 62.5, Y1: 100, Y2: 0},
					},
					LogBase: "2",
					Plot: []types.ScatterPoint{
						{Label: "a", X: 100, Y: 100},
						{Label: "b", X: 75, Y: 33.33},
						{Label: "c", X: 10.38, Y: 6.92},
					},
					Ticks: Ticks{
						X:      []float64{0, 25, 50, 75, 100},
						XLabel: []string{"-4", "-2", "-1", "1", "2"},
						Y:      []float64{0, 16.67, 33.33, 50, 66.67, 83.33, 100},
						YLabel: []string{"-8", "-4", "-2", "-1", "1", "2", "4"},
					},
				}
				scaleData(&scatter, axisLength)
				Expect(scatter).To(Equal(expected))
			})
		})
	})
})
