package geneid

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/spf13/afero"
)

var tsvFile = `A_X	1
C	3
D	
`

var _ = Describe("Map by file", func() {
	It("should map IDs using a file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		filename := "test/mapping.csv"
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, filename, []byte(tsvFile), 0444)

		data := []map[string]string{
			{"condition": "A_X"},
			{"condition": "B"},
			{"condition": "C"},
			{"condition": "D_X"},
		}
		sourceColumn := "condition"

		expected := map[string]string{
			"A_X": "1",
			"B":   "B",
			"C":   "3",
			"D_X": "D",
		}
		Expect(MapByFile(data, sourceColumn, filename)).To(Equal(expected))
	})
})
