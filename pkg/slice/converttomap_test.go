package slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/slice"
)

var _ = Describe("Convert to map", func() {
	It("should conver a slice of strings to a map of booleans", func() {
		s := []string{"a", "b", "c"}

		expected := map[string]bool{
			"a": true,
			"b": true,
			"c": true,
		}

		Expect(ConvertToMap(s)).To(Equal(expected))
	})
})
