package scv

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Create interactive file for SCV", func() {
	It("should create file when knowness requested", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("interactive", 0755)

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
			"\t\"plots\": [{\"name\":\"conditionA\",\"readouts\":[{\"known\":true,\"label\":\"readoutX\",\"segments\":{\"abundance\":1.00}},{\"known\":false,\"label\":\"readoutY\",\"segments\":{\"abundance\":4.00}}]},{\"name\":\"conditionB\",\"readouts\":[{\"known\":true,\"label\":\"readoutY\",\"segments\":{\"abundance\":2.00}},{\"known\":false,\"label\":\"readoutZ\",\"segments\":{\"abundance\":6.00}}]}]\n" +
			"}\n"

		createInteractive(plots, legend, settings)

		actual, _ := afero.ReadFile(fs.Instance, "interactive/scv.json")
		Expect(string(actual)).To(Equal(expected))
	})

	It("should create file when knowness not requested", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("interactive", 0755)

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
			"\t\"plots\": [{\"name\":\"conditionA\",\"readouts\":[{\"known\":false,\"label\":\"readoutY\",\"segments\":{\"abundance\":4.00}},{\"known\":false,\"label\":\"readoutX\",\"segments\":{\"abundance\":1.00}}]},{\"name\":\"conditionB\",\"readouts\":[{\"known\":false,\"label\":\"readoutZ\",\"segments\":{\"abundance\":6.00}},{\"known\":false,\"label\":\"readoutY\",\"segments\":{\"abundance\":2.00}}]}]\n" +
			"}\n"

		createInteractive(plots, legend, settings)

		actual, _ := afero.ReadFile(fs.Instance, "interactive/scv.json")
		Expect(string(actual)).To(Equal(expected))
	})
})
