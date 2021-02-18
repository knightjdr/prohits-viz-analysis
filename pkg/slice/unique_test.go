package slice

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unique ints", func() {
	It("should remove duplicates", func() {
		s := []int{1, 2, 4, 4, 2, 2}

		expected := []int{1, 2, 4}
		Expect(UniqueInts(s)).To(Equal(expected))
	})
})

var _ = Describe("Unique strings", func() {
	It("should remove duplicates", func() {
		s := []string{"a", "b", "c", "c", "c", "b"}

		expected := []string{"a", "b", "c"}
		Expect(UniqueStrings(s)).To(Equal(expected))
	})
})
