package slice

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dict strings", func() {
	It("should convert slice to map", func() {
		s := []string{"a", "b", "c"}

		expected := map[string]bool{
			"a": true,
			"b": true,
			"c": true,
		}
		Expect(Dict(s)).To(Equal(expected))
	})
})

var _ = Describe("Dict int", func() {
	It("should convert slice to map", func() {
		s := []int{1, 2, 4}

		expected := map[int]bool{
			1: true,
			2: true,
			4: true,
		}
		Expect(DictInt(s)).To(Equal(expected))
	})
})
