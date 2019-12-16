package parser

import (
	"bytes"
	"encoding/csv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CSV reader", func() {
	It("should parse unix line endings", func() {
		testFile := bytes.NewBufferString("column1\tcolumn2\tcolumn3\na\tb\tc\n").Bytes()

		r2 := csv.NewReader(createReader(bytes.NewReader(testFile)))
		parsed, _ := r2.ReadAll()

		expected := [][]string{
			[]string{"column1	column2	column3"},
			[]string{"a	b	c"},
		}
		Expect(parsed).To(Equal(expected))
	})

	It("should parse windows line endings", func() {
		testFile := bytes.NewBufferString("column1\tcolumn2\tcolumn3\r\na\tb\tc\r\n").Bytes()

		r2 := csv.NewReader(createReader(bytes.NewReader(testFile)))
		parsed, _ := r2.ReadAll()

		expected := [][]string{
			[]string{"column1	column2	column3"},
			[]string{"a	b	c"},
		}
		Expect(parsed).To(Equal(expected))
	})

	It("should parse classic mac osx line endings", func() {
		testFile := bytes.NewBufferString("column1\tcolumn2\tcolumn3\ra\tb\tc\r").Bytes()

		r2 := csv.NewReader(createReader(bytes.NewReader(testFile)))
		parsed, _ := r2.ReadAll()

		expected := [][]string{
			[]string{"column1	column2	column3"},
			[]string{"a	b	c"},
		}
		Expect(parsed).To(Equal(expected))
	})
})
