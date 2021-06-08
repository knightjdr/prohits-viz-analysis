package settings

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Convert settings", func() {
	It("should convert file settings type to analysis settings type", func() {
		fileSettings := &jsonSettings{
			AbundanceColumn: "AvgSpec",
			InvertColor:     1,
			PrimaryFilter:   0.01,
			ScoreColumn:     "BFDR",
			ScoreType:       0,
			SecondaryFilter: 0.05,
			Type:            "dotplot",
			XLabel:          "Prey",
			YLabel:          "Bait",
		}

		expected := types.Settings{
			Abundance:       "AvgSpec",
			InvertColor:     true,
			PrimaryFilter:   0.01,
			Score:           "BFDR",
			ScoreType:       "lte",
			SecondaryFilter: 0.05,
			Type:            "dotplot",
			XLabel:          "Prey",
			YLabel:          "Bait",
		}
		Expect(convert(fileSettings)).To(Equal(expected))
	})
})

var _ = Describe("Convert invert color", func() {
	It("should return true when value is 1", func() {
		Expect(invertColorToBool(1)).To(BeTrue())
	})

	It("should return false when value is 0", func() {
		Expect(invertColorToBool(0)).To(BeFalse())
	})
})

var _ = Describe("Convert score type", func() {
	It("should return 'gte' when value is 1", func() {
		Expect(scoreTypeToBool(1)).To(Equal("gte"))
	})

	It("should return 'lte' when value is 0", func() {
		Expect(scoreTypeToBool(0)).To(Equal("lte"))
	})
})
