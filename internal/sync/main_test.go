package sync

import (
	"encoding/base64"
	"os"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var dataText = `{
	"imageType": "heatmap",
	"settings": {"abundanceCap":4,"fillColor":"blue","imageType":"heatmap","minAbundance":0,"xLabel":"Bait","yLabel":"Prey"},
	"columnOrder": [0, 1, 2, 3],
	"rowOrder": [0, 1, 2, 3],
	"rowDB": [
		{
			"name": "prey1",
			"data": [
				{"value": 1},
				{"value": 1},
				{"value": 2},
				{"value": 2}
			]
		},
		{
			"name": "prey2",
			"data": [
				{"value": 2},
				{"value": 2},
				{"value": 2},
				{"value": 3}
			]
		},
		{
			"name": "prey3",
			"data": [
				{"value": 1},
				{"value": 1},
				{"value": 4},
				{"value": 1}
			]
		},
		{
			"name": "prey4",
			"data": [
				{"value": 4},
				{"value": 2},
				{"value": 2},
				{"value": 3}
			]
		}
	]
}`

var _ = Describe("Sync minimap", func() {
	It("should create minimap", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		afero.WriteFile(fs.Instance, "data.json", []byte(dataText), 0644)

		os.Args = []string{
			"cmd",
			"--file", "data.json",
		}

		expected := "iVBORw0KGgoAAAANSUhEUgAAAFAAAABQCAIAAAABc2X6AAAAsElEQVR4nOzaw" +
			"QnCQBBA0URSmJ2InayViJ3ZiWIHkhwWft67L8Nnb8Ns4/lZZhivKWOXy5yx8wiuE1wnuE5wne" +
			"A6wXWC6wTXCa5bl+ucndYh78fup6f7YcF1gusE1wmuE1wnuE5wneA6wXXbuM0ZPO7r/rcHbst" +
			"O98OC6wTXCa4TXCe4TnCd4DrBdYLr9i+WfmbdeLnT+p/gOsF1gusE1wmuE1wnuE5wneC6bwAA" +
			"AP//BQENAfnS30gAAAAASUVORK5CYII="

		Minimap()
		pngContent, _ := afero.ReadFile(fs.Instance, "minimap/minimap.png")
		Expect(base64.StdEncoding.EncodeToString((pngContent))).To(Equal(expected))
	})
})
