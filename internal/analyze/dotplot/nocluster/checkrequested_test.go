package nocluster

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check input list", func() {
	It("should return requested names that are present in data file", func() {
		fileData := []map[string]string{
			{"condition": "a"},
			{"condition": "b"},
			{"condition": "a"},
			{"condition": "c"},
		}
		requestedList := []string{"b", "a", "d"}

		expected := []string{"b", "a"}
		Expect(checkRequestedList(fileData, "condition", requestedList)).To(Equal(expected))
	})
})
