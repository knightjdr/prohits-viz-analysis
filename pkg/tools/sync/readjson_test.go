package sync

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/matrix/frontend"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var jsonText = `{
	"columnDB": ["column1", "column2", "column3"],
	"columnOrder": [1, 2, 0],
	"imageType": "dotplot",
	"rowOrder": [0, 1, 3],
	"rowDB": [
		{
			"name": "row1",
			"data": [{"value": 1}, {"value": 2}, {"value": 3}]
		},
		{
			"name": "row2",
			"data": [{"value": 4}, {"value": 5}, {"value": 6}]
		},
		{
			"name": "row3",
			"data": [{"value": 7}, {"value": 8}, {"value": 9}]
		},
		{
			"name": "row4",
			"data": [{"value": 10}, {"value": 11}, {"value": 12}]
		}
	],
	"settings": {
		"abundanceCap": 50,
		"scoreType": "lte"
	}
}`

var _ = Describe("Read JSON", func() {
	It("should read json", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/file.json", []byte(jsonText), 0644)

		expected := &minimap{
			ColumnDB:    []string{"column1", "column2", "column3"},
			ColumnOrder: []int{1, 2, 0},
			ImageType:   "dotplot",
			RowDB: []frontend.Row{
				{Name: "row1", Data: []frontend.Cell{{Value: 1}, {Value: 2}, {Value: 3}}},
				{Name: "row2", Data: []frontend.Cell{{Value: 4}, {Value: 5}, {Value: 6}}},
				{Name: "row3", Data: []frontend.Cell{{Value: 7}, {Value: 8}, {Value: 9}}},
				{Name: "row4", Data: []frontend.Cell{{Value: 10}, {Value: 11}, {Value: 12}}},
			},
			RowOrder: []int{0, 1, 3},
			Settings: types.Settings{
				AbundanceCap: 50,
				ScoreType:    "lte",
			},
		}
		Expect(readJSON("test/file.json")).To(Equal(expected))
	})
})
