package dotplot

import (
	"image"
	"image/color"
	"image/png"
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

	// Create png.
	pngImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1, 1}})
	c := color.RGBA{uint8(0), uint8(0), uint8(0), 255}
	pngImage.Set(0, 0, c)
	myfile, _ := fs.Instance.Create("test.png")
	png.Encode(myfile, pngImage)

	// Starting with some hypothetical data for a dotplot, generate a data matrix
	// with formatted rows for use in the heatmap.
	abundance := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	columns := []string{"col1", "col2", "col3"}
	params := typedef.Parameters{
		Abundance:        "abd",
		AnalysisType:     "dotplot",
		Bait:             "bait",
		Clustering:       "hierarchical",
		ClusteringMethod: "complete",
		Control:          "control",
		Distance:         "euclidean",
		Files:            []string{"file1.txt", "file2.txt"},
		FillColor:        "blueBlack",
		LogBase:          "none",
		MaximumAbundance: float64(50),
		MinimumAbundance: float64(0),
		Normalization:    "none",
		Prey:             "prey",
		PrimaryFilter:    0.01,
		Score:            "score",
		ScoreType:        "lte",
		SecondaryFilter:  0.05,
	}
	rows := []string{"row1", "row2", "row3"}

	// TEST1: typical date conversion to json.
	want := "{" +
		"\"columns\":" +
		"[\"col1\",\"col2\",\"col3\"]," +
		"\"params\":{" +
		"\"abundanceColumn\":\"abd\"," +
		"\"analysisType\":\"dotplot\"," +
		"\"baitColumn\":\"bait\"," +
		"\"clustering\":\"hierarchical\"," +
		"\"clusteringMethod\":\"complete\"," +
		"\"controlColumn\":\"control\"," +
		"\"distance\":\"euclidean\"," +
		"\"files\":[\"file1.txt\",\"file2.txt\"]," +
		"\"fillColor\":\"blueBlack\"," +
		"\"imageType\":\"heatmap\"," +
		"\"invert\":false," +
		"\"logBase\":\"none\"," +
		"\"maximumAbundance\":50," +
		"\"minimumAbundance\":0," +
		"\"normalization\":\"none\"," +
		"\"preyColumn\":\"prey\"," +
		"\"primaryFilter\":0.01," +
		"\"scoreColumn\":\"score\"," +
		"\"scoreType\":\"lte\"," +
		"\"secondaryFilter\":0.05}," +
		"\"rows\":[" +
		"{\"data\":[" +
		"{\"value\":1}," +
		"{\"value\":2}," +
		"{\"value\":3}]," +
		"\"name\":\"row1\"}," +
		"{\"data\":[" +
		"{\"value\":4}," +
		"{\"value\":5}," +
		"{\"value\":6}]," +
		"\"name\":\"row2\"}," +
		"{\"data\":[" +
		"{\"value\":7}," +
		"{\"value\":8}," +
		"{\"value\":9}]," +
		"\"name\":\"row3\"}]," +
		"\"minimap\":\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg==\"}"
	assert.Equal(
		t,
		want,
		InteractiveHeatmap(
			abundance,
			columns,
			rows,
			false,
			params,
			"test.png",
		),
		"JSON not generated correctly",
	)
}
