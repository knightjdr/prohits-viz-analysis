package data

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate data for minimum conditions", func() {
	It("should return nil when data satisfies minimum conditions", func() {
		data := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA"},
			{"condition": "conditionB", "readout": "readoutB"},
			{"condition": "conditionA", "readout": "readoutC"},
			{"condition": "conditionB", "readout": "readoutC"},
			{"condition": "conditionC", "readout": "readoutA"},
			{"condition": "conditionC", "readout": "readoutB"},
		}
		analysisType := "dotplot"

		Expect(confirmMinimumConditions(data, analysisType)).To(BeNil())
	})

	It("should return an error when data does not satisfy minimum conditions", func() {
		data := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA"},
			{"condition": "conditionA", "readout": "readoutC"},
		}
		analysisType := "dotplot"

		expected := errors.New("there are not enough conditions for analysis, min: 2")
		Expect(confirmMinimumConditions(data, analysisType)).To(Equal(expected))
	})
})

var _ = Describe("Minimum required conditions", func() {
	It("should return 1 for circheatmap", func() {
		Expect(getMinimumRequiredConditions("circheatmap")).To(Equal(1))
	})

	It("should return 2 for doplot", func() {
		Expect(getMinimumRequiredConditions("doplot")).To(Equal(2))
	})
})

var _ = Describe("Count unique conditions", func() {
	It("should count conditions but return once minimum is reach", func() {
		data := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA"},
			{"condition": "conditionB", "readout": "readoutB"},
			{"condition": "conditionA", "readout": "readoutC"},
			{"condition": "conditionB", "readout": "readoutC"},
			{"condition": "conditionC", "readout": "readoutA"},
			{"condition": "conditionC", "readout": "readoutB"},
		}
		minimumRequiredConditions := 2
		Expect(countUniqueConditions(data, minimumRequiredConditions)).To(Equal(2))
	})

	It("should return conditions length when minimum is not reached", func() {
		data := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA"},
			{"condition": "conditionB", "readout": "readoutB"},
			{"condition": "conditionA", "readout": "readoutC"},
			{"condition": "conditionB", "readout": "readoutC"},
			{"condition": "conditionC", "readout": "readoutA"},
			{"condition": "conditionC", "readout": "readoutB"},
		}
		minimumRequiredConditions := 4
		Expect(countUniqueConditions(data, minimumRequiredConditions)).To(Equal(3))
	})
})
