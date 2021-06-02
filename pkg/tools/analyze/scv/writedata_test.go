package scv

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

var _ = Describe("Write scv data to file", func() {
	It("should write data to file with knownness", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("other", 0755)

		plots := []types.CircHeatmap{
			{
				Name: "conditionA",
				Readouts: []types.CircHeatmapReadout{
					{
						Known: true,
						Label: "readoutX",
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
		legend := types.CircHeatmapLegend{
			{Attribute: "abundance", Color: "blue", Max: 50, Min: 0},
		}
		settings := types.Settings{
			Condition: "Bait",
			Known:     "interaction",
			Readout:   "Prey",
		}

		writeData(plots, legend, settings)

		expected := "Bait\tPrey\tabundance\tknown interaction\n" +
			"conditionA\treadoutX\t1.00\ttrue\n" +
			"conditionA\treadoutY\t4.00\tfalse\n" +
			"conditionB\treadoutY\t2.00\ttrue\n" +
			"conditionB\treadoutZ\t6.00\tfalse\n"

		actual, _ := afero.ReadFile(fs.Instance, "other/scv-data.txt")
		Expect(string(actual)).To(Equal(expected))
	})

	It("should write data to file without knownness", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("other", 0755)

		plots := []types.CircHeatmap{
			{
				Name: "conditionA",
				Readouts: []types.CircHeatmapReadout{
					{
						Known: false,
						Label: "readoutY",
						Segments: map[string]types.RoundedSegment{
							"abundance": 4,
						},
					},
					{
						Known: false,
						Label: "readoutX",
						Segments: map[string]types.RoundedSegment{
							"abundance": 1,
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
		legend := types.CircHeatmapLegend{
			{Attribute: "abundance", Color: "blue", Max: 50, Min: 0},
		}
		settings := types.Settings{
			Condition: "Bait",
			Known:     "",
			Readout:   "Prey",
		}

		writeData(plots, legend, settings)

		expected := "Bait\tPrey\tabundance\n" +
			"conditionA\treadoutY\t4.00\n" +
			"conditionA\treadoutX\t1.00\n" +
			"conditionB\treadoutZ\t6.00\n" +
			"conditionB\treadoutY\t2.00\n"

		actual, _ := afero.ReadFile(fs.Instance, "other/scv-data.txt")
		Expect(string(actual)).To(Equal(expected))
	})
})
