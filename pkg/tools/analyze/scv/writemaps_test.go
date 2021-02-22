package scv

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

var _ = Describe("Write id maps to file", func() {
	It("should write data to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("other", 0755)

		idMaps := map[string]map[string]string{
			"condition": {
				"A_X": "1",
				"B":   "",
				"C":   "3",
				"D":   "D",
			},
			"readout": {
				"X": "24",
				"Y": "",
				"Z": "26",
			},
		}
		settings := types.Settings{
			Condition: "Bait",
			Readout:   "Prey",
		}

		writeMaps(idMaps, settings)

		expectedCondition := "A_X\t1\n" +
			"B\t\n" +
			"C\t3\n" +
			"D\tD\n"
		expectedReadout := "X\t24\n" +
			"Y\t\n" +
			"Z\t26\n"

		actualCondition, _ := afero.ReadFile(fs.Instance, "other/map-Bait.txt")
		actualReadout, _ := afero.ReadFile(fs.Instance, "other/map-Prey.txt")
		Expect(string(actualCondition)).To(Equal(expectedCondition))
		Expect(string(actualReadout)).To(Equal(expectedReadout))
	})
})
