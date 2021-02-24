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

		data := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutX": {"abundance": 1},
				"readoutY": {"abundance": 4},
			},
			"conditionB": {
				"readoutX": {"abundance": 2},
				"readoutY": {"abundance": 8},
			},
		}
		known := map[string]map[string]bool{
			"conditionA": {"readoutX": true, "readoutY": false},
			"conditionB": {"readoutX": true, "readoutY": true},
		}
		legend := types.CircHeatmapLegend{
			{Attribute: "abundance", Color: "blue", Max: 50, Min: 0},
		}
		settings := types.Settings{
			Condition: "Bait",
			Known:     "interaction",
			Readout:   "Prey",
		}

		writeData(data, known, legend, settings)

		expected := "Bait\tPrey\tabundance\tknown interaction\n" +
			"conditionA\treadoutX\t1.00\ttrue\n" +
			"conditionA\treadoutY\t4.00\tfalse\n" +
			"conditionB\treadoutX\t2.00\ttrue\n" +
			"conditionB\treadoutY\t8.00\ttrue\n"

		actual, _ := afero.ReadFile(fs.Instance, "other/scv-data.txt")
		Expect(string(actual)).To(Equal(expected))
	})

	It("should write data to file without knownness", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("other", 0755)

		data := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutX": {"abundance": 1},
				"readoutY": {"abundance": 4},
			},
			"conditionB": {
				"readoutX": {"abundance": 2},
				"readoutY": {"abundance": 8},
			},
		}
		var known map[string]map[string]bool = nil
		legend := types.CircHeatmapLegend{
			{Attribute: "abundance", Color: "blue", Max: 50, Min: 0},
		}
		settings := types.Settings{
			Condition: "Bait",
			Known:     "",
			Readout:   "Prey",
		}

		writeData(data, known, legend, settings)

		expected := "Bait\tPrey\tabundance\n" +
			"conditionA\treadoutX\t1.00\n" +
			"conditionA\treadoutY\t4.00\n" +
			"conditionB\treadoutX\t2.00\n" +
			"conditionB\treadoutY\t8.00\n"

		actual, _ := afero.ReadFile(fs.Instance, "other/scv-data.txt")
		Expect(string(actual)).To(Equal(expected))
	})
})
