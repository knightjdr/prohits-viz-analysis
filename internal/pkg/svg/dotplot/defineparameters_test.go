package dotplot

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Define dotplot parameters", func() {
	It("should define parameters", func() {
		d := &Dotplot{
			CellSize: 20,
			Ratio:    1,
		}

		expected := dotplotparameters{
			cellSizeHalf: 10,
			edgeWidth:    2,
			maxRadius:    8.5,
		}
		Expect(defineParameters(d)).To(Equal(expected))
	})
})
