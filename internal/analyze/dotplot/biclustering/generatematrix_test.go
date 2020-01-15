package biclustering

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Generate matrix for nested cluster", func() {
	It("should write matrix to file and return list of singletons", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("biclustering", 0755)

		matrices := &types.Matrices{
			Abundance: [][]float64{
				{2, 4, 8},
				{4, 8, 16},
				{2, 8, 8},
			},
			Conditions: []string{"condition1", "condition2", "condition3"},
			Readouts:   []string{"readout1", "readout2", "readout3"},
		}

		actualSingles := generateMatrix(matrices, 8)

		expectedSingles := []string{"readout1"}
		Expect(actualSingles).To(Equal(expectedSingles), "should return a list of singletons")

		expectedMatrix := "PROT\tcondition1\tcondition2\tcondition3\n" +
			"readout2\t2.50000\t5.00000\t10.00000\n" +
			"readout3\t1.25000\t5.00000\t5.00000\n"
		actualMatrix, _ := afero.ReadFile(fs.Instance, "biclustering/matrix.txt")
		Expect(string(actualMatrix)).To(Equal(expectedMatrix), "should write matrix to file")
	})
})
