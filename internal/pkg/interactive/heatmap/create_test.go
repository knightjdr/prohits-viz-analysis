package heatmap_test

import (
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/interactive/heatmap"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Create file", func() {
	It("should create file for interactive heatmap viewer", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		pngImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1, 1}})
		c := color.RGBA{uint8(0), uint8(0), uint8(0), 255}
		pngImage.Set(0, 0, c)
		myfile, _ := fs.Instance.Create("test/image.png")
		png.Encode(myfile, pngImage)

		data := &Data{
			Filename:  "test/heatmap.json",
			ImageType: "heatmap",
			Matrices: &types.Matrices{
				Abundance: [][]float64{
					{25, 5, 50.2},
					{100, 30, 7},
					{5, 2.3, 8},
				},
				Conditions: []string{"bait1", "bait2", "bait3"},
				Readouts:   []string{"prey1", "prey2", "prey3"},
			},
			Minimap: "test/image.png",
			Settings: types.Settings{
				Abundance:          "AvgSpec",
				AbundanceCap:       50,
				Clustering:         "hierarchical",
				ClusteringMethod:   "complete",
				ClusteringOptimize: true,
				Condition:          "Bait",
				Control:            "ctrl",
				Distance:           "canberra",
				Files:              []string{"file1", "file2"},
				FillColor:          "blue",
				InvertColor:        true,
				LogBase:            "2",
				MinAbundance:       10,
				Normalization:      "total",
				PrimaryFilter:      0.01,
				Readout:            "Prey",
				Score:              "bfdr",
				ScoreType:          "lte",
				Type:               "dotplot",
				XLabel:             "Prey",
				YLabel:             "Bait",
			},
		}

		uri := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR4nGJiYGAABAAA//8ADAADcZGLFwAAAABJRU5ErkJggg=="
		expected := "{\n" +
			"\t\"parameters\": {\"abundanceColumn\":\"AvgSpec\",\"analysisType\":\"dotplot\",\"clustering\":\"hierarchical\",\"clusteringMethod\":\"complete\",\"clusteringOptimize\":true,\"conditionColumn\":\"Bait\",\"controlColumn\":\"ctrl\",\"distance\":\"canberra\",\"files\":[\"file1\",\"file2\"],\"imageType\":\"heatmap\",\"logBase\":\"2\",\"normalization\":\"total\",\"readoutColumn\":\"Prey\",\"scoreColumn\":\"bfdr\",\"scoreType\":\"lte\",\"xLabel\":\"Prey\",\"yLabel\":\"Bait\"},\n" +
			"\t\"settings\": {\"main\":{\"current\":{\"abundanceCap\":50,\"fillColor\":\"blue\",\"imageType\":\"heatmap\",\"invertColor\":true,\"minAbundance\":10,\"primaryFilter\":0.01}}},\n" +
			"\t\"columnDB\": [\"bait1\",\"bait2\",\"bait3\"],\n" +
			"\t\"rowDB\": [\n" +
			"\t\t{\n" +
			"\t\t\t\"name\": \"prey1\",\n" +
			"\t\t\t\"data\": [\n" +
			"\t\t\t\t{\"value\": 25.00},\n" +
			"\t\t\t\t{\"value\": 5.00},\n" +
			"\t\t\t\t{\"value\": 50.20}\n" +
			"\t\t\t]\n" +
			"\t\t},\n" +
			"\t\t{\n" +
			"\t\t\t\"name\": \"prey2\",\n" +
			"\t\t\t\"data\": [\n" +
			"\t\t\t\t{\"value\": 100.00},\n" +
			"\t\t\t\t{\"value\": 30.00},\n" +
			"\t\t\t\t{\"value\": 7.00}\n" +
			"\t\t\t]\n" +
			"\t\t},\n" +
			"\t\t{\n" +
			"\t\t\t\"name\": \"prey3\",\n" +
			"\t\t\t\"data\": [\n" +
			"\t\t\t\t{\"value\": 5.00},\n" +
			"\t\t\t\t{\"value\": 2.30},\n" +
			"\t\t\t\t{\"value\": 8.00}\n" +
			"\t\t\t]\n" +
			"\t\t}\n" +
			"\t],\n" +
			fmt.Sprintf("\t\"minimap\": {\"main\":{\"image\":\"%s\"}}\n", uri) +
			"}\n"

		Create(data)

		actual, _ := afero.ReadFile(fs.Instance, "test/heatmap.json")
		Expect(string(actual)).To(Equal(expected))
	})
})
