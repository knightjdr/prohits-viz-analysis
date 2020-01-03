package files_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knightjdr/prohits-viz-analysis/internal/pkg/files"
)

var _ = Describe("Names", func() {
	It("should parse file names", func() {
		files := []string{"file.txt", "some/path/file2.svg", "/path/file3.png"}

		expected := []string{"file.txt", "file2.svg", "file3.png"}
		Expect(ParseBaseNames(files)).To(Equal(expected))
	})
})
