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

		Expect(ConvertToBoolMap(s)).To(Equal(expected))
	})

	It("should conver a slice of strings to a map of ints", func() {
		s := []string{"a", "b", "c"}

		expected := map[string]int{
			"a": 0,
			"b": 1,
			"c": 2,
		}

		Expect(ConvertToIntMap(s)).To(Equal(expected))
	})
})
