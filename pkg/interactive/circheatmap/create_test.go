package circheatmap

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Create file for scv", func() {
	It("should create file for interactive circheatmap viewer", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		data := &Data{
			Filename: "test/scv.json",
			Legend: types.CircHeatmapLegend{
				{Attribute: "abundance", Color: "blue", Filter: 0, Max: 50, Min: 0},
			},
			Parameters: types.Settings{
				Abundance:     "AvgSpec",
				Condition:     "Bait",
				Control:       "ctrl",
				Files:         []string{"file1", "file2"},
				Normalization: "total",
				Readout:       "Prey",
				Score:         "bfdr",
				ScoreType:     "lte",
				Type:          "scv",
			},
			Plots: []types.CircHeatmap{
				{
					Name: "conditionA",
					Readouts: []types.CircHeatmapReadout{
						{
							Known: true,
							Label: "readoutX",
							Segments: map[string]types.RoundedSegment{
								"attribute1": 1,
							},
						},
						{
							Known: false,
							Label: "readoutY",
							Segments: map[string]types.RoundedSegment{
								"attribute1": 4,
							},
						},
					},
				},
			},
			Settings: map[string]interface{}{
				"sortByKnown": true,
			},
		}

		expected := "{\n" +
			"\t\"circles\": {\"order\": [{\"attribute\":\"abundance\",\"color\":\"blue\",\"filter\":0,\"max\":50,\"min\":0}]},\n" +
			"\t\"parameters\": {\"abundanceColumn\":\"AvgSpec\",\"analysisType\":\"scv\",\"conditionColumn\":\"Bait\",\"controlColumn\":\"ctrl\",\"files\":[\"file1\",\"file2\"],\"imageType\":\"circheatmap\",\"normalization\":\"total\",\"readoutColumn\":\"Prey\",\"scoreColumn\":\"bfdr\",\"scoreType\":\"lte\"},\n" +
			"\t\"settings\": {\"sortByKnown\":true},\n" +
			"\t\"plots\": [{\"name\":\"conditionA\",\"readouts\":[{\"known\":true,\"label\":\"readoutX\",\"segments\":{\"attribute1\":1.00}},{\"known\":false,\"label\":\"readoutY\",\"segments\":{\"attribute1\":4.00}}]}]\n" +
			"}\n"

		Create(data)

		actual, _ := afero.ReadFile(fs.Instance, "test/scv.json")
		Expect(string(actual)).To(Equal(expected))
	})
})
