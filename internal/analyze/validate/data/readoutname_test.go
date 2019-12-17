package data

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate readout names", func() {
	It("should return nil when all readouts have names", func() {
		data := []map[string]string{
			map[string]string{"condition": "conditionA", "readout": "readoutA"},
			map[string]string{"condition": "conditionB", "readout": "readoutB"},
			map[string]string{"condition": "conditionA", "readout": "readoutC"},
			map[string]string{"condition": "conditionB", "readout": "readoutC"},
			map[string]string{"condition": "conditionC", "readout": "readoutA"},
			map[string]string{"condition": "conditionC", "readout": "readoutB"},
		}

		Expect(confirmReadoutsHaveNames(data)).To(BeNil())
	})

	It("should return an error when data does not satisfy minimum conditions", func() {
		data := []map[string]string{
			map[string]string{"condition": "conditionA", "readout": "readoutA"},
			map[string]string{"condition": "conditionB", "readout": "readoutB"},
			map[string]string{"condition": "conditionA", "readout": ""},
			map[string]string{"condition": "conditionB", "readout": "readoutC"},
			map[string]string{"condition": "conditionC", "readout": ""},
			map[string]string{"condition": "conditionC", "readout": "readoutB"},
		}

		expected := errors.New("all readouts should have a name")
		Expect(confirmReadoutsHaveNames(data)).To(Equal(expected))
	})
})
