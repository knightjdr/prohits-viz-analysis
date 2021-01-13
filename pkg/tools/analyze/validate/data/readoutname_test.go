package data

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate readout names", func() {
	It("should return nil when all readouts have names", func() {
		data := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA"},
			{"condition": "conditionB", "readout": "readoutB"},
			{"condition": "conditionA", "readout": "readoutC"},
			{"condition": "conditionB", "readout": "readoutC"},
			{"condition": "conditionC", "readout": "readoutA"},
			{"condition": "conditionC", "readout": "readoutB"},
		}

		Expect(confirmReadoutsHaveNames(data)).To(BeNil())
	})

	It("should return an error when data does not satisfy minimum conditions", func() {
		data := []map[string]string{
			{"condition": "conditionA", "readout": "readoutA"},
			{"condition": "conditionB", "readout": "readoutB"},
			{"condition": "conditionA", "readout": ""},
			{"condition": "conditionB", "readout": "readoutC"},
			{"condition": "conditionC", "readout": ""},
			{"condition": "conditionC", "readout": "readoutB"},
		}

		expected := errors.New("all readouts should have a name")
		Expect(confirmReadoutsHaveNames(data)).To(Equal(expected))
	})
})
