package parse

import (
	"os"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestHeatmapJSON(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Define json.
	json := `{
		"abundanceCap": 50,
		"annotationFontSize": 15,
		"annotations": [
			{ "text": "a", "x": 0.2, "y": 0.5 },
			{ "text": "b", "x": 0.6, "y": 0.1 }
		],
		"edgeColor": "blueBlack",
		"fillColor": "blueBlack",
		"imageType": "dotplot",
		"invertColor": false,
		"primaryFilter": 0.01,
		"rows": [
			{
				"data": [{"value": 5}, {"value": 10}, {"value": 40}],
				"name": "a"
			},
			{
				"data": [{"value": 8}, {"value": 60}, {"value": 15}],
				"name": "b"
			},
			{
				"data": [{"value": 17}, {"value": 5}, {"value": 30}],
				"name": "c"
			}
		],
		"scoreType": "lte",
		"secondaryFilter": 0.05
	}`

	// create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/testfile1.txt",
		[]byte(json),
		0444,
	)

	// Argument unmocking.
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// TEST1: returns struct from json.
	annotations := []typedef.Annotation{
		{Text: "a", X: 0.2, Y: 0.5},
		{Text: "b", X: 0.6, Y: 0.1},
	}
	rows := []Row{
		{
			Data: []Column{{Value: 5}, {Value: 10}, {Value: 40}},
			Name: "a",
		},
		{
			Data: []Column{{Value: 8}, {Value: 60}, {Value: 15}},
			Name: "b",
		},
		{
			Data: []Column{{Value: 17}, {Value: 5}, {Value: 30}},
			Name: "c",
		},
	}
	dotplotData := Data{
		AnnotationFontSize: 15,
		Annotations:        annotations,
		EdgeColor:          "blueBlack",
		FillColor:          "blueBlack",
		ImageType:          "dotplot",
		Invert:             false,
		MaximumAbundance:   50,
		PrimaryFilter:      0.01,
		Rows:               rows,
		ScoreType:          "lte",
		SecondaryFilter:    0.05,
	}

	os.Args = []string{
		"cmd",
		"-json", "test/testfile1.txt",
	}
	dotplotOutput, dotplotErr := HeatmapJSON("test/testfile1.txt")
	assert.Nil(t, dotplotErr, "All required arguments specified should not return an error")
	assert.EqualValues(t, &dotplotData, dotplotOutput)
}
