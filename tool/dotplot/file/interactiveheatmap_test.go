package file

import (
	"image"
	"image/color"
	"image/png"
	"strings"
	"testing"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/typedef"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestInteractiveHeatmap(t *testing.T) {
	// Mock filesystem.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	fs.Instance.MkdirAll("interactive", 0755)
	fs.Instance.MkdirAll("minimap", 0755)

	// Create png.
	pngImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1, 1}})
	c := color.RGBA{uint8(0), uint8(0), uint8(0), 255}
	pngImage.Set(0, 0, c)
	myfile, _ := fs.Instance.Create("minimap/test.png")
	png.Encode(myfile, pngImage)

	// Starting with some hypothetical data for a dotplot, generate a data matrix
	// with formatted rows for use in the heatmap.
	abundance := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	columns := []string{"col1", "col2", "col3"}
	parameters := typedef.Parameters{
		Abundance:          "abd",
		AbundanceCap:       float64(50),
		AnalysisType:       "dotplot",
		Condition:          "condition",
		Clustering:         "hierarchical",
		ClusteringMethod:   "complete",
		ClusteringOptimize: false,
		Control:            "control",
		Distance:           "euclidean",
		Files:              []string{"file1.txt", "file2.txt"},
		FillColor:          "blueBlack",
		LogBase:            "none",
		MinAbundance:       float64(0),
		Normalization:      "none",
		Readout:            "readout",
		PrimaryFilter:      0.01,
		Score:              "score",
		ScoreType:          "lte",
		SecondaryFilter:    0.05,
	}
	rows := []string{"row1", "row2", "row3"}

	// TEST1: typical date conversion to json.
	want := "{" +
		"\"columns\":{" +
		"\"names\":[\"col1\",\"col2\",\"col3\"]," +
		"\"ref\":null}," +
		"\"parameters\":{" +
		"\"abundanceColumn\":\"abd\"," +
		"\"analysisType\":\"dotplot\"," +
		"\"clustering\":\"hierarchical\"," +
		"\"clusteringMethod\":\"complete\"," +
		"\"clusteringOptimize\":false," +
		"\"conditionColumn\":\"condition\"," +
		"\"controlColumn\":\"control\"," +
		"\"distance\":\"euclidean\"," +
		"\"files\":[\"file1.txt\",\"file2.txt\"]," +
		"\"imageType\":\"heatmap\"," +
		"\"logBase\":\"none\"," +
		"\"normalization\":\"none\"," +
		"\"readoutColumn\":\"readout\"," +
		"\"scoreColumn\":\"score\"," +
		"\"scoreType\":\"lte\"," +
		"\"xLabel\":\"condition\"," +
		"\"yLabel\":\"readout\"" +
		"}," +
		"\"settings\":{\"current\":{" +
		"\"abundanceCap\":50," +
		"\"fillColor\":\"blueBlack\"," +
		"\"imageType\":\"heatmap\"," +
		"\"invertColor\":false," +
		"\"minAbundance\":0," +
		"\"primaryFilter\":0.01," +
		"\"secondaryFilter\":0.05" +
		"}}," +
		"\"rows\":{\"list\":[" +
		"{\"data\":[{\"value\":1},{\"value\":2},{\"value\":3}],\"name\":\"row1\"}," +
		"{\"data\":[{\"value\":4},{\"value\":5},{\"value\":6}],\"name\":\"row2\"}," +
		"{\"data\":[{\"value\":7},{\"value\":8},{\"value\":9}],\"name\":\"row3\"}" +
		"]}," +
		"\"minimap\":{\"image\":\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg==\"}}"
	InteractiveHeatmap("heatmap", abundance, [][]float64{}, [][]float64{}, columns, rows, "test", parameters)
	bytes, _ := afero.ReadFile(fs.Instance, "interactive/test.json")
	json := string(bytes)
	json = strings.Replace(json, " ", "", -1)
	json = strings.Replace(json, "\n", "", -1)
	json = strings.Replace(json, "\t", "", -1)
	assert.Equal(
		t,
		want,
		json,
		"JSON not generated correctly",
	)
}
