package convert

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sort labels", func() {
	It("should sort labels by map value", func() {
		labels := map[string]int{
			"B": 2,
			"C": 0,
			"A": 4,
			"E": 1,
			"D": 3,
		}

		expected := []string{"C", "E", "B", "D", "A"}
		Expect(sortLabels(labels, false)).To(Equal(expected))
	})

	It("should sort labels alphabetically", func() {
		labels := map[string]int{
			"B": 2,
			"C": 0,
			"A": 4,
			"E": 1,
			"D": 3,
		}

		expected := []string{"A", "B", "C", "D", "E"}
		Expect(sortLabels(labels, true)).To(Equal(expected))
	})
})
