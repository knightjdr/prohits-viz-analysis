package treeview

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get labeller function", func() {
	It("should create a function for labelling tree nodes/leafs", func() {
		sorted := []string{"a", "c", "b"}
		unsorted := []string{"a", "b", "c"}

		label := getTreeLabeler(sorted, unsorted, "label")

		Expect(label(0)).To(Equal("label0X"), "should return label for first leaf indexed from zero")
		Expect(label(1)).To(Equal("label2X"), "should return label for second sorted leaf")
		Expect(label(3)).To(Equal("NODE1X"), "should return label for first node indexed from one")
	})
})
