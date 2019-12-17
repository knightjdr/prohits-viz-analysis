package float_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/float"
)

var _ = Describe("Join", func() {
	It("should join a slice of float64", func() {
		s := []float64{1, 2, 3.3, 2.2}
		sep := "|"

		expected := "1|2|3.3|2.2"
		Expect(Join(s, sep)).To(Equal(expected))
	})
})
