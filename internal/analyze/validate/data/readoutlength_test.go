package data

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate readout lengths", func() {
	It("should return nil when no readout length columns is specified", func() {
		data := []map[string]string{
			map[string]string{"condition": "conditionA", "readout": "readoutA", "readoutLength": "10"},
			map[string]string{"condition": "conditionB", "readout": "readoutB", "readoutLength": "15"},
			map[string]string{"condition": "conditionA", "readout": "readoutC", "readoutLength": "25"},
			map[string]string{"condition": "conditionB", "readout": "readoutC", "readoutLength": "7"},
			map[string]string{"condition": "conditionC", "readout": "readoutA", "readoutLength": "8"},
			map[string]string{"condition": "conditionC", "readout": "readoutB", "readoutLength": "12"},
		}
		readoutLength := ""

		Expect(confirmReadLengthIsInt(data, readoutLength)).To(BeNil())
	})

	It("should return nil when readout length columns contains ints", func() {
		data := []map[string]string{
			map[string]string{"condition": "conditionA", "readout": "readoutA", "readoutLength": "10"},
			map[string]string{"condition": "conditionB", "readout": "readoutB", "readoutLength": "15"},
			map[string]string{"condition": "conditionA", "readout": "readoutC", "readoutLength": "25"},
			map[string]string{"condition": "conditionB", "readout": "readoutC", "readoutLength": "7"},
			map[string]string{"condition": "conditionC", "readout": "readoutA", "readoutLength": "8"},
			map[string]string{"condition": "conditionC", "readout": "readoutB", "readoutLength": "12"},
		}
		readoutLength := "PreyLength"

		Expect(confirmReadLengthIsInt(data, readoutLength)).To(BeNil())
	})

	It("should return an error when a readout length value is not parseable as an int", func() {
		data := []map[string]string{
			map[string]string{"condition": "conditionA", "readout": "readoutA", "readoutLength": "10"},
			map[string]string{"condition": "conditionB", "readout": "readoutB", "readoutLength": "15"},
			map[string]string{"condition": "conditionA", "readout": "readoutC", "readoutLength": "25"},
			map[string]string{"condition": "conditionB", "readout": "readoutC", "readoutLength": "7.3"},
			map[string]string{"condition": "conditionC", "readout": "readoutA", "readoutLength": "8"},
			map[string]string{"condition": "conditionC", "readout": "readoutB", "readoutLength": "12"},
		}
		readoutLength := "PreyLength"

		expected := errors.New("readout length column must contain integer values, offending value: 7.3")
		Expect(confirmReadLengthIsInt(data, readoutLength)).To(Equal(expected))
	})
})
