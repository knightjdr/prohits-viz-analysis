package cc

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write scatter data to file", func() {
	It("should write data to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("other", 0755)

		data := []types.ScatterPoint{
			{Label: "readoutA", X: 1, Y: 3, Color: "#0066cc"},
			{Label: "readoutB", X: 0, Y: 4, Color: "#99ccff"},
			{Label: "readoutC", X: 2, Y: 0, Color: "#99ccff"},
		}
		settings := types.Settings{
			Readout: "Gene",
		}

		writeData(data, settings)

		expected := "Gene\tx\ty\n" +
			"readoutA\t1.00\t3.00\n" +
			"readoutB\t0.00\t4.00\n" +
			"readoutC\t2.00\t0.00\n"
		actual, _ := afero.ReadFile(fs.Instance, "other/x-y-data.txt")
		Expect(string(actual)).To(Equal(expected))
	})
})
