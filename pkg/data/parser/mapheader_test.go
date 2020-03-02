package parser

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map header", func() {
	It("should map columns name to header indices", func() {
		columnMap := map[string]string{
			"key1": "column1",
			"key2": "column3",
		}
		header := []string{"column1", "column2", "column3"}
		
		expected := map[string]int{
			"key1": 0,
			"key2": 2,
		}

		actual, err := mapHeader(columnMap, header, false)
		Expect(err).To(BeNil(), "should not return an error")
		Expect(actual).To(Equal(expected), "should return column to header map")
	})

	It("should return an error whem column cannot be found", func() {
		columnMap := map[string]string{
			"key1": "column1",
			"key2": "column4",
		}
		header := []string{"column1", "column2", "column3"}

		_, err := mapHeader(columnMap, header, false)
		Expect(err).To(Not(BeNil()))
	})

	It("should not return an error whem column cannot be found and missing columns are ignored", func() {
		columnMap := map[string]string{
			"key1": "column1",
			"key2": "column4",
		}
		header := []string{"column1", "column2", "column3"}

		_, err := mapHeader(columnMap, header, true)
		Expect(err).To(BeNil())
	})

	It("should ignore empty map values", func() {
		columnMap := map[string]string{
			"key1": "column1",
			"key2": "column3",
			"key3": "",
		}
		header := []string{"column1", "column2", "column3"}

		expected := map[string]int{
			"key1": 0,
			"key2": 2,
		}

		actual, err := mapHeader(columnMap, header, false)
		Expect(err).To(BeNil(), "should not return an error")
		Expect(actual).To(Equal(expected), "should return column to header map")
	})
})
