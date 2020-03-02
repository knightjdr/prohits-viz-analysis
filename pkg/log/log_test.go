package log_test

import (
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/spf13/afero"
)

var _ = Describe("Log message", func() {
	It("should log message to file", func() {
		// Mock filesystem.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		log.Write("test message")
		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		expected := "test message"
		matched, _ := regexp.MatchString(expected, string(logfile))
		Expect(matched).To(BeTrue())
	})
})
