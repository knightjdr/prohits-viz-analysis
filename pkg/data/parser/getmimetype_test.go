package parser

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/spf13/afero"
)

var _ = Describe("Get mime type", func() {
	It("should return mime type", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/testfile.txt", []byte("a\tb\tc\n"), 0444)

		expected := "text/plain"
		actualMimeType, err := getMimeType("test/testfile.txt")

		Expect(err).To(BeNil(), "should not return an error")
		Expect(actualMimeType).To(Equal(expected), "should return correct mime type")
	})

	It("should return 'unknown' for missing file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		expected := "unknown"
		actualMimeType, err := getMimeType("test/missing.txt")

		Expect(err).To(Not(BeNil()), "should return an error")
		Expect(actualMimeType).To(Equal(expected), "should return unknown for mime type")
	})

	It("should return 'unknown' for unreadable file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/unreadable.txt", []byte(""), 0444) // Unreadable because empty.

		expected := "unknown"
		actualMimeType, err := getMimeType("test/unreadable.txt")

		Expect(err).To(Not(BeNil()), "should return an error")
		Expect(actualMimeType).To(Equal(expected), "should return unknown for mime type")
	})
})
