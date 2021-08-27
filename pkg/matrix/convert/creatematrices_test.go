package convert

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Create matrices", func() {
	It("should create matrices", func() {
		data := &tableData{
			conditions: map[string]int{
				"condition1": 0,
				"condition2": 2,
				"condition3": 1,
			},
			readoutCondition: map[readoutCondition]readoutData{
				{"readout1", "condition1"}: {5, 0.01},
				{"readout3", "condition1"}: {10, 0.02},
				{"readout2", "condition1"}: {23, 0},
				{"readout3", "condition3"}: {7, 0.01},
				{"readout1", "condition3"}: {14.3, 0.08},
				{"readout2", "condition2"}: {17.8, 0.01},
				{"readout1", "condition2"}: {2, 0.01},
			},
			readouts: map[string]int{
				"readout1": 0,
				"readout2": 2,
				"readout3": 1,
			},
			worstScore: 0.08,
		}
		matrices := &types.Matrices{
			Conditions: []string{"condition1", "condition2", "condition3"},
			Readouts:   []string{"readout1", "readout2", "readout3"},
		}
		settings := ConversionSettings{
			CalculateRatios: true,
			RatioDimension:  "diameter",
		}

		expected := &types.Matrices{
			Abundance: [][]float64{
				{5, 2, 14.3},
				{23, 17.8, 0},
				{10, 0, 7},
			},
			Conditions: []string{"condition1", "condition2", "condition3"},
			Ratio: [][]float64{
				{0.35, 0.14, 1},
				{1, 0.77, 0},
				{1, 0, 0.7},
			},
			Readouts: []string{"readout1", "readout2", "readout3"},
			Score: [][]float64{
				{0.01, 0.01, 0.08},
				{0, 0.01, 0.08},
				{0.02, 0.08, 0.01},
			},
		}

		createMatrices(matrices, data, settings)
		for i, row := range matrices.Abundance {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected.Abundance[i][j], 0.01), "should return abundance matrix")
			}
		}
		for i, row := range matrices.Ratio {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected.Ratio[i][j], 0.01), "should return ratio matrix")
			}
		}
		for i, row := range matrices.Score {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected.Score[i][j], 0.01), "should return score matrix")
			}
		}
	})
})

var _ = Describe("Add abundance and score matrices", func() {
	It("should create matrices", func() {
		data := &tableData{
			conditions: map[string]int{
				"condition1": 0,
				"condition2": 2,
				"condition3": 1,
			},
			readoutCondition: map[readoutCondition]readoutData{
				{"readout1", "condition1"}: {5, 0.01},
				{"readout3", "condition1"}: {10, 0.02},
				{"readout2", "condition1"}: {23, 0},
				{"readout3", "condition3"}: {7, 0.01},
				{"readout1", "condition3"}: {14.3, 0.08},
				{"readout2", "condition2"}: {17.8, 0.01},
				{"readout1", "condition2"}: {2, 0.01},
			},
			readouts: map[string]int{
				"readout1": 0,
				"readout2": 2,
				"readout3": 1,
			},
			worstScore: 0.08,
		}
		matrices := &types.Matrices{
			Conditions: []string{"condition1", "condition2", "condition3"},
			Readouts:   []string{"readout1", "readout2", "readout3"},
		}

		expected := &types.Matrices{
			Abundance: [][]float64{
				{5, 2, 14.3},
				{23, 17.8, 0},
				{10, 0, 7},
			},
			Conditions: []string{"condition1", "condition2", "condition3"},
			Readouts:   []string{"readout1", "readout2", "readout3"},
			Score: [][]float64{
				{0.01, 0.01, 0.08},
				{0, 0.01, 0.08},
				{0.02, 0.08, 0.01},
			},
		}

		addAbundanceAndScoreMatrices(matrices, data)
		Expect(matrices).To(Equal(expected))
	})
})

var _ = Describe("Add ratio matrix", func() {
	It("should add matrix", func() {
		matrices := &types.Matrices{
			Abundance: [][]float64{
				{5, 2, 14.3},
				{23, 17.8, 0},
				{10, 0, 7},
				{0.1, 1, -2},
			},
			Conditions: []string{"condition1", "condition2", "condition3"},
			Readouts:   []string{"readout1", "readout2", "readout3"},
			Score: [][]float64{
				{0.01, 0.01, 0.08},
				{0, 0.01, 0.08},
				{0.02, 0.08, 0.01},
				{0.5, 0.1, 0},
			},
		}
		settings := ConversionSettings{
			CalculateRatios: true,
			RatioDimension:  "diameter",
		}

		expected := [][]float64{
			{0.35, 0.14, 1},
			{1, 0.77, 0},
			{1, 0, 0.7},
			{0.05, 0.5, 1},
		}

		addRatioMatrix(matrices, settings)
		for i, row := range matrices.Ratio {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected[i][j], 0.01))
			}
		}
	})

	It("should add matrix adjusting values for area calculation", func() {
		matrices := &types.Matrices{
			Abundance: [][]float64{
				{5, 2, 14.3},
				{23, 17.8, 0},
				{10, 0, 7},
			},
			Conditions: []string{"condition1", "condition2", "condition3"},
			Readouts:   []string{"readout1", "readout2", "readout3"},
			Score: [][]float64{
				{0.01, 0.01, 0.08},
				{0, 0.01, 0.08},
				{0.02, 0.08, 0.01},
			},
		}
		settings := ConversionSettings{
			CalculateRatios: true,
			RatioDimension:  "area",
		}

		expected := [][]float64{
			{0.592, 0.374, 1},
			{1, 0.877, 0},
			{1, 0, 0.837},
		}

		addRatioMatrix(matrices, settings)
		for i, row := range matrices.Ratio {
			for j, value := range row {
				Expect(value).To(BeNumerically("~", expected[i][j], 0.001))
			}
		}
	})

	It("should not add matrix", func() {
		matrices := &types.Matrices{
			Abundance: [][]float64{
				{5, 2, 14.3},
				{23, 17.8, 0},
				{10, 0, 7},
			},
			Conditions: []string{"condition1", "condition2", "condition3"},
			Readouts:   []string{"readout1", "readout2", "readout3"},
			Score: [][]float64{
				{0.01, 0.01, 0.08},
				{0, 0.01, 0.08},
				{0.02, 0.08, 0.01},
			},
		}
		settings := ConversionSettings{
			CalculateRatios: false,
		}

		addRatioMatrix(matrices, settings)
		Expect(matrices.Ratio).To(BeNil())
	})
})
