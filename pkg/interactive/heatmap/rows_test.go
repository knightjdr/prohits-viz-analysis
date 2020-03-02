package heatmap

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rows", func() {
	It("should write heatmap rows to file", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		data := &Data{
			AnalysisType: "heatmap",
			Matrices: &types.Matrices{
				Abundance: [][]float64{
					{25, 5, 50.2},
					{100, 30, 7},
					{5, 2.3, 8},
				},
				Readouts: []string{"prey1", "prey2", "prey3"},
			},
		}

		expected := "\t\"rowDB\": [\n" +
			"\t\t{\n" +
			"\t\t\t\"name\": \"prey1\",\n" +
			"\t\t\t\"data\": [\n" +
			"\t\t\t\t{\"value\": 25.00},\n" +
			"\t\t\t\t{\"value\": 5.00},\n" +
			"\t\t\t\t{\"value\": 50.20}\n" +
			"\t\t\t]\n" +
			"\t\t},\n" +
			"\t\t{\n" +
			"\t\t\t\"name\": \"prey2\",\n" +
			"\t\t\t\"data\": [\n" +
			"\t\t\t\t{\"value\": 100.00},\n" +
			"\t\t\t\t{\"value\": 30.00},\n" +
			"\t\t\t\t{\"value\": 7.00}\n" +
			"\t\t\t]\n" +
			"\t\t},\n" +
			"\t\t{\n" +
			"\t\t\t\"name\": \"prey3\",\n" +
			"\t\t\t\"data\": [\n" +
			"\t\t\t\t{\"value\": 5.00},\n" +
			"\t\t\t\t{\"value\": 2.30},\n" +
			"\t\t\t\t{\"value\": 8.00}\n" +
			"\t\t\t]\n" +
			"\t\t}\n" +
			"\t],\n"

		parseRows(data, writeString)
		Expect(svg.String()).To(Equal(expected))
	})

	It("should write dotplot rows to file", func() {
		var svg strings.Builder
		writeString := func(str string) {
			svg.WriteString(str)
		}

		data := &Data{
			AnalysisType: "dotplot",
			Matrices: &types.Matrices{
				Abundance: [][]float64{
					{25, 5, 50.2},
					{100, 30, 7},
					{5, 2.3, 8},
				},
				Ratio: [][]float64{
					{1, 0.5, 0.3},
					{1, 0.3, 0.1},
					{0.5, 0.25, 1},
				},
				Score: [][]float64{
					{0.01, 0, 0.02},
					{0, 0.01, 0.01},
					{0.02, 0.1, 0.01},
				},
				Readouts: []string{"prey1", "prey2", "prey3"},
			},
		}

		expected := "\t\"rowDB\": [\n" +
			"\t\t{\n" +
			"\t\t\t\"name\": \"prey1\",\n" +
			"\t\t\t\"data\": [\n" +
			"\t\t\t\t{\"ratio\": 1.00, \"score\": 0.01, \"value\": 25.00},\n" +
			"\t\t\t\t{\"ratio\": 0.50, \"score\": 0.00, \"value\": 5.00},\n" +
			"\t\t\t\t{\"ratio\": 0.30, \"score\": 0.02, \"value\": 50.20}\n" +
			"\t\t\t]\n" +
			"\t\t},\n" +
			"\t\t{\n" +
			"\t\t\t\"name\": \"prey2\",\n" +
			"\t\t\t\"data\": [\n" +
			"\t\t\t\t{\"ratio\": 1.00, \"score\": 0.00, \"value\": 100.00},\n" +
			"\t\t\t\t{\"ratio\": 0.30, \"score\": 0.01, \"value\": 30.00},\n" +
			"\t\t\t\t{\"ratio\": 0.10, \"score\": 0.01, \"value\": 7.00}\n" +
			"\t\t\t]\n" +
			"\t\t},\n" +
			"\t\t{\n" +
			"\t\t\t\"name\": \"prey3\",\n" +
			"\t\t\t\"data\": [\n" +
			"\t\t\t\t{\"ratio\": 0.50, \"score\": 0.02, \"value\": 5.00},\n" +
			"\t\t\t\t{\"ratio\": 0.25, \"score\": 0.10, \"value\": 2.30},\n" +
			"\t\t\t\t{\"ratio\": 1.00, \"score\": 0.01, \"value\": 8.00}\n" +
			"\t\t\t]\n" +
			"\t\t}\n" +
			"\t],\n"

		parseRows(data, writeString)
		Expect(svg.String()).To(Equal(expected))
	})
})
