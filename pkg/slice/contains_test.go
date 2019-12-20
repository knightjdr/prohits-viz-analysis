package slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/slice"
)

var _ = Describe("Contains int", func() {
	It("should return true when slice contains value", func() {
		s := []int{1, 2, 3}
		tests := []int{1, 2, 3}
		for _, test := range tests {
			Expect(ContainsInt(test, s)).To(BeTrue())
		}
	})

	It("should return false when slice does not contain value", func() {
		s := []int{1, 2, 3}
		tests := []int{11, 22}
		for _, test := range tests {
			Expect(ContainsInt(test, s)).To(BeFalse())
		}
	})
})

var _ = Describe("Contains string", func() {
	It("should return true when slice contains value", func() {
		s := []string{"a", "c", "d"}
		tests := []string{"a", "c", "d"}
		for _, test := range tests {
			Expect(ContainsString(test, s)).To(BeTrue())
		}
	})

	It("should return false when slice does not contain value", func() {
		s := []string{"a", "c", "d"}
		tests := []string{"aa", "b", "something"}
		for _, test := range tests {
			Expect(ContainsString(test, s)).To(BeFalse())
		}
	})
})
