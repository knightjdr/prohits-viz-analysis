package scatter

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Create file", func() {
	It("should create file for interactive scatter plot viewer", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		data := &Data{
			AnalysisType: "condition-condition",
			Filename:     "test/condition-condition.json",
			Legend: []map[string]string{
				{"color": "#ff0000", "text": "point1"},
				{"color": "#00ff00", "text": "point2"},
			},
			Parameters: types.Settings{
				Abundance:         "AvgSpec",
				Condition:         "Bait",
				Control:           "ctrl",
				Files:             []string{"file1", "file2"},
				Normalization:     "total",
				PrimaryFilter:     0.01,
				Readout:           "Prey",
				Score:             "bfdr",
				ScoreType:         "lte",
				SpecificityMetric: "fe",
				Type:              "condition-condition",
			},
			Plots: []types.ScatterPlot{
				{
					Labels: types.ScatterAxesLabels{X: "conditionX", Y: "conditionY"},
					Name:   "condition-condition",
					Points: []types.ScatterPoint{
						{Label: "readoutA", X: 1, Y: 3, Color: "#0066cc"},
						{Label: "readoutB", X: 0, Y: 4, Color: "#99ccff"},
						{Label: "readoutC", X: 2, Y: 0, Color: "#99ccff"},
					},
				},
			},
			Settings: map[string]interface{}{
				"logBase": 10,
				"xFilter": 0,
				"yFilter": 0,
			},
		}

		expected := "{\n" +
			"\t\"legend\": [{\"color\":\"#ff0000\",\"text\":\"point1\"},{\"color\":\"#00ff00\",\"text\":\"point2\"}],\n" +
			"\t\"parameters\": {\"abundanceColumn\":\"AvgSpec\",\"analysisType\":\"condition-condition\",\"conditionColumn\":\"Bait\",\"controlColumn\":\"ctrl\",\"files\":[\"file1\",\"file2\"],\"imageType\":\"scatter\",\"mockConditionAbundance\":false,\"normalization\":\"total\",\"readoutColumn\":\"Prey\",\"scoreColumn\":\"bfdr\",\"scoreType\":\"lte\"},\n" +
			"\t\"settings\": {\"logBase\":10,\"xFilter\":0,\"yFilter\":0},\n" +
			"\t\"plots\": [{\"labels\":{\"x\":\"conditionX\",\"y\":\"conditionY\"},\"name\":\"condition-condition\",\"points\":[{\"color\":\"#0066cc\",\"label\":\"readoutA\",\"x\":1,\"y\":3},{\"color\":\"#99ccff\",\"label\":\"readoutB\",\"x\":0,\"y\":4},{\"color\":\"#99ccff\",\"label\":\"readoutC\",\"x\":2,\"y\":0}]}]\n" +
			"}\n"

		Create(data)

		actual, _ := afero.ReadFile(fs.Instance, "test/condition-condition.json")
		Expect(string(actual)).To(Equal(expected))
	})
})
