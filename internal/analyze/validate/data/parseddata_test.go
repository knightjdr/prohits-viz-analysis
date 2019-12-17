package data

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate parsed data length", func() {
	It("should return nil when there is data", func() {
		data := []map[string]string{
			map[string]string{"condition": "conditionA", "readout": "readoutA"},
		}

		Expect(confirmParsedData(data)).To(BeNil())
	})

	It("should return an error when there is no parsed data", func() {
		data := []map[string]string{}

		expected := errors.New("no parsed results satisfying filter criteria")
		Expect(confirmParsedData(data)).To(Equal(expected))
	})
})
