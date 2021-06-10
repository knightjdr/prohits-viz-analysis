package scatter

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

var _ = Describe("Format 1", func() {
	It("should parse plots and settings, and define legend", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fileContents := "entry\tAMOT_HeLa\tAvgSpec\tSpecificity\tSpecificity\tFDR\n" +
			"ALMS1\t34.57\t9.49\t6\tblack\tScore: 0\n" +
			"AMOTL1\t59\t4.18\t6\tblack\tScore: 0.04\n" +
			"AMOTL2\t7.5\t4.48\t6\tblack\tScore: 0\n" +
			"entry\tANKRD28\tAvgSpec\tSpecificity\tSpecificity\tFDR\n" +
			"ACACB\t118.14\t3.84\t6\tblack\tScore: 0.04\n" +
			"ARHGEF2\t10.79\t29.83\t6\tred\tScore: 0.04\n"

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/file.txt", []byte(fileContents), 0444)

		actualPlots, actualSettings, actualLegend := readFormat1("test/file.txt")

		expectedPlots := []types.ScatterPlot{
			{
				Labels: types.ScatterAxesLabels{
					X: "AvgSpec",
					Y: "Specificity",
				},
				Name: "AMOT_HeLa",
				Points: []types.ScatterPoint{
					{
						Color: "black",
						Label: "ALMS1",
						X:     34.57,
						Y:     9.49,
					},
					{
						Color: "black",
						Label: "AMOTL1",
						X:     59,
						Y:     4.18,
					},
					{
						Color: "black",
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
						Color: "black",
						Label: "ACACB",
						X:     118.14,
						Y:     3.84,
					},
					{
						Color: "red",
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
				"color": "#ff0000",
				"text":  "Infinite specificity",
			},
		}

		Expect(actualPlots).To(Equal(expectedPlots))
		Expect(actualSettings).To(Equal(expectedSettings))
		Expect(actualLegend).To(Equal(expectedLegend))
	})
})
