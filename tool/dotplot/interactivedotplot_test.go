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

func TestInteractiveDotplot(t *testing.T) {
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
		ColorSpace:       "blueBlack",
		Control:          "control",
		Distance:         "euclidean",
		Files:            []string{"file1.txt", "file2.txt"},
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
	ratios := [][]float64{
		{0.2, 0.5, 1},
		{0.7, 0.8, 1},
		{1, 0.2, 0.5},
	}
	rows := []string{"row1", "row2", "row3"}
	scores := [][]float64{
		{0.01, 0.05, 0.08},
		{1, 0.07, 0.5},
		{0.2, 0.7, 0.01},
	}

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
		"\"colorSpace\":\"blueBlack\"," +
		"\"controlColumn\":\"control\"," +
		"\"distance\":\"euclidean\"," +
		"\"files\":[\"file1.txt\",\"file2.txt\"]," +
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
		"{\"ratio\":0.2,\"score\":0.01,\"value\":1}," +
		"{\"ratio\":0.5,\"score\":0.05,\"value\":2}," +
		"{\"ratio\":1,\"score\":0.08,\"value\":3}]," +
		"\"name\":\"row1\"}," +
		"{\"data\":[" +
		"{\"ratio\":0.7,\"score\":1,\"value\":4}," +
		"{\"ratio\":0.8,\"score\":0.07,\"value\":5}," +
		"{\"ratio\":1,\"score\":0.5,\"value\":6}]," +
		"\"name\":\"row2\"}," +
		"{\"data\":[" +
		"{\"ratio\":1,\"score\":0.2,\"value\":7}," +
		"{\"ratio\":0.2,\"score\":0.7,\"value\":8}," +
		"{\"ratio\":0.5,\"score\":0.01,\"value\":9}]," +
		"\"name\":\"row3\"}]," +
		"\"minimap\":\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg==\"}"
	assert.Equal(
		t,
		want,
		InteractiveDotplot(
			abundance,
			ratios,
			scores,
			columns,
			rows,
			params,
			"test.png",
		),
		"JSON not generated correctly",
	)
}
