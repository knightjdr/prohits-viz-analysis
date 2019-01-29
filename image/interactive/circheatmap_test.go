package interactive

import (
	"strings"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestCircHeatmap(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Starting with some hypothetical data for a dotplot, generate a data matrix
	// with formatted rows for use in the heatmap.
	plots := []typedef.CircHeatmapPlot{
		{
			Name: "condition1",
			Readouts: []map[string]interface{}{
				{"name": "readout1", "known": true},
				{"name": "readout3", "known": false},
			},
			Segments: []typedef.CircHeatmapSegments{
				{Name: "Abundance", Values: []float64{10, 5}},
				{Name: "FC", Values: []float64{2, 3}},
			},
		},
	}
	parameters := map[string]interface{}{
		"files":     []string{"file1.txt", "file2.txt"},
		"imageType": "circ-heatmap",
	}
	settings := map[string]interface{}{
		"known": true,
		"plot":  0,
	}
	segmentSettings := []typedef.CircHeatmapSetttings{
		{AbundanceCap: 50, Color: "blueBlack", MinAbundance: 0, Name: "Abundance"},
		{AbundanceCap: 50, Color: "redBlack", MinAbundance: 0, Name: "FC"},
	}

	// TEST: typical data conversion to json.
	want := "{" +
		"\"availablePlots\":[" +
		"{\"name\":\"condition1\"," +
		"\"readouts\":[{\"known\":true,\"name\":\"readout1\"},{\"known\":false,\"name\":\"readout3\"}]," +
		"\"segments\":[{\"name\":\"Abundance\",\"values\":[10,5]},{\"name\":\"FC\",\"values\":[2,3]}]}]," +
		"\"circHeatmapSettings\":[" +
		"{\"abundanceCap\":50,\"color\":\"blueBlack\",\"minAbundance\":0,\"name\":\"Abundance\"}," +
		"{\"abundanceCap\":50,\"color\":\"redBlack\",\"minAbundance\":0,\"name\":\"FC\"}]," +
		"\"parameters\":{\"files\":[\"file1.txt\",\"file2.txt\"],\"imageType\":\"circ-heatmap\"}," +
		"\"settings\":{\"current\":{\"known\":true,\"plot\":0}}}"
	CircHeatmap(plots, parameters, settings, segmentSettings, "file.json")
	bytes, _ := afero.ReadFile(fs.Instance, "file.json")
	json := string(bytes)
	json = strings.Replace(json, " ", "", -1)
	json = strings.Replace(json, "\n", "", -1)
	json = strings.Replace(json, "\t", "", -1)
	assert.Equal(t, want, json, "JSON should be formatted correctly")
}
