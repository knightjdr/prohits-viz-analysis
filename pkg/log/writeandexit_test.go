package log_test

import (
	"os"
	"regexp"

	"github.com/bouk/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/spf13/afero"
)

var _ = Describe("Write message and exit", func() {
	It("should log message and exit", func() {
		// Mock filesystem.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		// Mock exit.
		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		Expect(func() { log.WriteAndExit("test message") }).To(Panic())

		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		expected := "test message"
		matched, _ := regexp.MatchString(expected, string(logfile))
		Expect(matched).To(BeTrue())
	})
})
