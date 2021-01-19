package specificity

import (
	"math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write specificity data to file", func() {
	It("should write data to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("other", 0755)

		data := map[string]map[string]map[string]float64{
			"a": {
				"x": {"abundance": 10, "score": 0.01, "specificity": 0.67},
				"y": {"abundance": 20, "score": 0.01, "specificity": 2.67},
			},
			"b": {
				"x": {"abundance": 30, "score": 0, "specificity": 6},
			},
			"c": {
				"y": {"abundance": 15, "score": 0.02, "specificity": 1.5},
				"z": {"abundance": 25, "score": 0.01, "specificity": math.Inf(1)},
			},
		}
		settings := types.Settings{
			Abundance: "AvgSpec",
			Condition: "Bait",
			Readout:   "Prey",
			Score:     "FDR",
		}

		writeData(data, settings)

		expected := "Bait\tPrey\tAvgSpec\tSpecificity\tFDR\n" +
			"a\tx\t10.00\t0.67\t0.01\n" +
			"a\ty\t20.00\t2.67\t0.01\n" +
			"b\tx\t30.00\t6.00\t0.00\n" +
			"c\ty\t15.00\t1.50\t0.02\n" +
			"c\tz\t25.00\t+Inf\t0.01\n"
		actual, _ := afero.ReadFile(fs.Instance, "other/specificity-data.txt")
		Expect(string(actual)).To(Equal(expected))
	})
})
