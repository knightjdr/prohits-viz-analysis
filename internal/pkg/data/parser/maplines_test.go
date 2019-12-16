package parser

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map lines", func() {
	It("should map lines", func() {
		headerMap := map[string]int{
			"key1": 0,
			"key2": 2,
		}
		lines := [][]string{
			{"condition1", "x", "readout1"},
			{"condition2", "x", "readout2"},
		}

		expected := []map[string]string{
			{"key1": "condition1", "key2": "readout1"},
			{"key1": "condition2", "key2": "readout2"},
		}
		Expect(mapLines(lines, headerMap)).To(Equal(expected))
	})
})
