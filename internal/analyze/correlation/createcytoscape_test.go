package correlation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/spf13/afero"
)

var _ = Describe("Create folders", func() {
	It("create default folders", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("cytoscape", 0755)

		data := &correlationData{
			matrix: [][]float64{
				{1, 0.8, 0.5},
				{0.8, 1, 0.7},
				{0.5, 0.7, 1},
			},
			sortedLabels: []string{"b", "a", "c"},
		}

		expected := "test\ttest\tcorrelation\n" +
			"b\ta\t0.8\n" +
			"a\tc\t0.7\n"

		createCytoscape(data, 0.7, "test")
		actual, _ := afero.ReadFile(fs.Instance, "cytoscape/test-test.txt")
		Expect(string(actual)).To(Equal(expected))
	})
})
