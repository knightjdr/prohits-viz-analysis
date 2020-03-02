package hierarchical

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write matrices to file", func() {
	It("should write matrices to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("other", 0755)

		data := &SortedData{
			Matrices: &types.Matrices{
				Abundance: [][]float64{
					{0, 10, 74.2},
					{5, 7.2, 90.12},
					{8.3, 2, 1.4},
				},
				Conditions: []string{"condition1", "condition2", "condition3"},
				Ratio: [][]float64{
					{5, 7.2, 90.12},
					{0, 10, 74.2},
					{8.3, 2, 1.4},
				},
				Readouts: []string{"readout1", "readout2", "readout3"},
			},
		}

		WriteMatrices(data)

		expectedAbundance := "\tcondition1\tcondition2\tcondition3\n" +
			"readout1\t0.00\t10.00\t74.20\n" +
			"readout2\t5.00\t7.20\t90.12\n" +
			"readout3\t8.30\t2.00\t1.40\n"
		actualAbundance, _ := afero.ReadFile(fs.Instance, "other/data-transformed.txt")
		Expect(string(actualAbundance)).To(Equal(expectedAbundance), "should write transformed abundance data to file")

		expectedRatio := "\tcondition1\tcondition2\tcondition3\n" +
			"readout1\t5.00\t7.20\t90.12\n" +
			"readout2\t0.00\t10.00\t74.20\n" +
			"readout3\t8.30\t2.00\t1.40\n"
		actualRatio, _ := afero.ReadFile(fs.Instance, "other/data-transformed-ratios.txt")
		Expect(string(actualRatio)).To(Equal(expectedRatio), "should write transformed ratio data to file")
	})
})

var _ = Describe("Write matrix to file", func() {
	It("should write matrix to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		conditions := []string{"condition1", "condition2", "condition3"}
		matrix := [][]float64{
			{0, 10, 74.2},
			{5, 7.2, 90.12},
			{8.3, 2, 1.4},
		}
		readouts := []string{"readout1", "readout2", "readout3"}

		expected := "\tcondition1\tcondition2\tcondition3\n" +
			"readout1\t0.00\t10.00\t74.20\n" +
			"readout2\t5.00\t7.20\t90.12\n" +
			"readout3\t8.30\t2.00\t1.40\n"

		writeMatrix(matrix, conditions, readouts, "other/test.txt")
		actual, _ := afero.ReadFile(fs.Instance, "other/test.txt")
		Expect(string(actual)).To(Equal(expected))
	})
})
