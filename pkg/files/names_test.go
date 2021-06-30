package files_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/files"
)

var _ = Describe("Names", func() {
	It("should parse file names from a slices of names", func() {
		files := []string{"file.txt", "some/path/file2.svg", "/path/file3.png"}

		expected := []string{"file.txt", "file2.svg", "file3.png"}
		Expect(ParseBaseNames(files)).To(Equal(expected))
	})

	It("should parse file name without extension", func() {
		files := []string{"file.txt", "some/path/file2.svg", "/path/file3.png"}

		expected := []string{"file", "file2", "file3"}
		for i, file := range files {
			Expect(ParseBaseNameWithoutExtension(file)).To(Equal(expected[i]))
		}
	})
})
