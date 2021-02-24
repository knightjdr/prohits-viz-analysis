package scv

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Create interactive file for SCV", func() {
	It("should create file with when knowness requested", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("interactive", 0755)

		data := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutX": {"abundance": 1},
				"readoutY": {"abundance": 4},
			},
		}
		known := map[string]map[string]bool{
			"conditionA": {"readoutX": true, "readoutY": false},
		}
		legend := types.CircHeatmapLegend{
			{Attribute: "abundance", Color: "blue", Max: 50, Min: 0},
		}
		settings := types.Settings{
			Abundance:     "AvgSpec",
			Condition:     "Bait",
			Control:       "ctrl",
			Files:         []string{"file1", "file2"},
			Known:         "interaction",
			Normalization: "total",
			Readout:       "Prey",
			Score:         "bfdr",
			ScoreType:     "lte",
			Type:          "scv",
		}

		expected := "{\n" +
			"\t\"circles\": {\"order\": [{\"attribute\":\"abundance\",\"color\":\"blue\",\"max\":50,\"min\":0}]},\n" +
			"\t\"parameters\": {\"abundanceColumn\":\"AvgSpec\",\"analysisType\":\"scv\",\"conditionColumn\":\"Bait\",\"controlColumn\":\"ctrl\",\"files\":[\"file1\",\"file2\"],\"imageType\":\"circheatmap\",\"normalization\":\"total\",\"readoutColumn\":\"Prey\",\"scoreColumn\":\"bfdr\",\"scoreType\":\"lte\"},\n" +
			"\t\"settings\": {\"sortByKnown\":true},\n" +
			"\t\"plots\": [{\"name\":\"conditionA\",\"readouts\":[{\"known\":true,\"label\":\"readoutX\",\"segments\":{\"abundance\":1.00}},{\"known\":false,\"label\":\"readoutY\",\"segments\":{\"abundance\":4.00}}]}]\n" +
			"}\n"

		createInteractive(data, known, legend, settings)

		actual, _ := afero.ReadFile(fs.Instance, "interactive/scv.json")
		Expect(string(actual)).To(Equal(expected))
	})

	It("should create file with when knowness not requested", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("interactive", 0755)

		data := map[string]map[string]map[string]float64{
			"conditionA": {
				"readoutX": {"abundance": 1},
				"readoutY": {"abundance": 4},
			},
		}
		var known map[string]map[string]bool = nil
		legend := types.CircHeatmapLegend{
			{Attribute: "abundance", Color: "blue", Max: 50, Min: 0},
		}
		settings := types.Settings{
			Abundance:     "AvgSpec",
			Condition:     "Bait",
			Control:       "ctrl",
			Files:         []string{"file1", "file2"},
			Known:         "",
			Normalization: "total",
			Readout:       "Prey",
			Score:         "bfdr",
			ScoreType:     "lte",
			Type:          "scv",
		}

		expected := "{\n" +
			"\t\"circles\": {\"order\": [{\"attribute\":\"abundance\",\"color\":\"blue\",\"max\":50,\"min\":0}]},\n" +
			"\t\"parameters\": {\"abundanceColumn\":\"AvgSpec\",\"analysisType\":\"scv\",\"conditionColumn\":\"Bait\",\"controlColumn\":\"ctrl\",\"files\":[\"file1\",\"file2\"],\"imageType\":\"circheatmap\",\"normalization\":\"total\",\"readoutColumn\":\"Prey\",\"scoreColumn\":\"bfdr\",\"scoreType\":\"lte\"},\n" +
			"\t\"settings\": {\"sortByKnown\":false},\n" +
			"\t\"plots\": [{\"name\":\"conditionA\",\"readouts\":[{\"known\":false,\"label\":\"readoutX\",\"segments\":{\"abundance\":1.00}},{\"known\":false,\"label\":\"readoutY\",\"segments\":{\"abundance\":4.00}}]}]\n" +
			"}\n"

		createInteractive(data, known, legend, settings)

		actual, _ := afero.ReadFile(fs.Instance, "interactive/scv.json")
		Expect(string(actual)).To(Equal(expected))
	})
})
