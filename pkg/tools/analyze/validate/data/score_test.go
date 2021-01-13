package data

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate score", func() {
	It("should return nil when score column contains numeric values", func() {
		data := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA", "score": "10"},
			{"condition": "conditionB", "readout": "readoutB", "score": "15.5"},
			{"condition": "conditionA", "readout": "readoutC", "score": "25.3"},
			{"condition": "conditionB", "readout": "readoutC", "score": "7"},
			{"condition": "conditionC", "readout": "readoutA", "score": "8"},
			{"condition": "conditionC", "readout": "readoutB", "score": "12"},
		}

		Expect(confirmScoreIsFloat(data)).To(BeNil())
	})

	It("should return an error when score column contains an NaN", func() {
		data := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA", "score": "10"},
			{"condition": "conditionB", "readout": "readoutB", "score": "15"},
			{"condition": "conditionA", "readout": "readoutC", "score": "25"},
			{"condition": "conditionB", "readout": "readoutC", "score": "NaN"},
			{"condition": "conditionC", "readout": "readoutA", "score": "8"},
			{"condition": "conditionC", "readout": "readoutB", "score": "12"},
		}

		expected := errors.New("score column must contain numeric values, offending value: NaN")
		Expect(confirmScoreIsFloat(data)).To(Equal(expected))
	})

	It("should return an error when score column contains a value not parseable as a float", func() {
		data := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA", "score": "10"},
			{"condition": "conditionB", "readout": "readoutB", "score": "15"},
			{"condition": "conditionA", "readout": "readoutC", "score": "25"},
			{"condition": "conditionB", "readout": "readoutC", "score": "a"},
			{"condition": "conditionC", "readout": "readoutA", "score": "8"},
			{"condition": "conditionC", "readout": "readoutB", "score": "12"},
		}

		expected := errors.New("score column must contain numeric values, offending value: a")
		Expect(confirmScoreIsFloat(data)).To(Equal(expected))
	})
})
