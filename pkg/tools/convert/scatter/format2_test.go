package scatter

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

var _ = Describe("Format 2", func() {
	It("should parse plots and settings, and define legend", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fileContents := "details:\t{\"tool\": \"Specificity\", \"type\": \"Specificity\", \"bait\": \"AMOT_HeLa\", \"xAxis\": \"AvgSpec\", \"yAxis\": \"Specificity\", \"score\": \"FDR\"}\n" +
			"ALMS1\t34.57\t9.49\t6\t#dfcd06\tBFDR: 0<br>Specificity: 9.49<br>AvgSpec: 34.57\n" +
			"AMOTL1\t59\t4.18\t6\t#dfcd06\tBFDR: 0<br>Specificity: 4.18<br>AvgSpec: 59\n" +
			"AMOTL2\t7.5\t4.48\t6\t#dfcd06\tBFDR: 0<br>Specificity: 4.48<br>AvgSpec: 7.5\n" +
			"details:\t{\"tool\": \"Specificity\", \"type\": \"Specificity\", \"bait\": \"ANKRD28\", \"xAxis\": \"AvgSpec\", \"yAxis\": \"Specificity\", \"score\": \"FDR\"}\n" +
			"ACACB\t118.14\t3.84\t6\t#dfcd06\tBFDR: 0<br>Specificity: 3.84<br>AvgSpec: 118.14\n" +
			"ARHGEF2\t10.79\t29.83\t6\t#0066cc\tBFDR: 0<br>Specificity: infinite<br>AvgSpec: 10.79\n"

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/file.txt", []byte(fileContents), 0444)

		actualPlots, actualSettings, actualLegend := readFormat2("test/file.txt")

		expectedPlots := []types.ScatterPlot{
			{
				Labels: types.ScatterAxesLabels{
					X: "AvgSpec",
					Y: "Specificity",
				},
				Name: "AMOT_HeLa",
				Points: []types.ScatterPoint{
					{
						Color: "#dfcd06",
						Label: "ALMS1",
						X:     34.57,
						Y:     9.49,
					},
					{
						Color: "#dfcd06",
						Label: "AMOTL1",
						X:     59,
						Y:     4.18,
					},
					{
						Color: "#dfcd06",
						Label: "AMOTL2",
						X:     7.5,
						Y:     4.48,
					},
				},
			},
			{
				Labels: types.ScatterAxesLabels{
					X: "AvgSpec",
					Y: "Specificity",
				},
				Name: "ANKRD28",
				Points: []types.ScatterPoint{
					{
						Color: "#dfcd06",
						Label: "ACACB",
						X:     118.14,
						Y:     3.84,
					},
					{
						Color: "#0066cc",
						Label: "ARHGEF2",
						X:     10.79,
						Y:     29.83,
					},
				},
			},
		}
		expectedSettings := types.Settings{
			Abundance: "AvgSpec",
			Score:     "FDR",
			Type:      "specificity",
		}
		expectedLegend := []map[string]string{
			{
				"color": "#0066cc",
				"text":  "Infinite specificity",
			},
		}

		Expect(actualPlots).To(Equal(expectedPlots))
		Expect(actualSettings).To(Equal(expectedSettings))
		Expect(actualLegend).To(Equal(expectedLegend))
	})
})
